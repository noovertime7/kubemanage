package v1

import (
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/cmdb"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/runtime/checker"
	"github.com/noovertime7/kubemanage/runtime/queue"
)

var hostLocalQueue = queue.NewQueue()

type CMDBGetter interface {
	CMDB() CMDBService
}

type CMDBService interface {
	Permission() cmdb.PermissionService
	HostGroup() cmdb.HostGroupService
	Host() cmdb.HostService
	Secret() cmdb.SecretService
	StartChecker()
}

type cmdbService struct {
	factory dao.ShareDaoFactory
}

func (c *cmdbService) HostGroup() cmdb.HostGroupService {
	return cmdb.NewHostGroupService(c.factory)
}

func (c *cmdbService) Host() cmdb.HostService {
	return cmdb.NewHostService(c.factory, hostLocalQueue)
}

func (c *cmdbService) Secret() cmdb.SecretService {
	return cmdb.NewSecretService(c.factory)
}

func (c *cmdbService) Permission() cmdb.PermissionService {
	return cmdb.NewPermissionService(c.factory)
}

func (c *cmdbService) StartChecker() {
	telnetChecker := cmdb.NewTelnetHandler(c.factory, hostLocalQueue, logger.New(logger.LG))
	factory := checker.NewSharedCheckerFactory()
	// 向factory中注册 checker
	_ = factory.CheckerFor(telnetChecker)
	factory.Start()
}

func NewCMDBService(factory dao.ShareDaoFactory) CMDBService {
	return &cmdbService{factory: factory}
}
