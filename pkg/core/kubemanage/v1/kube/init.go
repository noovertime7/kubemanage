package kube

import (
	"flag"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	Config    *rest.Config
	ClientSet *kubernetes.Clientset
}

func (k *k8s) Init() error {
	var err error
	var config *rest.Config
	var kubeConfig *string

	if home := homeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 使用 ServiceAccount 创建集群配置（InCluster模式）
	if config, err = rest.InClusterConfig(); err != nil {
		// 使用 KubeConfig 文件创建集群配置
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig); err != nil {
			return err
		}
	}

	// 创建 clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	log := logger.New()
	log.Info("获取k8s clientSet 成功")
	k.ClientSet = clientSet
	k.Config = config
	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
