package cmdb

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/runtime"
)

type PermissionService interface {
	PagePermission(ctx context.Context, pager runtime.Pager) (dto.PageCMDBPermissionOut, error)
	GetPermission(ctx context.Context, in model.Permission) (*model.Permission, error)
	GetPermissionWithHosts(ctx context.Context, in model.Permission) (*model.Permission, error)
	DeletePermission(ctx context.Context, uuid uuid.UUID, instanceID string) error
	DeletePermissions(ctx context.Context, uuid uuid.UUID, instanceIDs []string) error
	CheckPermissionWithDB(ctx context.Context, uuid uuid.UUID, instanceID string) (bool, error)
	CheckPermissionWithCache(ctx context.Context, uuid uuid.UUID, instanceID string) (bool, error)
}

func NewPermissionService(factory dao.ShareDaoFactory) PermissionService {
	return &permissionService{factory: factory, PermissionCache: NewPermissionCache(factory)}
}

type permissionService struct {
	PermissionCache
	factory dao.ShareDaoFactory
}

func (p *permissionService) CheckPermissionWithDB(ctx context.Context, uuid uuid.UUID, instanceID string) (bool, error) {
	permission, err := p.GetPermissionWithHosts(ctx, model.Permission{UserUUID: uuid})
	if err != nil {
		return false, err
	}
	if permission.Id == 0 {
		logger.LG.Warn("CheckPermission: permission not found")
		return false, nil
	}
	if len(permission.Hosts) == 0 {
		logger.LG.Warn("CheckPermission: current user owned host is empty")
		return false, nil
	}
	// 构建该用户拥有的所有主机
	ownHosts := sets.NewString()
	now := time.Now()
	for _, host := range permission.Hosts {
		// 判断 当前时间在 在开始时间之后，在结束时间之前
		if now.After(permission.StartTime) && now.Before(permission.EndTime) {
			ownHosts.Insert(host.InstanceID)
		}
	}
	ok := ownHosts.Has(instanceID)
	return ok, nil
}

// CheckPermissionWithCache 检查主机是否有权限 返回false没有权限
func (p *permissionService) CheckPermissionWithCache(ctx context.Context, uuid uuid.UUID, instanceID string) (bool, error) {
	return p.CheckPermission(ctx, uuid, instanceID)
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

func (p *permissionService) DeletePermission(ctx context.Context, uuid uuid.UUID, instanceID string) error {
	p.Remove(ctx, uuid)
	return p.factory.CMDB().Permission().Delete(ctx, model.Permission{InstanceID: instanceID}, false)
}

func (p *permissionService) DeletePermissions(ctx context.Context, uuid uuid.UUID, instanceIDs []string) error {
	for _, ins := range instanceIDs {
		if err := p.DeletePermission(ctx, uuid, ins); err != nil {
			return err
		}
	}
	return nil
}
