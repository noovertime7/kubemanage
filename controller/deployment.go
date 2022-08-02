package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var Deployment deployment

type deployment struct{}

func DeploymentRegister(router *gin.RouterGroup) {
	router.POST("/create", Deployment.CreateDeployment)
	router.DELETE("/del", Deployment.DeleteDeployment)
	router.PUT("/update", Deployment.UpdateDeployment)
	router.GET("/list", Deployment.GetDeploymentList)
	router.GET("/detail", Deployment.GetDeploymentDetail)
	router.PUT("/restart", Deployment.RestartDeployment)
	router.GET("/scale", Deployment.ScaleDeployment)
	router.GET("/numnp", Deployment.GetDeploymentNumPreNS)
}

func (d *deployment) CreateDeployment(ctx *gin.Context) {
	params := &dto.DeployCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.Deployment.CreateDeployment(params); err != nil {
		logger.Error("新增deploy失败:", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "新增成功")
}

func (d *deployment) DeleteDeployment(ctx *gin.Context) {
	params := &dto.DeploymentNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.Deployment.DeleteDeployment(params.Name, params.NameSpace); err != nil {
		logger.Error("删除deploy失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (d *deployment) UpdateDeployment(ctx *gin.Context) {
	params := &dto.UpdateDeployInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.Deployment.UpdateDeployment(params.NameSpace, params.Content); err != nil {
		logger.Error("更新deploy失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

func (d *deployment) GetDeploymentList(ctx *gin.Context) {
	params := &dto.DeployListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.Deployment.GetDeployments(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取deploy列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (d *deployment) GetDeploymentDetail(ctx *gin.Context) {
	params := &dto.DeploymentNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.Deployment.GetDeploymentDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取deploy详细信息失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (d *deployment) GetDeploymentNumPreNS(ctx *gin.Context) {
	data, err := service.Deployment.GetDeployNumPerNS()
	if err != nil {
		logger.Error("获取deploy数量失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (d *deployment) RestartDeployment(ctx *gin.Context) {
	params := &dto.DeploymentNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.Deployment.RestartDeployment(params.Name, params.NameSpace); err != nil {
		logger.Error("获取deploy详细信息失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "重启Deployment成功")
}

func (d *deployment) ScaleDeployment(ctx *gin.Context) {
	params := &dto.DeployScaleInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	num, err := service.Deployment.ScaleDeployment(params.Name, params.NameSpace, params.ScaleNum)
	if err != nil {
		logger.Error("获取deploy详细信息失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, num)
}
