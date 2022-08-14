package service

import (
	"fmt"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dto"
)

var WorkFlow workflow

type workflow struct {
}

func (w *workflow) GetWorkFlowList(name string, page, limit int) (*dao.WorkflowResp, error) {
	var workflow *dao.Workflow
	workflows, total, err := workflow.PageList(&dto.WorkFlowListInput{
		FilterName: name,
		Page:       page,
		Limit:      limit,
	})
	if err != nil {
		return nil, err
	}
	return &dao.WorkflowResp{
		Items: workflows,
		Total: total,
	}, nil
}

func (w *workflow) GetWorkFlowByID(id int) (data *dao.Workflow, err error) {
	in := &dao.Workflow{ID: uint(id)}
	return in.Find(in)
}

func (w *workflow) CreateWorkFlow(params *dto.WorkFlowCreateInput) error {
	//若workflow不是ingress类型，传入空字符串即可
	var ingressName string
	if params.Type == "Ingress" {
		ingressName = getIngressName(params.Name)
	} else {
		ingressName = ""
	}
	dataWorkFlow := &dao.Workflow{
		Name:        params.Name,
		NameSpace:   params.NameSpace,
		Replicas:    params.Replicas,
		Deployment:  params.Deployment,
		Service:     getServiceName(params.Name),
		Ingress:     ingressName,
		ServiceType: params.Type,
		IsDeleted:   0,
	}
	//创建k8s资源
	if err := createWorkflowRes(params); err != nil {
		return err
	}
	return dataWorkFlow.Save()
}

// DelById 删除workflow
func (w *workflow) DelById(id int) (err error) {
	//获取workflow数据
	in := &dao.Workflow{ID: uint(id)}
	fmt.Println("in = ", in)
	workflow, err := in.Find(in)
	fmt.Println(workflow, in)
	if err != nil {
		return err
	}
	//删除k8s资源
	if err := delWorkflowRes(workflow); err != nil {
		return err
	}
	//删除数据库数据
	if err := in.DeleteById(); err != nil {
		return err
	}
	return nil
}

func delWorkflowRes(w *dao.Workflow) error {
	//删除deployment
	if err := Deployment.DeleteDeployment(w.Name, w.NameSpace); err != nil {
		return err
	}
	//删除service
	if err := Service.DeleteService(getServiceName(w.Name), w.NameSpace); err != nil {
		return err
	}
	//删除ingress，这里多了一层判断，因为只有type为ingress的workflow才有ingress资源
	if w.ServiceType == "Ingress" {
		if err := Ingress.DeleteIngress(getIngressName(w.Name), w.NameSpace); err != nil {
			return err
		}
	}
	return nil
}

func createWorkflowRes(params *dto.WorkFlowCreateInput) error {
	//声明service类型
	var serviceType string
	//组装DeployCreate类型的数据
	dc := &dto.DeployCreateInput{
		Name:          params.Name,
		NameSpace:     params.NameSpace,
		Replicas:      params.Replicas,
		Image:         params.Image,
		Labels:        params.Label,
		Cpu:           params.Cpu,
		Memory:        params.Memory,
		ContainerPort: params.ContainerPort,
		HealthCheck:   params.HealthCheck,
		HealthPath:    params.HealthPath,
	}
	//创建deployment
	if err := Deployment.CreateDeployment(dc); err != nil {
		return err
	}
	//判断service类型
	if params.Type != "Ingress" {
		serviceType = params.Type
	} else {
		serviceType = "ClusterIP"
	}

	//组装ServiceCreate类型的数据
	sc := &dto.ServiceCreateInput{
		Name:          getServiceName(params.Name),
		NameSpace:     params.NameSpace,
		Type:          serviceType,
		ContainerPort: params.ContainerPort,
		Port:          params.Port,
		NodePort:      params.NodePort,
		Label:         params.Label,
	}
	if err := Service.CreateService(sc); err != nil {
		return err
	}
	//组装IngressCreate类型的数据，创建ingress，只有ingress类型的workflow才有ingress资源，所以这里做了一层判断
	if params.Type == "Ingress" {
		ic := &dto.IngressCreteInput{
			Name:      getIngressName(params.Name),
			NameSpace: params.NameSpace,
			Label:     params.Label,
			Hosts:     params.Hosts,
		}
		if err := Ingress.CreateIngress(ic); err != nil {
			return err
		}
	}
	return nil
}

//workflow名字转换成service名字，添加-svc后缀
func getServiceName(workflowName string) (serviceName string) {
	return workflowName + "-svc"
}

//workflow名字转换成ingress名字，添加-ing后缀
func getIngressName(workflowName string) (ingressName string) {
	return workflowName + "-ing"
}
