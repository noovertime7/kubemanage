package cmdb

import (
	"context"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/runtime"
)

type PermissionService interface {
	PagePermission(ctx context.Context, pager runtime.Pager) (dto.PageCMDBPermissionOut, error)
	GetPermission(ctx context.Context, in model.Permission) (*model.Permission, error)
	GetPermissionWithHosts(ctx context.Context, in model.Permission) (*model.Permission, error)
	DeletePermission(ctx context.Context, instanceID string) error
	DeletePermissions(ctx context.Context, instanceIDs []string) error
}

func NewPermissionService(factory dao.ShareDaoFactory) PermissionService {
	return &permissionService{factory: factory}
}

type permissionService struct {
	factory dao.ShareDaoFactory
}

func (p *permissionService) PagePermission(ctx context.Context, pager runtime.Pager) (dto.PageCMDBPermissionOut, error) {
	list, total, err := p.factory.CMDB().Permission().PageList(ctx, pager)
	if err != nil {
		return dto.PageCMDBPermissionOut{}, err
	}
	return dto.PageCMDBPermissionOut{Total: total, List: list, Page: pager.GetPage(), PageSize: pager.GetPageSize()}, nil
}

func (p *permissionService) GetPermission(ctx context.Context, in model.Permission) (*model.Permission, error) {
	return p.factory.CMDB().Permission().Find(ctx, in)
}

func (p *permissionService) GetPermissionWithHosts(ctx context.Context, in model.Permission) (*model.Permission, error) {
	return p.factory.CMDB().Permission().FindWithHosts(ctx, in)
}

func (p *permissionService) DeletePermission(ctx context.Context, instanceID string) error {
	return p.factory.CMDB().Permission().Delete(ctx, model.Permission{InstanceID: instanceID}, false)
}

func (p *permissionService) DeletePermissions(ctx context.Context, instanceIDs []string) error {
	for _, ins := range instanceIDs {
		if err := p.DeletePermission(ctx, ins); err != nil {
			return err
		}
	}
	return nil
}
