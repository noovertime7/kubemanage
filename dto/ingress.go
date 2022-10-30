package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg"
	nwV1 "k8s.io/api/networking/v1"
)

type IngressCreteInput struct {
	Name      string                 `json:"name"`
	NameSpace string                 `json:"namespace"`
	Label     map[string]string      `json:"label"`
	Hosts     map[string][]*HttpPath `json:"hosts"`
}

type HttpPath struct {
	Path        string        `json:"path"`
	PathType    nwV1.PathType `json:"path_type"`
	ServiceName string        `json:"service_name"`
	ServicePort int32         `json:"service_port"`
}

type IngressNameNS struct {
	Name      string `json:"name" form:"name" comment:"服务名称" validate:"required"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
}

type IngressUpdateInput struct {
	Content   string `json:"content" form:"content" validate:"required" comment:"更新内容"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
}

type IngressListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
	NameSpace  string `json:"namespace" form:"namespace" validate:"" comment:"命名空间"`
	Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
}

func (params *IngressCreteInput) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}

func (params *IngressNameNS) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}

func (params *IngressUpdateInput) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}

func (params *IngressListInput) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}
