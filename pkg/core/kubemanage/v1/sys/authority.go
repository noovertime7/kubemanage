package sys

import (
	"context"
	"github.com/noovertime7/kubemanage/dto"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
)

type AuthorityGetter interface {
	Authority() Authority
}

type Authority interface {
	SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error
	GetAuthorityList(ctx context.Context, pageInfo dto.PageInfo) (*dto.AuthorityList, error)
}

type authority struct {
	factory dao.ShareDaoFactory
}

func NewAuthority(factory dao.ShareDaoFactory) *authority {
	return &authority{factory: factory}
}

func (a *authority) SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error {
	return a.factory.Authority().SetMenuAuthority(ctx, auth)
}

func (a *authority) GetAuthorityList(ctx context.Context, pageInfo dto.PageInfo) (*dto.AuthorityList, error) {
	list, total, err := a.factory.Authority().PageList(ctx, pageInfo)
	if err != nil {
		return nil, err
	}
	return &dto.AuthorityList{
		PageInfo:          pageInfo,
		Total:             total,
		AuthorityListItem: list,
	}, nil
}
