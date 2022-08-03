package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type NameSpaceNameInput struct {
	Name string `json:"name" form:"name" comment:"命名空间名称" validate:"required"`
}

type NameSpaceListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"required" comment:"过滤名"`
	Limit      int    `json:"limit" form:"limit" validate:"required" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"required" comment:"页码"`
}

func (params *NameSpaceListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *NameSpaceNameInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
