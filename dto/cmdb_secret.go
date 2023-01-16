package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
)

type CMDBSecretCreateInput struct {
	Name         string `json:"name" comment:"名称" validate:"required"`
	Protocol     uint   `json:"protocol" comment:"协议" validate:"required"`
	SecretType   uint   `json:"secretType" comment:"类型" validate:"required"`
	HostUserName string `json:"hostUserName" comment:"用户名" validate:"required"`
	HostPassword string `json:"hostPassword" comment:"密码"`
	Content      string `json:"content" comment:"备注"`
	PrivateKey   string `json:"privateKey" comment:"秘钥" `
}

func (params *CMDBSecretCreateInput) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}

type CMDBSecretUpdateInput struct {
	InstanceID   string `json:"instanceID" validate:"required"`
	Name         string `json:"name" comment:"名称" validate:"required"`
	Protocol     uint   `json:"protocol" comment:"协议" validate:"required"`
	SecretType   uint   `json:"secretType" comment:"类型" validate:"required"`
	HostUserName string `json:"hostUserName" comment:"用户名" validate:"required"`
	HostPassword string `json:"hostPassword" comment:"密码"`
	Content      string `json:"content" comment:"备注"`
	PrivateKey   string `json:"privateKey" comment:"秘钥" `
}

func (params *CMDBSecretUpdateInput) BindingValidParams(c *gin.Context) error {
	return pkg.DefaultGetValidParams(c, params)
}

type PageCMDBSecretOut struct {
	Total    int64              `json:"total"`
	List     []model.CMDBSecret `json:"list"`
	Page     int                `json:"page" form:"page"`         // 页码
	PageSize int                `json:"pageSize" form:"pageSize"` // 每页大小
}

type PageListCMDBSecretInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageListCMDBSecretInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageListCMDBSecretInput) GetPage() int {
	return p.Page
}

func (p *PageListCMDBSecretInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageListCMDBSecretInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageListCMDBSecretInput) Do(tx *gorm.DB) {
	tx.Where("name like ? or content like ?", "%"+p.Keyword+"%", "%"+p.Keyword+"%")
}
