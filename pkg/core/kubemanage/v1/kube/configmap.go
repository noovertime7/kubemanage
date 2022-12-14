package kube

import (
	"context"
	"encoding/json"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Configmap configmap

type configmap struct{}

type ConfigmapResp struct {
	Total int                `json:"total"`
	Items []coreV1.ConfigMap `json:"items"`
}

type ConfigmapNp struct {
	NameSpace    string `json:"namespace"`
	ConfigmapNum int    `json:"configmap_num"`
}

func (d *configmap) toCells(Configmaps []coreV1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(Configmaps))
	for i := range Configmaps {
		cells[i] = configmapCell(Configmaps[i])
	}
	return cells
}

func (d *configmap) FromCells(cells []DataCell) []coreV1.ConfigMap {
	Configmaps := make([]coreV1.ConfigMap, len(cells))
	for i := range cells {
		Configmaps[i] = coreV1.ConfigMap(cells[i].(configmapCell))
	}
	return Configmaps
}

func (d *configmap) GetConfigmaps(filterName, namespace string, limit, page int) (*ConfigmapResp, error) {
	ConfigmapList, err := K8s.ClientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	selectableData := &dataSelector{
		GenericDataList: d.toCells(ConfigmapList.Items),
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
	Configmaps := d.FromCells(data.GenericDataList)
	return &ConfigmapResp{
		Total: total,
		Items: Configmaps,
	}, nil
}

func (d *configmap) GetConfigmapDetail(name, namespace string) (*coreV1.ConfigMap, error) {
	data, err := K8s.ClientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *configmap) DeleteConfigmap(name, namespace string) error {
	return K8s.ClientSet.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (d *configmap) UpdateConfigmap(content, namespace string) error {
	var Configmap = &coreV1.ConfigMap{}
	if err := json.Unmarshal([]byte(content), Configmap); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.CoreV1().ConfigMaps(namespace).Update(context.TODO(), Configmap, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}
