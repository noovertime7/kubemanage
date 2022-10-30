package app

import (
	"fmt"
	"github.com/noovertime7/kubemanage/cmd/app/options"
	"github.com/noovertime7/kubemanage/pkg"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/router"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
)

func NewServerCommand() *cobra.Command {
	opts, err := options.NewOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}

	cmd := &cobra.Command{
		Use:  "gopixiu-server",
		Long: "The gopixiu server controller is a daemon that embeds the core control loops.",
		Run: func(cmd *cobra.Command, args []string) {
			if err = opts.Complete(); err != nil {
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
	// 设置核心应用接口
	kubemanage.Setup(opt)
	//初始化K8s client
	InitLocalK8s()
	// 初始化 api 路由
	InitRouters()
	// 初始化内置K8S组件
	return nil
}

func InitLocalK8s() {
	//初始化K8s client
	if err := kube.K8s.Init(); err != nil {
		pkg.Must(err)
	}
}

func InitRouters() {
	pkg.PrintColor()
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
