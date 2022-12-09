package kube

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Pod pod

type pod struct{}

type PodsResp struct {
	Total int          `json:"total"`
	Items []coreV1.Pod `json:"items"`
}

type PodsNp struct {
	Namespace string `json:"namespace"`
	PodNum    int    `json:"pod_num"`
}

// GetPods 获取pod列表支持、过滤、排序以及分页
func (p *pod) GetPods(filterName, namespace string, limit, page int) (podsResp *PodsResp, err error) {
	podlist, err := K8s.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
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

// GetPodDetail 获取Pod详情
func (p *pod) GetPodDetail(podName, namespace string) (pod *coreV1.Pod, err error) {
	podRes, err := K8s.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return podRes, nil
}

// DeletePod 删除Pod
func (p *pod) DeletePod(podName, namespace string) error {
	err := K8s.ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metaV1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

// UpdatePod 更新Pod
func (p *pod) UpdatePod(namespace, content string) error {
	var pod = &coreV1.Pod{}
	//将json反序列换为pod类型
	if err := json.Unmarshal([]byte(content), pod); err != nil {
		return err
	}
	_, err := K8s.ClientSet.CoreV1().Pods(namespace).Update(context.TODO(), pod, metaV1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

// GetPodContainer 获取Pod容器名
func (p *pod) GetPodContainer(podName, namespace string) (containers []string, err error) {
	pod, err := p.GetPodDetail(podName, namespace)
	if err != nil {
		return nil, err
	}
	//从pod中获取containers
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

// GetPodLog 获取容器日志
func (p *pod) GetPodLog(containerName, podName, namespace string) (log string, err error) {
	//设置日志的配置，容器名，获取的内容的配置
	lineLimit := int64(100)
	op := &coreV1.PodLogOptions{
		Container: containerName,
		TailLines: &lineLimit,
	}
	//获取request的实例
	req := K8s.ClientSet.CoreV1().Pods(namespace).GetLogs(podName, op)
	//发起stream连接，得到response.body
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer func(podLogs io.ReadCloser) {
		err := podLogs.Close()
		if err != nil {

		}
	}(podLogs)
	//将body写入缓冲区，转换为可读的string类型
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, podLogs); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// GetPodNumPerNp 获取namespace下的Pod数量
func (p *pod) GetPodNumPerNp() (podsNps []*PodsNp, err error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, namespace := range namespaceList.Items {
		podList, err := K8s.ClientSet.CoreV1().Pods(namespace.Name).List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			return nil, err
		}
		//组装数据
		podsNp := &PodsNp{
			Namespace: namespace.Name,
			PodNum:    len(podList.Items),
		}
		podsNps = append(podsNps, podsNp)
	}
	return podsNps, nil
}

// 类型转换的方法 coreV1.pod => DataCell,DataCell => coreV1.pod
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
