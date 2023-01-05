package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/pkg"
)

// Empty 用于swag生成不需要传参的api
type Empty struct{}

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (a PageInfo) GetPage() int {
	return a.Page
}

func (a PageInfo) GetPageSize() int {
	return a.PageSize
}

func (a PageInfo) IsFitter() bool {
	return a.Keyword != ""
}

func (a PageInfo) GetFitter() string {
	return a.Keyword
}

func (a *PageInfo) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// BindingValidParams 绑定并校验参数
func (a *IdsReq) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}

type InstancesReq struct {
	Ids []string `json:"ids" form:"ids"`
}

func (a *InstancesReq) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, a)
}
