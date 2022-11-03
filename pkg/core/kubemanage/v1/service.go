package v1

import (
	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/dao"
)

type CoreService interface {
	WorkFlowServiceGetter
	UserServiceGetter
}

type KubeManage struct {
	Cfg     config.Config
	Factory dao.ShareDaoFactory
}

func (c *KubeManage) WorkFlow() WorkFlowService {
	return NewWorkFlow(c)
}

func (c *KubeManage) User() UserService {
	return NewUserService(c)
}

func New(cfg config.Config, factory dao.ShareDaoFactory) CoreService {
	return &KubeManage{
		Cfg:     cfg,
		Factory: factory,
	}
}
