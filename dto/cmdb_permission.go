package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
	"gorm.io/gorm"
)

type PageCMDBPermissionOut struct {
	Total    int64               `json:"total"`
	List     []*model.Permission `json:"list"`
	Page     int                 `json:"page" form:"page"`         // 页码
	PageSize int                 `json:"pageSize" form:"pageSize"` // 每页大小
}

type PageListCMDBPermissionInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageListCMDBPermissionInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageListCMDBPermissionInput) GetPage() int {
	return p.Page
}

func (p *PageListCMDBPermissionInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageListCMDBPermissionInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageListCMDBPermissionInput) Do(tx *gorm.DB) {
	tx.Where("name like ? or content like ?", "%"+p.Keyword+"%", "%"+p.Keyword+"%")
}
