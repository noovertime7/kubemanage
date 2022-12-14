package kube

import (
	"context"
	"encoding/json"

	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DaemonSet daemonSet

type daemonSet struct {
}

type DaemonSetResp struct {
	Total int                `json:"total"`
	Items []appsV1.DaemonSet `json:"items"`
}

type DaemonSetNp struct {
	NameSpace    string `json:"namespace"`
	DaemonSetNum int    `json:"daemonset_num"`
}

func (d *daemonSet) toCells(daemonsets []appsV1.DaemonSet) []DataCell {
	cells := make([]DataCell, len(daemonsets))
	for i := range daemonsets {
		cells[i] = daemonSetCell(daemonsets[i])
	}
	return cells
}

func (d *daemonSet) FromCells(cells []DataCell) []appsV1.DaemonSet {
	daemonSets := make([]appsV1.DaemonSet, len(cells))
	for i := range cells {
		daemonSets[i] = appsV1.DaemonSet(cells[i].(daemonSetCell))
	}
	return daemonSets
}

func (d *daemonSet) GetDaemonSets(filterName, namespace string, limit, page int) (*DaemonSetResp, error) {
	daemonSetList, err := K8s.ClientSet.AppsV1().DaemonSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	selectableData := &dataSelector{
		GenericDataList: d.toCells(daemonSetList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginatite: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filterd := selectableData.Filter()
	total := len(filterd.GenericDataList)
	data := filterd.Sort().Paginate()
	daemonSets := d.FromCells(data.GenericDataList)
	return &DaemonSetResp{
		Total: total,
		Items: daemonSets,
	}, nil
}

func (d *daemonSet) GetDaemonSetDetail(name, namespace string) (*appsV1.DaemonSet, error) {
	data, err := K8s.ClientSet.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *daemonSet) DeleteDaemonSet(name, namespace string) error {
	return K8s.ClientSet.AppsV1().DaemonSets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (d *daemonSet) UpdateDaemonSet(content, namespace string) error {
	var daemonset = &appsV1.DaemonSet{}
	if err := json.Unmarshal([]byte(content), daemonset); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.AppsV1().DaemonSets(namespace).Update(context.TODO(), daemonset, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}
