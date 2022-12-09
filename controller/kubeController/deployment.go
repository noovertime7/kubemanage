package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubeDto"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	_ "k8s.io/api/apps/v1"

	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

var Deployment deployment

type deployment struct{}

// CreateDeployment 创建deployment
// ListPage godoc
// @Summary      创建deployment
// @Description  创建deployment
// @Tags         deployment
// @ID           /api/k8s/deployment/create
// @Accept       json
// @Produce      json
// @Param        body  body  kubernetes.DeployCreateInput  true  "body"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "新增成功}"
// @Router       /api/k8s/deployment/create [post]
func (d *deployment) CreateDeployment(ctx *gin.Context) {
	params := &kubeDto.DeployCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Deployment.CreateDeployment(params); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "新增成功")
}

// DeleteDeployment 删除deployment
// ListPage godoc
// @Summary      删除deployment
// @Description  删除deployment
// @Tags         deployment
// @ID           /api/k8s/deployment/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Deployment名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/deployment/del [delete]
func (d *deployment) DeleteDeployment(ctx *gin.Context) {
	params := &kubeDto.DeploymentNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Deployment.DeleteDeployment(params.Name, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateDeployment 更新deployment
// ListPage godoc
// @Summary      更新deployment
// @Description  更新deployment
// @Tags         deployment
// @ID           /api/k8s/deployment/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/deployment/update [put]
func (d *deployment) UpdateDeployment(ctx *gin.Context) {
	params := &kubeDto.UpdateDeployInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Deployment.UpdateDeployment(params.NameSpace, params.Content); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetDeploymentList 查看deployment列表
// ListPage godoc
// @Summary      查看deployment列表
// @Description  查看deployment列表
// @Tags         deployment
// @ID           /api/k8s/deployment/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/deployment/list [get]
func (d *deployment) GetDeploymentList(ctx *gin.Context) {
	params := &kubeDto.DeployListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Deployment.GetDeployments(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetDeploymentDetail 获取deployment详情
// ListPage godoc
// @Summary      获取deployment详情
// @Description  获取deployment详情
// @Tags         pod
// @ID           /api/k8s/deployment/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Deployment名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/deployment/detail [get]
func (d *deployment) GetDeploymentDetail(ctx *gin.Context) {
	params := &kubeDto.DeploymentNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Deployment.GetDeploymentDetail(params.Name, params.NameSpace)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetDeploymentNumPreNS 根据命名空间获取无状态控制器数量
// ListPage godoc
// @Summary      根据命名空间获取无状态控制器数量
// @Description  根据命名空间获取无状态控制器数量
// @Tags         deployment
// @ID           /api/k8s/deployment/numnp
// @Accept       json
// @Produce      json
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data":service.DeployNp }"
// @Router       /api/k8s/deployment/numnp [get]
func (d *deployment) GetDeploymentNumPreNS(ctx *gin.Context) {
	data, err := kube.Deployment.GetDeployNumPerNS()
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// RestartDeployment 重启deployment
// ListPage godoc
// @Summary      重启deployment
// @Description  重启deployment
// @Tags         deployment
// @ID           /api/k8s/deployment/restart
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": 重启Deployment成功}"
// @Router       /api/k8s/deployment/restart [put]
func (d *deployment) RestartDeployment(ctx *gin.Context) {
	params := &kubeDto.DeploymentNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Deployment.RestartDeployment(params.Name, params.NameSpace); err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "重启Deployment成功")
}

// ScaleDeployment 扩容deployment
// ListPage godoc
// @Summary      扩容deployment
// @Description  扩容deployment
// @Tags         deployment
// @ID           /api/k8s/deployment/scale
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        scale_num  query  int     true  "期望副本数"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": num}"
// @Router       /api/k8s/deployment/scale [get]
func (d *deployment) ScaleDeployment(ctx *gin.Context) {
	params := &kubeDto.DeployScaleInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	num, err := kube.Deployment.ScaleDeployment(params.Name, params.NameSpace, params.ScaleNum)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.ServerError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ServerError, err))
		return
	}
	middleware.ResponseSuccess(ctx, num)
}
