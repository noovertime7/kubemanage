package v1

import (
	"github.com/noovertime7/kubemanage/cmd/app/config"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/pkg/logger"
)

type CoreService interface {
	WorkFlowServiceGetter
	UserServiceGetter
	MenuGetter
	CasbinServiceGetter
	AuthorityGetter
	OperationServiceGetter
	CloudGetter
}

func New(cfg *config.Config, factory dao.ShareDaoFactory) CoreService {
	return &KubeManage{
		Cfg:     cfg,
		Factory: factory,
	}
}

type Logger interface {
	logger.Logger
}

type KubeManage struct {
	Cfg     *config.Config
	Factory dao.ShareDaoFactory
}

func (c *KubeManage) WorkFlow() WorkFlowService {
	return NewWorkFlow(c)
}

func (c *KubeManage) User() UserService {
	return NewUserService(c)
}

func (c *KubeManage) Menu() MenuService {
	return NewMenuService(c)
}

func (c *KubeManage) CasbinService() CasbinService {
	return NewCasbinService(c)
}

func (c *KubeManage) Authority() Authority {
	return NewAuthority(c)
}

func (c *KubeManage) Operation() OperationService {
	return NewOperationService(c)
}

func (c *KubeManage) Cloud() CloudInterface {
	return newCloud(c)
}
