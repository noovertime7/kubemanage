package kube

import (
	"net/http"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dto/kubeDto"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/pkg/types"
)

type PodsGetter interface {
	Pods(cloud string) PodInterface
}

type PodInterface interface {
	WebShellHandler(webShellOptions *kubeDto.WebShellOptions, w http.ResponseWriter, r *http.Request) error
}

type pods struct {
	client  *kubernetes.Clientset
	cloud   string
	factory dao.ShareDaoFactory
}

func NewPods(c *kubernetes.Clientset, cloud string, factory dao.ShareDaoFactory) *pods {
	return &pods{
		client:  c,
		cloud:   cloud,
		factory: factory,
	}
}

func (c *pods) WebShellHandler(webShellOptions *kubeDto.WebShellOptions, w http.ResponseWriter, r *http.Request) error {
	log := logger.New(logger.LG)
	session, err := types.NewTerminalSession(w, r)
	if err != nil {
		return err
	}
	// 处理关闭
	defer func() {
		_ = session.Close()
	}()

	// 组装 POST 请求
	req := K8s.ClientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(webShellOptions.Pod).
		Namespace(webShellOptions.Namespace).
		SubResource("exec").
		VersionedParams(&coreV1.PodExecOptions{
			Container: webShellOptions.Container,
			Command:   []string{"/bin/sh"},
			Stderr:    true,
			Stdin:     true,
			Stdout:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	// remotecommand 主要实现了http 转 SPDY 添加X-Stream-Protocol-Version相关header 并发送请求
	executor, err := remotecommand.NewSPDYExecutor(K8s.Config, "POST", req.URL())
	if err != nil {
		log.ErrorWithErr("remotecommand pod error", err)
		return err
	}
	// 与 kubelet 建立 stream 连接
	if err = executor.Stream(remotecommand.StreamOptions{
		Stdout:            session,
		Stdin:             session,
		Stderr:            session,
		TerminalSizeQueue: session,
		Tty:               true,
	}); err != nil {
		log.ErrorWithErr("exec pod error", err)
		_, _ = session.Write([]byte("exec pod command failed," + err.Error()))
		// 标记关闭terminal
		session.Done()
	}
	return nil
}
