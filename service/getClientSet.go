package service

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClient(configPath string) (*kubernetes.Clientset, error) {
	var err error
	var config *rest.Config
	if config, err = clientcmd.BuildConfigFromFlags("", configPath); err != nil {
		return nil, err
	}
	// 创建 clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet, nil
}
