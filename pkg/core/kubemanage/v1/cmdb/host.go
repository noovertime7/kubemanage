package cmdb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/cmdb/webshell"

	"gorm.io/gorm"

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
	GetHostListWithGroupName(ctx context.Context, search *model.CMDBHost) ([]*model.CMDBHost, error)
	PageHost(ctx context.Context, groupID uint, pager runtime.Pager) (dto.PageCMDBHostOut, error)
	DeleteHost(ctx context.Context, instanceID string) error
	DeleteHosts(ctx context.Context, instanceIDs []string) error
	StartHostCheck() error
	WebShell(w http.ResponseWriter, r *http.Request, cols, rows int) error
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
		UseSecret:       in.UseSecret,
		Name:            in.Name,
		Address:         in.Address,
		Port:            in.Port,
		HostUserName:    in.HostUserName,
		HostPassword:    in.HostPassword,
		PrivateKey:      in.PrivateKey,
		SecretID:        in.SecretID,
		Protocol:        in.Protocol,
		SecretType:      in.SecretType,
		Status:          1,
		CMDBHostGroupID: in.CMDBHostGroupID,
	}
	return h.factory.CMDB().Host().Save(ctx, hostDB)
}

func (h *hostService) UpdateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error {
	if utils.IsStrEmpty(in.InstanceID) {
		return fmt.Errorf("instance id is empty")
	}
	hostDB := &model.CMDBHost{
		InstanceID:      in.InstanceID,
		Name:            in.Name,
		UseSecret:       in.UseSecret,
		Address:         in.Address,
		Port:            in.Port,
		HostUserName:    in.HostUserName,
		HostPassword:    in.HostPassword,
		PrivateKey:      in.PrivateKey,
		Protocol:        in.Protocol,
		SecretType:      in.SecretType,
		SecretID:        in.SecretID,
		CMDBHostGroupID: in.CMDBHostGroupID,
	}
	return h.factory.CMDB().Host().Updates(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("instanceID = ?", in.InstanceID)
	}, hostDB)
}

func (h *hostService) PageHost(ctx context.Context, groupID uint, pager runtime.Pager) (dto.PageCMDBHostOut, error) {
	list, total, err := h.factory.CMDB().Host().PageList(ctx, groupID, pager)
	newList, err := h.buildHostGroupName(ctx, list)
	if err != nil {
		return dto.PageCMDBHostOut{}, err
	}
	return dto.PageCMDBHostOut{Total: total, List: newList}, nil
}

func (h *hostService) buildHostGroupName(ctx context.Context, in []*model.CMDBHost) ([]*model.CMDBHost, error) {
	var newList []*model.CMDBHost
	for _, host := range in {
		group, err := h.factory.CMDB().HostGroup().Find(ctx, model.CMDBHostGroup{Id: host.CMDBHostGroupID})
		if err != nil {
			return nil, err
		}
		host.GroupName = group.GroupName
		newList = append(newList, host)
	}
	return newList, nil
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

func (h *hostService) getHostList(ctx context.Context, search model.CMDBHost) ([]*model.CMDBHost, error) {
	data, err := h.factory.CMDB().Host().FindList(ctx, search)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *hostService) GetHostListWithGroupName(ctx context.Context, search *model.CMDBHost) ([]*model.CMDBHost, error) {
	if search == nil {
		search = &model.CMDBHost{}
	}
	list, err := h.getHostList(ctx, *search)
	if err != nil {
		return nil, err
	}
	groupNameList, err := h.buildHostGroupName(ctx, list)
	if err != nil {
		return nil, err
	}
	return groupNameList, nil

}

func (h *hostService) WebShell(w http.ResponseWriter, r *http.Request, cols, rows int) error {
	// TODO 需要优化
	if err := webshell.SSHWsHandler.SetUp(w, r, cols, rows); err != nil {
		return err
	}
	return nil
}
