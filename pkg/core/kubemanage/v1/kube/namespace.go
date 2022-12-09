package kube

import (
	"context"
	"github.com/pkg/errors"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NameSpace namespace

type namespace struct{}

type NameSpaceResp struct {
	Total int                `json:"total"`
	Items []coreV1.Namespace `json:"items"`
}

func (n *namespace) toCells(nodes []coreV1.Namespace) []DataCell {
	cells := make([]DataCell, len(nodes))
	for i := range nodes {
		cells[i] = namespaceCell(nodes[i])
	}
	return cells
}

func (n *namespace) FromCells(cells []DataCell) []coreV1.Namespace {
	nodes := make([]coreV1.Namespace, len(cells))
	for i := range cells {
		nodes[i] = coreV1.Namespace(cells[i].(namespaceCell))
	}
	return nodes
}

func (n *namespace) GetNameSpaces(filterName string, limit, page int) (nodesResp *NameSpaceResp, err error) {
	NamespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, errors.New("获取Pod列表失败")
	}
	//实例化dataSelector结构体，组装数据
	selectableData := &dataSelector{
		GenericDataList: n.toCells(NamespaceList.Items),
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
	namespaces := n.FromCells(data.GenericDataList)
	return &NameSpaceResp{
		total,
		namespaces,
	}, nil
}

// GetNameSpacesDetail 获取Node详情
func (n *namespace) GetNameSpacesDetail(Name string) (*coreV1.Namespace, error) {
	namespacesRes, err := K8s.ClientSet.CoreV1().Namespaces().Get(context.TODO(), Name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return namespacesRes, nil
}

func (n *namespace) CreateNameSpace(name string) error {
	ns := &coreV1.Namespace{
		TypeMeta: metaV1.TypeMeta{},
		ObjectMeta: metaV1.ObjectMeta{
			Name: name,
		},
		Spec:   coreV1.NamespaceSpec{},
		Status: coreV1.NamespaceStatus{},
	}
	if _, err := K8s.ClientSet.CoreV1().Namespaces().Create(context.TODO(), ns, metaV1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

func (n *namespace) DeleteNameSpace(name string) error {
	return K8s.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), name, metaV1.DeleteOptions{})
}
