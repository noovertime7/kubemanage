package v1

import (
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/cmdb"
)

type CMDBGetter interface {
	CMDB() CMDBService
}

type CMDBService interface {
	HostGroup() cmdb.HostGroupService
}

type cmdbService struct {
	factory dao.ShareDaoFactory
}

func (c *cmdbService) HostGroup() cmdb.HostGroupService {
	return cmdb.NewHostGroupService(c.factory)
}

func NewCMDBService(factory dao.ShareDaoFactory) CMDBService {
	return &cmdbService{factory: factory}
}
