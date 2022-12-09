package kube

import (
	"context"
	"encoding/json"

	nwV1 "k8s.io/api/networking/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/noovertime7/kubemanage/dto/kubeDto"
)

var Ingress ingress

type ingress struct{}

type ingressResp struct {
	Total int            `json:"total"`
	Items []nwV1.Ingress `json:"items"`
}

type ingressNp struct {
	NameSpace  string `json:"namespace"`
	IngressNum int    `json:"ingres_num"`
}

func (i *ingress) toCells(ingress []nwV1.Ingress) []DataCell {
	cells := make([]DataCell, len(ingress))
	for i := range ingress {
		cells[i] = ingressCell(ingress[i])
	}
	return cells
}

func (i *ingress) FromCells(cells []DataCell) []nwV1.Ingress {
	ingress := make([]nwV1.Ingress, len(cells))
	for i := range cells {
		ingress[i] = nwV1.Ingress(cells[i].(ingressCell))
	}
	return ingress
}

func (i *ingress) CreateIngress(data *kubeDto.IngressCreteInput) error {
	//声明两个变量,用于后面组装数据
	var ingressRules []nwV1.IngressRule
	var httpIngressPaths []nwV1.HTTPIngressPath
	//将data中的数据组装成ingress对象
	ingress := &nwV1.Ingress{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.NameSpace,
			Labels:    data.Label,
		},
		Status: nwV1.IngressStatus{},
	}
	//第一层循环,将host组装成ingressRule对象,一个host对应一个ingressRule,每个ingressRule中包含一个host和多个path
	for key, value := range data.Hosts {
		ir := nwV1.IngressRule{
			Host: key,
			IngressRuleValue: nwV1.IngressRuleValue{
				HTTP: &nwV1.HTTPIngressRuleValue{
					Paths: nil,
				},
			},
		}
		//第二层循环，将path柱状成httpIngressPath对象
		for _, httpPath := range value {
			hip := nwV1.HTTPIngressPath{
				Path:     httpPath.Path,
				PathType: &httpPath.PathType,
				Backend: nwV1.IngressBackend{
					Service: &nwV1.IngressServiceBackend{
						Name: httpPath.ServiceName,
						Port: nwV1.ServiceBackendPort{
							Number: httpPath.ServicePort,
						},
					},
				},
			}
			//将每个hip对象组装成数组
			httpIngressPaths = append(httpIngressPaths, hip)
		}
		//给paths复制，前面置空了
		ir.IngressRuleValue.HTTP.Paths = httpIngressPaths
		//将每个ir对象组装成数组
		ingressRules = append(ingressRules, ir)
	}
	ingress.Spec.Rules = ingressRules
	//创建ingress
	if _, err := K8s.ClientSet.NetworkingV1().Ingresses(data.NameSpace).Create(context.TODO(), ingress, metaV1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

func (i *ingress) DeleteIngress(namespace, name string) error {
	return K8s.ClientSet.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (i *ingress) UpdateIngress(namespace, content string) error {
	ingress := &nwV1.Ingress{}
	if err := json.Unmarshal([]byte(content), ingress); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.NetworkingV1().Ingresses(namespace).Update(context.TODO(), ingress, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (i *ingress) GetIngressList(filterName, namespace string, limit, page int) (*ingressResp, error) {
	ingressList, err := K8s.ClientSet.NetworkingV1().Ingresses(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	selectableData := &dataSelector{
		GenericDataList: i.toCells(ingressList.Items),
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
	ingress := i.FromCells(data.GenericDataList)
	return &ingressResp{
		Total: total,
		Items: ingress,
	}, nil
}

func (i *ingress) GetIngressDetail(namespace, name string) (*nwV1.Ingress, error) {
	data, err := K8s.ClientSet.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i *ingress) GetIngressNp() ([]*ingressNp, error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var ingressnps []*ingressNp
	for _, namespace := range namespaceList.Items {
		ingress, err := K8s.ClientSet.NetworkingV1().Ingresses(namespace.Name).List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			return nil, err
		}
		ingressNp := &ingressNp{
			NameSpace:  namespace.Name,
			IngressNum: len(ingress.Items),
		}
		ingressnps = append(ingressnps, ingressNp)
	}
	return ingressnps, err
}
