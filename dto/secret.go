package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type SecretNameNS struct {
	Name      string `json:"name" form:"name" comment:"有状态控制器名称" validate:"required"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
}

type SecretUpdateInput struct {
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	Content   string `json:"content" validate:"required" comment:"更新内容"`
}

type SecretListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
	NameSpace  string `json:"namespace" form:"namespace" validate:"" comment:"命名空间"`
	Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
}

func (params *SecretNameNS) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *SecretUpdateInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *SecretListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
