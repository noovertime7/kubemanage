package service

import (
	"context"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MonitorService struct{}

func NewMonitorService() *MonitorService {
	return &MonitorService{}
}

func (m *MonitorService) GetClusterImageList(in *dto.ImageListInput) (*dto.ImageListOut, error) {
	//获取clientSet
	var total int
	k8sDB := &dao.K8SDB{Name: in.ClusterName}
	k8s, err := k8sDB.Find(k8sDB)
	if err != nil {
		return nil, err
	}
	client, err := kube.GetClient(k8s.Config)
	if err != nil {
		return nil, err
	}
	var out []dto.ImageListItem
	//查询所有Namespace
	NamespaceList, err := client.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, namespace := range NamespaceList.Items {
		// 查询所有deployment 镜像
		podlist, err := kube.K8s.clientSet.CoreV1().Pods(namespace.Name).List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, pod := range podlist.Items {
			for _, container := range pod.Spec.Containers {
				outItem := dto.ImageListItem{
					ClusterName: k8s.Name,
					NameSpace:   namespace.Name,
					AppName:     pod.Name,
					Image:       container.Image,
				}
				out = append(out, outItem)
				total++
			}
		}
	}
	return &dto.ImageListOut{
		Total: total,
		List:  out,
	}, nil
}
