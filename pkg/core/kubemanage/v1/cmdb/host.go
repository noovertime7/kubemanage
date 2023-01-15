package cmdb

import (
	"context"
	"fmt"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/utils"
	"github.com/noovertime7/kubemanage/runtime"
	"github.com/noovertime7/kubemanage/runtime/queue"
)

type HostService interface {
	CreateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error
	UpdateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error
	PageHost(ctx context.Context, pager runtime.Pager) (dto.PageCMDBHostOut, error)
	DeleteHost(ctx context.Context, instanceID string) error
	DeleteHosts(ctx context.Context, instanceIDs []string) error
	StartHostCheck() error
}

func NewHostService(factory dao.ShareDaoFactory, q queue.Queue) HostService {
	return &hostService{factory: factory, queue: q}
}

type hostService struct {
	queue   queue.Queue
	factory dao.ShareDaoFactory
}

func (h *hostService) CreateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error {
	//  查询是否ip重复添加
	tempDB := model.CMDBHost{Address: in.Address}
	temp, err := h.factory.CMDB().Host().Find(ctx, tempDB)
	if err != nil {
		return err
	}
	if temp.Id != 0 {
		return fmt.Errorf("主机已添加")
	}
	hostDB := &model.CMDBHost{
		InstanceID:      utils.GetSnowflakeID(),
		Address:         in.Address,
		Port:            in.Port,
		HostUserName:    in.HostUserName,
		HostPassword:    in.HostPassword,
		PrivateKey:      in.PrivateKey,
		SecretID:        in.SecretID,
		Status:          1,
		CMDBHostGroupID: in.CMDBHostGroupID,
	}
	return h.factory.CMDB().Host().Save(ctx, hostDB)
}

func (h *hostService) UpdateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error {
	hostDB := &model.CMDBHost{
		InstanceID:      in.InstanceID,
		Address:         in.Address,
		Port:            in.Port,
		HostUserName:    in.HostUserName,
		HostPassword:    in.HostPassword,
		PrivateKey:      in.PrivateKey,
		SecretID:        in.SecretID,
		Status:          1,
		CMDBHostGroupID: in.CMDBHostGroupID,
	}
	return h.factory.CMDB().Host().Updates(ctx, hostDB)
}

func (h *hostService) PageHost(ctx context.Context, pager runtime.Pager) (dto.PageCMDBHostOut, error) {
	list, total, err := h.factory.CMDB().Host().PageList(ctx, pager)
	if err != nil {
		return dto.PageCMDBHostOut{}, err
	}
	return dto.PageCMDBHostOut{Total: total, List: list}, nil
}

func (h *hostService) DeleteHost(ctx context.Context, instanceID string) error {
	return h.factory.CMDB().Host().Delete(ctx, model.CMDBHost{InstanceID: instanceID}, false)
}

func (h *hostService) DeleteHosts(ctx context.Context, instanceIDs []string) error {
	for _, ins := range instanceIDs {
		if err := h.DeleteHost(ctx, ins); err != nil {
			return err
		}
	}
	return nil
}

func (h *hostService) GetHostList(ctx context.Context, search model.CMDBHost) ([]model.CMDBHost, error) {
	data, err := h.factory.CMDB().Host().FindList(ctx, search)
	if err != nil {
		return nil, err
	}
	return data, nil
}
