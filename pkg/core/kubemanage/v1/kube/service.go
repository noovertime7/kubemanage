package kube

import (
	"context"
	"encoding/json"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var Service service

type service struct{}

type serviceResp struct {
	Total int              `json:"total"`
	Items []coreV1.Service `json:"items"`
}

type serviceNp struct {
	NameSpace  string `json:"namespace"`
	ServiceNum int    `json:"service_num"`
}

func (s *service) toCells(services []coreV1.Service) []DataCell {
	cells := make([]DataCell, len(services))
	for i := range services {
		cells[i] = serviceCell(services[i])
	}
	return cells
}

func (s *service) FromCells(cells []DataCell) []coreV1.Service {
	services := make([]coreV1.Service, len(cells))
	for i := range cells {
		services[i] = coreV1.Service(cells[i].(serviceCell))
	}
	return services
}

func (s *service) CreateService(data *kubernetes.ServiceCreateInput) error {
	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.NameSpace,
			Labels:    data.Label,
		},
		Spec: coreV1.ServiceSpec{
			Type: coreV1.ServiceType(data.Type),
			Ports: []coreV1.ServicePort{
				{
					Name:     "http",
					Port:     data.Port,
					Protocol: "TCP",
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			Selector: data.Label,
		},
		Status: coreV1.ServiceStatus{},
	}
	if data.NodePort != 0 && data.Type == "NodePort" {
		service.Spec.Ports[0].NodePort = data.NodePort
	}
	//创建service
	if _, err := K8s.ClientSet.CoreV1().Services(data.NameSpace).Create(context.TODO(), service, metaV1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteService(name, namespace string) error {
	return K8s.ClientSet.CoreV1().Services(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (s *service) UpdateService(namespace, content string) error {
	var Service = &coreV1.Service{}
	if err := json.Unmarshal([]byte(content), Service); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.CoreV1().Services(namespace).Update(context.TODO(), Service, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetServiceList(filterName, namespace string, limit, page int) (*serviceResp, error) {
	ServiceList, err := K8s.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logger.Error("获取Pod列表失败:", err.Error())
		return nil, errors.New("获取Pod列表失败")
	}
	//实例化dataSelector结构体，组装数据
	selectableData := &dataSelector{
		GenericDataList: s.toCells(ServiceList.Items),
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
	Services := s.FromCells(data.GenericDataList)
	return &serviceResp{
		total,
		Services,
	}, nil
}

func (s *service) GetServiceDetail(name, namespace string) (*coreV1.Service, error) {
	data, err := K8s.ClientSet.CoreV1().Services(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *service) GetServiceNp() ([]*serviceNp, error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var services []*serviceNp
	for _, namespace := range namespaceList.Items {
		serviceList, err := K8s.ClientSet.CoreV1().Services(namespace.Name).List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			return nil, err
		}
		//组装数据
		ServiceNp := &serviceNp{
			NameSpace:  namespace.Name,
			ServiceNum: len(serviceList.Items),
		}
		services = append(services, ServiceNp)
	}
	return services, nil
}
