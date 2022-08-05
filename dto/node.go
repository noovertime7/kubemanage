package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type NodeNameInput struct {
	Name string `json:"name" form:"name" comment:"Node名称" validate:"required"`
}

type NodeListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
	Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
}

func (params *NodeNameInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *NodeListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
