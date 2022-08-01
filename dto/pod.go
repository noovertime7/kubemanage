package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type PodListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"required" comment:"年龄"`
	NameSpace  string `json:"name_space" form:"namespace" validate:"required" comment:"年龄"`
	Limit      int    `json:"limit" form:"limit" validate:"required" comment:"年龄"`
	Page       int    `json:"page" form:"page" validate:"required" comment:"年龄"`
}

func (params *PodListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
