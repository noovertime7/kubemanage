package v1

import (
	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/workflow"
	"k8s.io/client-go/kubernetes"
)

type CoreService interface {
	workflow.WorkFlowServiceGetter
}

type KubeManage struct {
	Cfg        config.Config
	Factory    dao.ShareDaoFactory
	clientSets map[string]*kubernetes.Clientset
}

func (c *KubeManage) WorkFlow() workflow.WorkFlowService {
	return workflow.NewWorkFlow(c)
}

func New(cfg config.Config, factory dao.ShareDaoFactory) CoreService {
	return &KubeManage{
		Cfg:        cfg,
		Factory:    factory,
		clientSets: map[string]*kubernetes.Clientset{},
	}
}
