package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/service"
	"github.com/wonderivan/logger"
)

var StatefulSet statefulSet

type statefulSet struct{}

func StatefulSetRegister(router *gin.RouterGroup) {
	router.DELETE("/del", StatefulSet.DeleteStatefulSet)
	router.PUT("/update", StatefulSet.UpdateStatefulSet)
	router.GET("/list", StatefulSet.GetStatefulSetList)
	router.GET("/detail", StatefulSet.GetStatefulSetDetail)
}

func (s *statefulSet) DeleteStatefulSet(ctx *gin.Context) {
	params := &dto.StatefulSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.StatefulSet.DeleteStatefulSet(params.Name, params.NameSpace); err != nil {
		logger.Error("删除StatefulSet失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (s *statefulSet) UpdateStatefulSet(ctx *gin.Context) {
	params := &dto.StatefulSetUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := service.StatefulSet.UpdateStatefulSet(params.Content, params.NameSpace); err != nil {
		logger.Error("更新StatefulSet失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

func (s *statefulSet) GetStatefulSetList(ctx *gin.Context) {
	params := &dto.StatefulSetListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.StatefulSet.GetStatefulSets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取StatefulSet列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *statefulSet) GetStatefulSetDetail(ctx *gin.Context) {
	params := &dto.StatefulSetNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := service.StatefulSet.GetStatefulSetDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取StatefulSet详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
