package kubeController

import (
	"github.com/gin-gonic/gin"
)

type kubeRouter struct{}

func NewKubeRouter(ginEngine *gin.RouterGroup) {
	k := kubeRouter{}
	k.initRoutes(ginEngine)
}

func (k *kubeRouter) initRoutes(ginEngine *gin.RouterGroup) {
	k8sRoute := ginEngine.Group("/k8s")
	{
		k8sRoute.POST("/deployment/create", Deployment.CreateDeployment)
		k8sRoute.DELETE("/deployment/del", Deployment.DeleteDeployment)
		k8sRoute.PUT("/deployment/update", Deployment.UpdateDeployment)
		k8sRoute.GET("/deployment/list", Deployment.GetDeploymentList)
		k8sRoute.GET("/deployment/detail", Deployment.GetDeploymentDetail)
		k8sRoute.PUT("/deployment/restart", Deployment.RestartDeployment)
		k8sRoute.GET("/deployment/scale", Deployment.ScaleDeployment)
		k8sRoute.GET("/deployment/numnp", Deployment.GetDeploymentNumPreNS)
	}
	{
		k8sRoute.GET("/pod/list", Pod.GetPods)
		k8sRoute.GET("/pod/detail", Pod.GetPodDetail)
		k8sRoute.DELETE("/pod/del", Pod.DeletePod)
		k8sRoute.PUT("/pod/update", Pod.UpdatePod)
		k8sRoute.GET("/pod/container", Pod.GetPodContainer)
		k8sRoute.GET("/pod/log", Pod.GetPodLog)
		k8sRoute.GET("/pod/numnp", Pod.GetPodNumPreNp)
		k8sRoute.GET("/pod/webshell", Pod.WebShell)
	}
	{
		k8sRoute.DELETE("/daemonset/del", DaemonSet.DeleteDaemonSet)
		k8sRoute.PUT("/daemonset/update", DaemonSet.UpdateDaemonSet)
		k8sRoute.GET("/daemonset/list", DaemonSet.GetDaemonSetList)
		k8sRoute.GET("/daemonset/detail", DaemonSet.GetDaemonSetDetail)
	}
	{
		k8sRoute.DELETE("/statefulset/del", StatefulSet.DeleteStatefulSet)
		k8sRoute.PUT("/statefulset/update", StatefulSet.UpdateStatefulSet)
		k8sRoute.GET("/statefulset/list", StatefulSet.GetStatefulSetList)
		k8sRoute.GET("/statefulset/detail", StatefulSet.GetStatefulSetDetail)
	}
	{
		k8sRoute.GET("/node/list", Node.GetNodeList)
		k8sRoute.GET("/node/detail", Node.GetNodeDetail)
	}

	{
		k8sRoute.PUT("/namespace/create", NameSpace.CreateNameSpace)
		k8sRoute.DELETE("/namespace/del", NameSpace.DeleteNameSpace)
		k8sRoute.GET("/namespace/list", NameSpace.GetNameSpaceList)
		k8sRoute.GET("/namespace/detail", NameSpace.GetNameSpaceDetail)
	}

	{
		k8sRoute.DELETE("/persistentvolume/del", PersistentVolume.DeletePersistentVolume)
		k8sRoute.GET("/persistentvolume/list", PersistentVolume.GetPersistentVolumeList)
		k8sRoute.GET("/persistentvolume/detail", PersistentVolume.GetPersistentVolumeDetail)
	}

	{
		k8sRoute.POST("/service/create", ServiceController.CreateService)
		k8sRoute.DELETE("/service/del", ServiceController.DeleteService)
		k8sRoute.PUT("/service/update", ServiceController.UpdateService)
		k8sRoute.GET("/service/list", ServiceController.GetServiceList)
		k8sRoute.GET("/service/detail", ServiceController.GetServiceDetail)
		k8sRoute.GET("/service/numnp", ServiceController.GetServicePerNS)
	}

	{
		k8sRoute.PUT("/ingress/create", IngressController.CreateIngress)
		k8sRoute.DELETE("/ingress/del", IngressController.DeleteIngress)
		k8sRoute.PUT("/ingress/update", IngressController.UpdateIngress)
		k8sRoute.GET("/ingress/list", IngressController.GetIngressList)
		k8sRoute.GET("/ingress/detail", IngressController.GetIngressDetail)
		k8sRoute.GET("/ingress/numnp", IngressController.GetIngressNumPreNp)
	}

	{
		k8sRoute.DELETE("/configmap/del", Configmap.DeleteConfigmap)
		k8sRoute.PUT("/configmap/update", Configmap.UpdateConfigmap)
		k8sRoute.GET("/configmap/list", Configmap.GetConfigmapList)
		k8sRoute.GET("/configmap/detail", Configmap.GetConfigmapDetail)
	}

	{
		k8sRoute.DELETE("/persistentvolumeclaim/del", PersistentVolumeClaim.DeletePersistentVolumeClaim)
		k8sRoute.PUT("/persistentvolumeclaim/update", PersistentVolumeClaim.UpdatePersistentVolumeClaim)
		k8sRoute.GET("/persistentvolumeclaim/list", PersistentVolumeClaim.GetPersistentVolumeClaimList)
		k8sRoute.GET("/persistentvolumeclaim/detail", PersistentVolumeClaim.GetPersistentVolumeClaimDetail)
	}

	{
		k8sRoute.DELETE("/secret/del", Secret.DeleteSecret)
		k8sRoute.PUT("/secret/update", Secret.UpdateSecret)
		k8sRoute.GET("/secret/list", Secret.GetSecretList)
		k8sRoute.GET("/secret/detail", Secret.GetSecretDetail)
	}

	{
		k8sRoute.POST("/workflow/create", WorkFlow.CreateWorkFlow)
		k8sRoute.DELETE("/workflow/del", WorkFlow.DeleteWorkflow)
		k8sRoute.GET("/workflow/list", WorkFlow.GetWorkflowList)
		k8sRoute.GET("/workflow/id", WorkFlow.GetWorkflowByID)
	}

}
