package cmdb

import (
	"context"
	"strconv"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
)

type HostGroupService interface {
	GetHostGroupTree(ctx context.Context) ([]model.CMDBHostGroup, error)
	GetHostGroupList(ctx context.Context) ([]model.CMDBHostGroup, error)
}

func NewHostGroupService(factory dao.ShareDaoFactory) HostGroupService {
	return &hostGroupService{factory: factory}
}

var _ HostGroupService = &hostGroupService{}

type hostGroupService struct {
	factory dao.ShareDaoFactory
}

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
