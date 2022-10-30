package kube

import (
	"context"
	"encoding/json"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StatefulSet statefulSet

type statefulSet struct{}

type statefulSetResp struct {
	Total int                  `json:"total"`
	Items []appsV1.StatefulSet `json:"items"`
}

type StatefulSetNp struct {
	NameSpace    string `json:"namespace"`
	DaemonSetNum int    `json:"daemonset_num"`
}

func (d *statefulSet) toCells(statefulSets []appsV1.StatefulSet) []DataCell {
	cells := make([]DataCell, len(statefulSets))
	for i := range statefulSets {
		cells[i] = statefulSetCell(statefulSets[i])
	}
	return cells
}

func (d *statefulSet) FromCells(cells []DataCell) []appsV1.StatefulSet {
	statefulSets := make([]appsV1.StatefulSet, len(cells))
	for i := range cells {
		statefulSets[i] = appsV1.StatefulSet(cells[i].(statefulSetCell))
	}
	return statefulSets
}

func (d *statefulSet) GetStatefulSets(filterName, namespace string, limit, page int) (*statefulSetResp, error) {
	statefulSetList, err := K8s.ClientSet.AppsV1().StatefulSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	selectableData := &dataSelector{
		GenericDataList: d.toCells(statefulSetList.Items),
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
	statefulSets := d.FromCells(data.GenericDataList)
	return &statefulSetResp{
		Total: total,
		Items: statefulSets,
	}, nil
}

func (d *statefulSet) GetStatefulSetDetail(name, namespace string) (*appsV1.StatefulSet, error) {
	data, err := K8s.ClientSet.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *statefulSet) DeleteStatefulSet(name, namespace string) error {
	return K8s.ClientSet.AppsV1().StatefulSets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (d *statefulSet) UpdateStatefulSet(content, namespace string) error {
	var statefulSet = &appsV1.StatefulSet{}
	if err := json.Unmarshal([]byte(content), statefulSet); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.AppsV1().StatefulSets(namespace).Update(context.TODO(), statefulSet, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}
