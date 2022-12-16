package dto

import "github.com/noovertime7/kubemanage/dao/model"

type AuthorityList struct {
	PageInfo
	Total             int64                `json:"total"`
	AuthorityListItem []model.SysAuthority `json:"list"`
}
