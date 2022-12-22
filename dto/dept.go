package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
	"gorm.io/gorm"
)

type PageListDeptInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageListDeptInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageListDeptInput) GetPage() int {
	return p.Page
}

func (p *PageListDeptInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageListDeptInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageListDeptInput) Do(tx *gorm.DB) {
	tx.Where("dept_name = ?", p.Keyword)
}

type PageListDeptOut struct {
	Total int64              `json:"total"`
	List  []model.Department `json:"list"`
}
