package kube

import (
	"sort"
	"strings"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	nwV1 "k8s.io/api/networking/v1"
)

// 用于封装排序、过滤、分页方法
type dataSelector struct {
	GenericDataList []DataCell
	DataSelect      *DataSelectQuery
}

// DataCell 用于各种资源list的类型转化，转换后可以使用data-selector的自定义排序、过滤、分页方法
type DataCell interface {
	GetCreation() time.Time
	GetName() string
}

// DataSelectQuery 定义过滤和分页的结构体
type DataSelectQuery struct {
	Filter     *FilterQuery
	Paginatite *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

//  实现自定义结构的排序。需要重新len、swap、less方法

// Len 方法用于获取数组的长度
func (d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

// Swap swap方法用于数据比较大小之后的位置变更
func (d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// Less 比较大小
func (d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a)
}

func (d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

// Filter 用于过滤数据，比较数据的name属性，若包含，则返回
func (d *dataSelector) Filter() *dataSelector {
	//判断是否为空，若为空，则返回所有数据
	if d.DataSelect.Filter.Name == "" {
		return d
	}
	//若不为空，则按照入参name进行过滤,若name包含，则吧数据放进新数组中返回出去
	var filtered []DataCell
	for _, v := range d.GenericDataList {
		objName := v.GetName()
		if !strings.Contains(objName, d.DataSelect.Filter.Name) {
			continue
		}
		filtered = append(filtered, v)
	}
	d.GenericDataList = filtered
	return d
}

// Paginate 方法用于数组的分页，根据limit与page的传参取一定范围内的数据返回
func (d *dataSelector) Paginate() *dataSelector {
	//根据limit和page的入参定义快捷变量
	limit := d.DataSelect.Paginatite.Limit
	page := d.DataSelect.Paginatite.Page
	//检验参数的合法性
	if limit <= 0 || page <= 0 {
		return d
	}
	//定义取数范围需要的start index和end-index
	startIndex := limit * (page - 1)
	endIndex := limit*page - 1
	if startIndex > len(d.GenericDataList) {
		startIndex = 0
	}
	//处理end index
	if endIndex > len(d.GenericDataList) {
		endIndex = len(d.GenericDataList)
	}
	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d
}

// 定义podCell，重写GetCreation与getName方法后，可以实现数据转换
type podCell coreV1.Pod

// GetCreation 重写dataCell接口的两个方法
func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}

type deploymentCell appsV1.Deployment

func (d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d deploymentCell) GetName() string {
	return d.Name
}

type daemonSetCell appsV1.DaemonSet

func (d daemonSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d daemonSetCell) GetName() string {
	return d.Name
}

type statefulSetCell appsV1.StatefulSet

func (d statefulSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d statefulSetCell) GetName() string {
	return d.Name
}

type nodeCell coreV1.Node

func (d nodeCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d nodeCell) GetName() string {
	return d.Name
}

type namespaceCell coreV1.Namespace

func (d namespaceCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d namespaceCell) GetName() string {
	return d.Name
}

type persistentvolumesCell coreV1.PersistentVolume

func (d persistentvolumesCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d persistentvolumesCell) GetName() string {
	return d.Name
}

type serviceCell coreV1.Service

func (d serviceCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d serviceCell) GetName() string {
	return d.Name
}

type ingressCell nwV1.Ingress

func (d ingressCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d ingressCell) GetName() string {
	return d.Name
}

type configmapCell coreV1.ConfigMap

func (d configmapCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d configmapCell) GetName() string {
	return d.Name
}

type persistentVolumeClaimCell coreV1.PersistentVolumeClaim

func (d persistentVolumeClaimCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d persistentVolumeClaimCell) GetName() string {
	return d.Name
}

type secretCell coreV1.Secret

func (d secretCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d secretCell) GetName() string {
	return d.Name
}
