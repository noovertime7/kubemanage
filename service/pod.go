package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Pod pod

type pod struct{}

type PodsResp struct {
	Total int          `json:"total"`
	Items []coreV1.Pod `json:"items"`
}

// GetPods 获取pod列表支持、过滤、排序以及分页
func (p *pod) GetPods(filterName, namespace string, limit, page int) (podsResp *PodsResp, err error) {
	podlist, err := K8s.clientSet.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Error("获取Pod列表失败:", err.Error())
		return nil, errors.New("获取Pod列表失败")
	}
	//实例化dataSelector结构体，组装数据
	selectableData := &dataSelector{
		GenericDataList: p.toCells(podlist.Items),
		DataSelect: &DataSelectQuery{
			Filter:     &FilterQuery{Name: filterName},
			Paginatite: &PaginateQuery{limit, page},
		},
	}
	//先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	//排序、分页
	data := filtered.Sort().Paginate()
	//将dataCell类型转换为coreV1.Pod
	pods := p.FromCells(data.GenericDataList)
	return &PodsResp{
		total,
		pods,
	}, nil
}

//类型转换的方法 coreV1.pod => DataCell,DataCell => coreV1.pod
func (p *pod) toCells(pods []coreV1.Pod) []DataCell {
	cells := make([]DataCell, len(pods))
	for i := range pods {
		cells[i] = podCell(pods[i])
	}
	return cells
}

func (p *pod) FromCells(cells []DataCell) []coreV1.Pod {
	pods := make([]coreV1.Pod, len(cells))
	for i := range cells {
		pods[i] = coreV1.Pod(cells[i].(podCell))
	}
	return pods
}
