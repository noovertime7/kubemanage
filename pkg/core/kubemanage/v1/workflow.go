package v1

import (
	"context"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
)

type WorkFlowServiceGetter interface {
	WorkFlow() WorkFlowService
}

type WorkFlowService interface {
	Save(context.Context, *dto.WorkFlowCreateInput) error
	Find(context.Context, *dto.WorkFlowIDInput) (*model.Workflow, error)
	FindList(context.Context, *dto.WorkFlowListInput) (*WorkflowResp, error)
	Delete(context.Context, int) error
}

type workflow struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

var _ WorkFlowService = &workflow{}

func NewWorkFlow(app *KubeManage) *workflow {
	return &workflow{
		app:     app,
		factory: app.Factory,
	}
}

type WorkflowResp struct {
	Items []*model.Workflow `json:"items"`
	Total int               `json:"total"`
}

func (w *workflow) Save(ctx context.Context, params *dto.WorkFlowCreateInput) error {
	//若workflow不是ingress类型，传入空字符串即可
	var ingressName string
	if params.Type == "Ingress" {
		ingressName = getIngressName(params.Name)
	} else {
		ingressName = ""
	}
	dataWorkFlow := &model.Workflow{
		Name:        params.Name,
		NameSpace:   params.NameSpace,
		Replicas:    params.Replicas,
		Deployment:  params.Deployment,
		Service:     getServiceName(params.Name),
		Ingress:     ingressName,
		ServiceType: params.Type,
	}
	//创建k8s资源
	if err := createWorkflowRes(params); err != nil {
		return err
	}
	return w.factory.WorkFlow().Save(ctx, dataWorkFlow)
}

// Delete 删除workflow
func (w *workflow) Delete(ctx context.Context, id int) (err error) {
	//删除k8s资源
	if err := w.delWorkflowRes(ctx, id); err != nil {
		return err
	}
	//删除数据库数据
	if err := w.factory.WorkFlow().Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (w *workflow) Find(ctx context.Context, params *dto.WorkFlowIDInput) (data *model.Workflow, err error) {
	return w.factory.WorkFlow().Find(ctx, params.ID)
}

func (w *workflow) FindList(ctx context.Context, params *dto.WorkFlowListInput) (*WorkflowResp, error) {
	workflows, total, err := w.factory.WorkFlow().PageList(ctx, &dto.WorkFlowListInput{
		FilterName: params.FilterName,
		Page:       params.Page,
		Limit:      params.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &WorkflowResp{
		Items: workflows,
		Total: total,
	}, nil
}

func (w *workflow) delWorkflowRes(ctx context.Context, id int) error {
	workFlowInfo, err := w.Find(ctx, &dto.WorkFlowIDInput{ID: id})
	if err != nil {
		return err
	}
	//删除deployment
	if err := kube.Deployment.DeleteDeployment(workFlowInfo.Name, workFlowInfo.NameSpace); err != nil {
		return err
	}
	//删除service
	if err := kube.Service.DeleteService(getServiceName(workFlowInfo.Name), workFlowInfo.NameSpace); err != nil {
		return err
	}
	//删除ingress，这里多了一层判断，因为只有type为ingress的workflow才有ingress资源
	if workFlowInfo.ServiceType == "Ingress" {
		if err := kube.Ingress.DeleteIngress(getIngressName(workFlowInfo.Name), workFlowInfo.NameSpace); err != nil {
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
	if err := kube.Deployment.CreateDeployment(dc); err != nil {
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
	if err := kube.Service.CreateService(sc); err != nil {
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
		if err := kube.Ingress.CreateIngress(ic); err != nil {
			return err
		}
	}
	return nil
}

// workflow名字转换成service名字，添加-svc后缀
func getServiceName(workflowName string) (serviceName string) {
	return workflowName + "-svc"
}

// workflow名字转换成ingress名字，添加-ing后缀
func getIngressName(workflowName string) (ingressName string) {
	return workflowName + "-ing"
}
