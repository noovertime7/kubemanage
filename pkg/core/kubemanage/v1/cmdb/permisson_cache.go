package cmdb

import (
	"context"
	"sync"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg/logger"

	"github.com/noovertime7/kubemanage/dao"
	uuid "github.com/satori/go.uuid"
	"k8s.io/apimachinery/pkg/util/sets"
)

type PermissionCache interface {
	ReSync(ctx context.Context)
	Remove(ctx context.Context, uuid uuid.UUID)
	CheckPermission(ctx context.Context, uuid uuid.UUID, instanceID string) (bool, error)
}

type permissionCache struct {
	factory        dao.ShareDaoFactory
	lock           sync.RWMutex
	permissionsMap map[uuid.UUID]sets.String
}

func NewPermissionCache(factory dao.ShareDaoFactory) PermissionCache {
	return &permissionCache{factory: factory, permissionsMap: map[uuid.UUID]sets.String{}}
}

func (p *permissionCache) ReSync(ctx context.Context) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	list, err := p.factory.CMDB().Permission().FindList(ctx, model.Permission{})
	if err != nil {
		logger.LG.Warn(err.Error())
		return
	}
	for _, permission := range list {
		temp := sets.NewString()
		for _, item := range permission.Hosts {
			temp.Insert(item.InstanceID)
		}
		p.permissionsMap[permission.UserUUID] = temp
	}
}

func (p *permissionCache) Remove(ctx context.Context, uuid uuid.UUID) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	delete(p.permissionsMap, uuid)
}

func (p *permissionCache) CheckPermission(ctx context.Context, uuid uuid.UUID, instanceID string) (bool, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	list, ok := p.permissionsMap[uuid]
	if !ok {
		p.ReSync(ctx)
		return false, nil
	}
	return list.Has(instanceID), nil
}
