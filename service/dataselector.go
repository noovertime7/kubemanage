package service

import (
	corev1 "k8s.io/api/core/v1"
	"sort"
	"strings"
	"time"
)

//用于封装排序、过滤、分页方法
type dataSelector struct {
	GenericDataList []DataCell
	DataSelect      *DataSelectQuery
}

//DataCell 用于各种资源list的类型转化，转换后可以使用data-selector的自定义排序、过滤、分页方法
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

//Len 方法用于获取数组的长度
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

//Paginate 方法用于数组的分页，根据limit与page的传参取一定范围内的数据返回
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
	//处理end index
	if endIndex > len(d.GenericDataList) {
		endIndex = len(d.GenericDataList)
	}
	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d
}

//定义podCell，重写GetCreation与getName方法后，可以实现数据转换
type podCell corev1.Pod

// GetCreation 重写dataCell接口的两个方法
func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}