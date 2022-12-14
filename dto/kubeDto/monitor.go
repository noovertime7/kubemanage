package kubeDto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg"
)

type ImageListInput struct {
	ClusterName string `json:"cluster_name" form:"cluster_name" validate:"required"`
}

type ImageListOut struct {
	Total int             `json:"total"`
	List  []ImageListItem `json:"list"`
}

type ImageListItem struct {
	ClusterName string `json:"cluster_name"`
	NameSpace   string `json:"name_space"`
	AppName     string `json:"app_name"`
	Image       string `json:"image"`
}

func (params *ImageListInput) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}
