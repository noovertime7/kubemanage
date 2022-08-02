package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type PodListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"required" comment:"过滤名"`
	NameSpace  string `json:"namespace" form:"namespace" validate:"required" comment:"命名空间"`
	Limit      int    `json:"limit" form:"limit" validate:"required" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"required" comment:"页码"`
}

func (params *PodListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type PodNameNsInput struct {
	PodName   string `json:"pod_name" form:"pod_name" comment:"POD名称" validate:"required"`
	NameSpace string `json:"name_space" form:"namespace" comment:"命名空间" validate:"required"`
}

type PodUpdateInput struct {
	PodNameNsInput
	Content string `json:"content" form:"content" comment:"内容" validate:"required"`
}

type PodGetLogInput struct {
	PodNameNsInput
	ContainerName string `json:"container_name" form:"container_name" comment:"容器名称" validate:"required"`
}

func (params *PodNameNsInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
