package kube

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/noovertime7/kubemanage/dto/kubeDto"
)

var Deployment deployment

type deployment struct{}

type DeploymentResp struct {
	Total int                 `json:"total"`
	Items []appsV1.Deployment `json:"items"`
}

type DeployNp struct {
	NameSpace string `json:"namespace"`
	DeployNum int    `json:"deployment_num"`
}

func (d *deployment) toCells(deployments []appsV1.Deployment) []DataCell {
	cells := make([]DataCell, len(deployments))
	for i := range deployments {
		cells[i] = deploymentCell(deployments[i])
	}
	return cells
}

func (d *deployment) FromCells(cells []DataCell) []appsV1.Deployment {
	deployments := make([]appsV1.Deployment, len(cells))
	for i := range cells {
		deployments[i] = appsV1.Deployment(cells[i].(deploymentCell))
	}
	return deployments
}

func (d *deployment) GetDeployments(filterName, namespace string, limit, page int) (*DeploymentResp, error) {
	deploymentList, err := K8s.ClientSet.AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	selectableData := &dataSelector{
		GenericDataList: d.toCells(deploymentList.Items),
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
	deployments := d.FromCells(data.GenericDataList)
	return &DeploymentResp{
		Total: total,
		Items: deployments,
	}, nil
}

// GetDeploymentDetail 获取deployment详情
func (d *deployment) GetDeploymentDetail(deployName, namespace string) (*appsV1.Deployment, error) {
	deploy, err := K8s.ClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deployName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

// ScaleDeployment 设置deployment副本数
func (d *deployment) ScaleDeployment(deployName, namespace string, scaleNum int) (int32, error) {
	scale, err := K8s.ClientSet.AppsV1().Deployments(namespace).GetScale(context.TODO(), deployName, metaV1.GetOptions{})
	if err != nil {
		return 0, err
	}
	//修改副本数
	scale.Spec.Replicas = int32(scaleNum)
	//更新副本数，传入scale对象
	newScale, err := K8s.ClientSet.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deployName, scale, metaV1.UpdateOptions{})
	if err != nil {
		return 0, err
	}
	return newScale.Spec.Replicas, nil
}

// CreateDeployment 新增deployment,接收deployCreate的对象
func (d *deployment) CreateDeployment(data *kubeDto.DeployCreateInput) error {
	//初始化appsV1.deployment类型的对象
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.NameSpace,
			Labels:    data.Labels,
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &data.Replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels:      data.Labels,
				MatchExpressions: nil,
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name:   data.Name,
					Labels: data.Labels,
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  data.Name,
							Image: data.Image,
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.ProtocolTCP,
									ContainerPort: data.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
		Status: appsV1.DeploymentStatus{},
	}
	//判断健康检查功能是否打开
	if data.HealthCheck {
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &coreV1.Probe{
			ProbeHandler: coreV1.ProbeHandler{
				HTTPGet: &coreV1.HTTPGetAction{
					Path: data.HealthPath,
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
					Host:        "",
					Scheme:      "",
					HTTPHeaders: nil,
				},
			},
			InitialDelaySeconds: 5,
			TimeoutSeconds:      5,
			PeriodSeconds:       5,
		}
		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = &coreV1.Probe{
			ProbeHandler: coreV1.ProbeHandler{
				HTTPGet: &coreV1.HTTPGetAction{
					Path: data.HealthPath,
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
					Host:        "",
					Scheme:      "",
					HTTPHeaders: nil,
				},
			},
			InitialDelaySeconds: 15,
			TimeoutSeconds:      5,
			PeriodSeconds:       5,
		}
	}
	//定义容器的limit与request资源
	deployment.Spec.Template.Spec.Containers[0].Resources.Limits = map[coreV1.ResourceName]resource.Quantity{
		coreV1.ResourceCPU:    resource.MustParse(data.Cpu),
		coreV1.ResourceMemory: resource.MustParse(data.Memory),
	}
	deployment.Spec.Template.Spec.Containers[0].Resources.Requests = map[coreV1.ResourceName]resource.Quantity{
		coreV1.ResourceCPU:    resource.MustParse(data.Cpu),
		coreV1.ResourceMemory: resource.MustParse(data.Memory),
	}
	//调用sdk去更新deployment
	if _, err := K8s.ClientSet.AppsV1().Deployments(data.NameSpace).Create(context.TODO(), deployment, metaV1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}

// DeleteDeployment 删除deployment
func (d *deployment) DeleteDeployment(deployName, namespace string) error {
	return K8s.ClientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), deployName, metaV1.DeleteOptions{})
}

// UpdateDeployment 更新deployment
func (d *deployment) UpdateDeployment(namespace, content string) error {
	var deploy = &appsV1.Deployment{}
	if err := json.Unmarshal([]byte(content), deploy); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (d *deployment) RestartDeployment(deployName, namespace string) error {
	//随便改一个无关的值,就会触发重启
	patchData := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{
							"name": deployName,
							"env": []map[string]string{{
								"name":  "RESTART_",
								"value": strconv.FormatInt(time.Now().Unix(), 10),
							}},
						},
					},
				},
			},
		},
	}
	//序列化为字节，因为patch方法只能接受字节类型的参数
	patchByte, err := json.Marshal(patchData)
	if err != nil {
		return err
	}
	//调用patch方法更新deployment
	if _, err := K8s.ClientSet.AppsV1().Deployments(namespace).Patch(context.TODO(), deployName, "application/strategic-merge-patch+json", patchByte, metaV1.PatchOptions{}); err != nil {
		return err
	}
	return nil
}

// GetDeployNumPerNS 获取每个namespace下的deploy数量
func (d *deployment) GetDeployNumPerNS() ([]*DeployNp, error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var deploys []*DeployNp
	for _, namespace := range namespaceList.Items {
		deployList, err := K8s.ClientSet.AppsV1().Deployments(namespace.Name).List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			return nil, err
		}
		//组装数据
		deployNp := &DeployNp{
			NameSpace: namespace.Name,
			DeployNum: len(deployList.Items),
		}
		deploys = append(deploys, deployNp)
	}
	return deploys, nil
}
