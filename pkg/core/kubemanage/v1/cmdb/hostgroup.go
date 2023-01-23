package cmdb

import (
	"context"
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/pkg/utils"

	"github.com/noovertime7/kubemanage/dto"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
)

type HostGroupService interface {
	GetHostGroupTree(ctx context.Context) ([]model.CMDBHostGroup, error)
	GetHostGroupList(ctx context.Context) ([]model.CMDBHostGroup, error)
	DeleteHostGroup(ctx context.Context, instanceID string) error
	CreateHostGroup(ctx context.Context, in *dto.HostGroupCreateOrUpdate) error
	CreateSonHostGroup(ctx context.Context, in *dto.HostGroupCreateOrUpdate) error
	UpdateHostGroup(ctx context.Context, in *dto.HostGroupCreateOrUpdate) error
}

func NewHostGroupService(factory dao.ShareDaoFactory) HostGroupService {
	return &hostGroupService{factory: factory}
}

var _ HostGroupService = &hostGroupService{}

type hostGroupService struct {
	factory dao.ShareDaoFactory
}

// DeleteHostGroup 删除主机组
func (h *hostGroupService) DeleteHostGroup(ctx context.Context, instanceID string) error {
	g, err := h.factory.CMDB().HostGroup().Find(ctx, model.CMDBHostGroup{InstanceID: instanceID})
	if err != nil {
		return err
	}
	if len(g.Hosts) > 0 {
		return fmt.Errorf("there are hosts in the host group, please delete or move the host first")
	}

	// 判断当前主机组下是否存在其他主机组
	list, err := h.factory.CMDB().HostGroup().FindList(ctx, model.CMDBHostGroup{ParentId: strconv.Itoa(int(g.Id))})
	if err != nil {
		return err
	}
	if len(list) > 0 {
		return fmt.Errorf("there are other hostgroup in this hostgroup, please delete or move this hostgroups first")
	}

	return h.factory.CMDB().HostGroup().Delete(ctx, g, false)
}

func (h *hostGroupService) FindHostGroupByInsID(ctx context.Context, instanceID string) (model.CMDBHostGroup, error) {
	g, err := h.factory.CMDB().HostGroup().Find(ctx, model.CMDBHostGroup{InstanceID: instanceID})
	if err != nil {
		return model.CMDBHostGroup{}, err
	}
	// This will never happen
	if utils.IsStrEmpty(g.InstanceID) {
		return model.CMDBHostGroup{}, fmt.Errorf("hostGroup not found ")
	}
	return g, nil
}

// CreateHostGroup 创建同级别主机组
func (h *hostGroupService) CreateHostGroup(ctx context.Context, in *dto.HostGroupCreateOrUpdate) error {
	g, err := h.FindHostGroupByInsID(ctx, in.InstanceID)
	if err != nil {
		return err
	}
	return h.factory.CMDB().HostGroup().Save(ctx, model.CMDBHostGroup{
		InstanceID: utils.GetSnowflakeID(),
		ParentId:   g.ParentId,
		GroupName:  in.GroupName,
		Sort:       in.Sort,
	})
}

// CreateSonHostGroup 创建子主机组
func (h *hostGroupService) CreateSonHostGroup(ctx context.Context, in *dto.HostGroupCreateOrUpdate) error {
	g, err := h.FindHostGroupByInsID(ctx, in.InstanceID)
	if err != nil {
		return err
	}
	return h.factory.CMDB().HostGroup().Save(ctx, model.CMDBHostGroup{
		InstanceID: utils.GetSnowflakeID(),
		ParentId:   strconv.Itoa(int(g.Id)),
		GroupName:  in.GroupName,
		Sort:       in.Sort,
	})
}

func (h *hostGroupService) UpdateHostGroup(ctx context.Context, in *dto.HostGroupCreateOrUpdate) error {
	return h.factory.CMDB().HostGroup().Updates(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("instanceID = ?", in.InstanceID)
	}, &model.CMDBHostGroup{
		GroupName: in.GroupName,
		Sort:      in.Sort,
	})
}

// GetHostGroupTree 获取主机组树形结构 通过 ParentID 与 id 的对应关系
func (h *hostGroupService) GetHostGroupTree(ctx context.Context) ([]model.CMDBHostGroup, error) {
	treeMap, err := h.getHostGroupTreeMap(ctx)
	if err != nil {
		return nil, err
	}
	hostGroups := treeMap["0"]
	for i := 0; i < len(hostGroups); i++ {
		if err := h.getDeptChildrenList(&hostGroups[i], treeMap); err != nil {
			return nil, err
		}
	}
	return hostGroups, nil
}

func (h *hostGroupService) GetHostGroupList(ctx context.Context) ([]model.CMDBHostGroup, error) {
	return h.factory.CMDB().HostGroup().FindList(ctx, model.CMDBHostGroup{})
}

func (h *hostGroupService) getDeptChildrenList(hostGroup *model.CMDBHostGroup, treeMap map[string][]model.CMDBHostGroup) (err error) {
	hostGroup.Children = treeMap[strconv.Itoa(int(hostGroup.Id))]
	for i := 0; i < len(hostGroup.Children); i++ {
		err = h.getDeptChildrenList(&hostGroup.Children[i], treeMap)
	}
	return err
}

func (h *hostGroupService) getHostGroupTreeMap(ctx context.Context) (treeMap map[string][]model.CMDBHostGroup, err error) {
	var in model.CMDBHostGroup
	treeMap = make(map[string][]model.CMDBHostGroup)
	hostGroups, err := h.factory.CMDB().HostGroup().FindListWithHosts(ctx, in)
	if err != nil {
		return nil, err
	}
	for _, v := range hostGroups {
		// 处理主机组名称,添加主机数，如 研发部门 (10)
		v.HostNum = int64(len(v.Hosts))
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}
