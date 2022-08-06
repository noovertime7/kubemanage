package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type WorkFlowListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"" comment:"过滤名"`
	Limit      int    `json:"limit" form:"limit" validate:"" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"" comment:"页码"`
}

type WorkFlowCreateInput struct {
	Name          string                 `json:"name"`
	NameSpace     string                 `json:"namespace"`
	Replicas      int32                  `json:"replicas"`
	Deployment    string                 `json:"deployment"`
	Image         string                 `json:"image"`
	Label         map[string]string      `json:"label"`
	Cpu           string                 `json:"cpu"`
	Memory        string                 `json:"memory"`
	ContainerPort int32                  `json:"container_port"`
	HealthPath    string                 `json:"health_path"`
	HealthCheck   bool                   `json:"healthCheck"`
	Type          string                 `json:"type"`
	Port          int32                  `json:"port"`
	NodePort      int32                  `json:"node_port"`
	Hosts         map[string][]*HttpPath `json:"hosts"`
}

type WorkFlowIDInput struct {
	ID int `json:"id"`
}

func (params *WorkFlowCreateInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *WorkFlowListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *WorkFlowIDInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
