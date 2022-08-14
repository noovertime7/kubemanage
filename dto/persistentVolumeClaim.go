package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type PersistentVolumeClaimNameNS struct {
	Name      string `json:"name" form:"name" comment:"配置卷名称" validate:"required"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
}

type PersistentVolumeClaimUpdateInput struct {
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	Content   string `json:"content" form:"content"  validate:"required" comment:"更新内容"`
}

type PersistentVolumeClaimListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
	NameSpace  string `json:"namespace" form:"namespace" validate:"" comment:"命名空间"`
	Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
}

func (params *PersistentVolumeClaimNameNS) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *PersistentVolumeClaimUpdateInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *PersistentVolumeClaimListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
