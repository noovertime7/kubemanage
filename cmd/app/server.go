package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/cmd/app/options"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/pkg/utils"
	"github.com/noovertime7/kubemanage/router"
	"github.com/noovertime7/kubemanage/runtime"
)

func NewServerCommand() *cobra.Command {
	opts, err := options.NewOptions()
	if err != nil {
		logger.LG.Fatal("unable to initialize command options: %v", zap.Error(err))
	}

	cmd := &cobra.Command{
		Use:  "kubemanage-server",
		Long: "The kubemanage server controller is a daemon that embeds the core control loops.",
		Run: func(cmd *cobra.Command, args []string) {
			if err = opts.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = opts.InitDB(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			if err = Run(opts); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 绑定命令行参数
	opts.BindFlags(cmd)
	return cmd
}

func Run(opt *options.Options) error {
	// 打印logo
	utils.PrintLogo()
	// 设置核心应用接口
	v1.Setup(opt)
	//初始化K8s client  TODO 未来移除
	InitLocalK8s()
	// 初始化 APIs 路由
	router.InstallRouters(opt)
	// 启动优雅服务
	runServer(opt)
	return nil
}

func InitLocalK8s() {
	//初始化K8s client
	if err := kube.K8s.Init(); err != nil {
		utils.Must(err)
	}
}

// 优雅启动貔貅服务
func runServer(opt *options.Options) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s", config.SysConfig.Default.ListenAddr),
		Handler: opt.GinEngine,
	}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		logger.LG.Info("Success", zap.String("starting kubemanage server running on", config.SysConfig.Default.ListenAddr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.LG.Fatal("failed to listen kubemanage server: ", zap.Error(err))
		}
	}()

	<-runtime.SystemContext.Done()

	// Close
	runtime.CloserHandler.Close()

	logger.LG.Info("shutting kubemanage server down ...")

	// The context is used to inform the server it has 5 seconds to finish the request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.LG.Fatal("kubemanage server forced to shutdown: ", zap.Error(err))
		os.Exit(1)
	}
	logger.LG.Info("kubemanage server exit successful")
}
