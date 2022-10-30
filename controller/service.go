package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/wonderivan/logger"
)

var ServiceController serviceController

type serviceController struct{}

func ServiceRegister(router *gin.RouterGroup) {
	router.POST("/create", ServiceController.CreateService)
	router.DELETE("/del", ServiceController.DeleteService)
	router.PUT("/update", ServiceController.UpdateService)
	router.GET("/list", ServiceController.GetServiceList)
	router.GET("/detail", ServiceController.GetServiceDetail)
	router.GET("/numnp", ServiceController.GetServicePerNS)
}

// CreateService 创建service
// ListPage godoc
// @Summary      创建service
// @Description  创建service
// @Tags         service
// @ID           /api/k8s/service/create
// @Accept       json
// @Produce      json
// @Param        body  body  dto.ServiceCreateInput  true  "body"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/service/create [post]
func (s *serviceController) CreateService(ctx *gin.Context) {
	params := &dto.ServiceCreateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 40001, err)
		return
	}
	if err := kube.Service.CreateService(params); err != nil {
		logger.Error("创建ervice失败", err)
		middleware.ResponseError(ctx, 40002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// DeleteService 删除service
// ListPage godoc
// @Summary      删除service
// @Description  删除service
// @Tags         service
// @ID           /api/k8s/service/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "service名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/service/del [delete]
func (s *serviceController) DeleteService(ctx *gin.Context) {
	params := &dto.ServiceNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 40001, err)
		return
	}
	if err := kube.Service.DeleteService(params.Name, params.NameSpace); err != nil {
		logger.Error("删除service失败", err)
		middleware.ResponseError(ctx, 40002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateService 更新service
// ListPage godoc
// @Summary      更新service
// @Description  更新service
// @Tags         service
// @ID           /api/k8s/service/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "service名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/service/update [put]
func (s *serviceController) UpdateService(ctx *gin.Context) {
	params := &dto.ServiceUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 40001, err)
		return
	}
	if err := kube.Service.UpdateService(params.NameSpace, params.Content); err != nil {
		logger.Error("修改service失败", err)
		middleware.ResponseError(ctx, 40002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "修改成功")
}

// GetServiceList 查看service列表
// ListPage godoc
// @Summary      查看service列表
// @Description  查看service列表
// @Tags         service
// @ID           /api/k8s/service/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/service/list [get]
func (s *serviceController) GetServiceList(ctx *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 40001, err)
		return
	}
	data, err := kube.Service.GetServiceList(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取service列表失败", err)
		middleware.ResponseError(ctx, 40002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetServiceDetail 获取service详情
// ListPage godoc
// @Summary      获取service详情
// @Description  获取service详情
// @Tags         service
// @ID           /api/k8s/service/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "service名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/service/detail [get]
func (s *serviceController) GetServiceDetail(ctx *gin.Context) {
	params := &dto.ServiceNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err)
		middleware.ResponseError(ctx, 40001, err)
		return
	}
	data, err := kube.Service.GetServiceDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取service详情失败", err)
		middleware.ResponseError(ctx, 40002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetServicePerNS 根据命名空间获取service数量
// ListPage godoc
// @Summary      根据命名空间获取service数量
// @Description  根据命名空间获取service数量
// @Tags         service
// @ID           /api/k8s/service/numnp
// @Accept       json
// @Produce      json
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data":service.serviceNp }"
// @Router       /api/k8s/service/numnp [get]
func (s *serviceController) GetServicePerNS(ctx *gin.Context) {
	data, err := kube.Service.GetServiceNp()
	if err != nil {
		logger.Error("获取service失败", err)
		middleware.ResponseError(ctx, 40002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
