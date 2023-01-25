package cmdb

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/cmdb/webshell"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/pkg/utils"
	"github.com/noovertime7/kubemanage/runtime"
	"github.com/noovertime7/kubemanage/runtime/queue"
)

type HostService interface {
	CreateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error
	UpdateHost(ctx context.Context, uuid uuid.UUID, in *dto.CMDBHostCreateInput) error
	GetHostListWithGroupName(ctx context.Context, uuid uuid.UUID, search *model.CMDBHost) ([]*model.CMDBHost, error)
	PageHost(ctx context.Context, uuid uuid.UUID, groupID uint, pager runtime.Pager) (dto.PageCMDBHostOut, error)
	DeleteHost(ctx context.Context, uuid uuid.UUID, instanceID string) error
	DeleteHosts(ctx context.Context, uuid uuid.UUID, instanceIDs []string) error
	StartHostCheck() error
	WebShell(ctx *gin.Context, instanceID string, cols, rows int) error
}

func NewHostService(factory dao.ShareDaoFactory, q queue.Queue) HostService {
	return &hostService{factory: factory, queue: q, permission: NewPermissionService(factory)}
}

type hostService struct {
	queue      queue.Queue
	permission PermissionService
	factory    dao.ShareDaoFactory
}

func (h *hostService) CreateHost(ctx context.Context, in *dto.CMDBHostCreateInput) error {
	var (
		enHostPassword string
		enPrivateKey   string
		err            error
	)
	//  查询是否ip重复添加
	tempDB := model.CMDBHost{Address: in.Address}
	_, err = h.factory.CMDB().Host().Find(ctx, tempDB)
	if err != nil {
		if utils.GormExist(err) {
			return fmt.Errorf("存在相同IP地址主机，请重新填写")
		}
	}

	if in.HostPassword != "" {
		enHostPassword, err = utils.Encrypt([]byte(in.HostPassword))
		if err != nil {
			return err
		}
	}

	if in.PrivateKey != "" {
		enPrivateKey, err = utils.Encrypt([]byte(in.PrivateKey))
		if err != nil {
			return err
		}
	}

	hostDB := &model.CMDBHost{
		InstanceID:      utils.GetSnowflakeID(),
		UseSecret:       in.UseSecret,
		Name:            in.Name,
		Address:         in.Address,
		Port:            in.Port,
		HostUserName:    in.HostUserName,
		HostPassword:    enHostPassword,
		PrivateKey:      enPrivateKey,
		SecretID:        in.SecretID,
		Protocol:        in.Protocol,
		SecretType:      in.SecretType,
		Status:          1,
		CMDBHostGroupID: in.CMDBHostGroupID,
	}
	return h.factory.CMDB().Host().Save(ctx, hostDB)
}

// UpdateHost TODO 应该判断是否有权限更新
func (h *hostService) UpdateHost(ctx context.Context, uuid uuid.UUID, in *dto.CMDBHostCreateInput) error {
	if utils.IsStrEmpty(in.InstanceID) {
		return fmt.Errorf("instance id is empty")
	}

	// 检查当前用户是否拥有这个主机的权限
	ok, err := h.permission.CheckPermissionWithCache(ctx, uuid, in.InstanceID)
	if err != nil {
		return err
	}
	// This will never happen
	if !ok {
		return fmt.Errorf("403 当前用户没有更新这台主机的权限")
	}

	host, err := h.factory.CMDB().Host().Find(ctx, model.CMDBHost{InstanceID: in.InstanceID})
	if err != nil {
		return err
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

	// 表示密码发生变化
	if in.HostPassword != host.HostPassword && in.HostPassword != "" {
		enHostPassword, err := utils.Encrypt([]byte(in.HostPassword))
		if err != nil {
			return err
		}
		hostDB.HostPassword = enHostPassword
	}

	// 秘钥发生变化
	if in.PrivateKey != host.PrivateKey && in.PrivateKey != "" {
		enPrivateKey, err := utils.Encrypt([]byte(in.PrivateKey))
		if err != nil {
			return err
		}
		hostDB.PrivateKey = enPrivateKey
	}

	return h.factory.CMDB().Host().Updates(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("instanceID = ?", in.InstanceID)
	}, hostDB)
}

func (h *hostService) PageHost(ctx context.Context, uuid uuid.UUID, groupID uint, pager runtime.Pager) (dto.PageCMDBHostOut, error) {
	list, total, err := h.factory.CMDB().Host().PageList(ctx, groupID, pager)
	// 进行权限过滤
	list, err = h.buildPermissionFitter(ctx, uuid, list)
	if err != nil {
		return dto.PageCMDBHostOut{}, err
	}
	// 补充主机组名
	newList, err := h.buildHostGroupName(ctx, list)
	if err != nil {
		return dto.PageCMDBHostOut{}, err
	}
	total = int64(len(newList))
	return dto.PageCMDBHostOut{Total: total, List: newList}, nil
}

func (h *hostService) buildPermissionFitter(ctx context.Context, uuid uuid.UUID, in []*model.CMDBHost) ([]*model.CMDBHost, error) {
	var newList = make([]*model.CMDBHost, 0)
	permission, err := h.permission.GetPermissionWithHosts(ctx, model.Permission{UserUUID: uuid})
	if err != nil {
		return nil, err
	}
	if permission.Id == 0 {
		logger.LG.Warn("buildPermissionFitter: permission not found")
		return newList, nil
	}
	if len(permission.Hosts) == 0 {
		logger.LG.Warn("buildPermissionFitter: current user owned host is empty")
		return newList, nil
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

	for _, realHost := range in {
		if ownHosts.Has(realHost.InstanceID) {
			newList = append(newList, realHost)
		}
	}
	return newList, err
}

func (h *hostService) buildHostGroupName(ctx context.Context, in []*model.CMDBHost) ([]*model.CMDBHost, error) {
	var newList = make([]*model.CMDBHost, 0)
	if len(in) == 0 {
		return newList, nil
	}
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

func (h *hostService) DeleteHost(ctx context.Context, uuid uuid.UUID, instanceID string) error {
	// 检查当前用户是否拥有这个主机的权限
	ok, err := h.permission.CheckPermissionWithCache(ctx, uuid, instanceID)
	if err != nil {
		return err
	}
	// This will never happen
	if !ok {
		return fmt.Errorf("403 当前用户没有更新这台主机的权限")
	}
	return h.factory.CMDB().Host().Delete(ctx, model.CMDBHost{InstanceID: instanceID}, false)
}

func (h *hostService) DeleteHosts(ctx context.Context, uuid uuid.UUID, instanceIDs []string) error {
	for _, ins := range instanceIDs {
		if err := h.DeleteHost(ctx, uuid, ins); err != nil {
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

func (h *hostService) GetHostListWithGroupName(ctx context.Context, uuid uuid.UUID, search *model.CMDBHost) ([]*model.CMDBHost, error) {
	if search == nil {
		search = &model.CMDBHost{}
	}
	list, err := h.getHostList(ctx, *search)
	if err != nil {
		return nil, err
	}

	fitterList, err := h.buildPermissionFitter(ctx, uuid, list)
	if err != nil {
		return nil, err
	}

	groupNameList, err := h.buildHostGroupName(ctx, fitterList)
	if err != nil {
		return nil, err
	}
	return groupNameList, nil

}

func (h *hostService) WebShell(ctx *gin.Context, instanceID string, cols, rows int) error {
	// TODO 需要优化
	//获取主机信息
	host, err := h.factory.CMDB().Host().Find(ctx, model.CMDBHost{InstanceID: instanceID})
	if err != nil {
		return err
	}
	info, err := h.buildHostConnectionInfo(ctx, host)
	if err != nil {
		return err
	}
	if err := webshell.SSHWsHandler.SetUp(ctx.Writer, ctx.Request, info, cols, rows); err != nil {
		return err
	}
	return nil
}

// 构建主机连接信息
func (h *hostService) buildHostConnectionInfo(ctx context.Context, host model.CMDBHost) (*webshell.HostConnectionInfo, error) {
	var (
		info          *webshell.HostConnectionInfo
		dePassword    []byte
		dePrivateKey  []byte
		usePrivateKey bool
		err           error
	)
	if host.UseSecret == 1 {
		secret, err := h.factory.CMDB().Secret().Find(ctx, model.CMDBSecret{Id: host.SecretID})
		if err != nil {
			return nil, err
		}
		// 密码解密
		if secret.SecretType == 1 {
			dePassword, err = utils.Decrypt(secret.HostPassword)
			if err != nil {
				return nil, err
			}
			// 秘钥认证方式
		} else if secret.SecretType == 2 {
			usePrivateKey = true
			dePrivateKey, err = utils.Decrypt(secret.PrivateKey)
			if err != nil {
				return nil, err
			}
		}

		info = &webshell.HostConnectionInfo{
			Address:       host.Address,
			Port:          host.Port,
			UserName:      secret.HostUserName,
			Password:      string(dePassword),
			UsePrivateKey: usePrivateKey,
			PrivateKey:    string(dePrivateKey),
		}
	} else {
		// 代表用户名密码登录
		if host.SecretType == 1 {
			dePassword, err = utils.Decrypt(host.HostPassword)
			if err != nil {
				return nil, err
			}
			// 秘钥认证方式
		} else if host.SecretType == 2 {
			usePrivateKey = true
			dePrivateKey, err = utils.Decrypt(host.PrivateKey)
			if err != nil {
				return nil, err
			}
		}
		info = &webshell.HostConnectionInfo{
			Address:       host.Address,
			Port:          host.Port,
			UserName:      host.HostUserName,
			Password:      string(dePassword),
			UsePrivateKey: usePrivateKey,
			PrivateKey:    string(dePrivateKey),
		}
	}
	return info, nil
}
