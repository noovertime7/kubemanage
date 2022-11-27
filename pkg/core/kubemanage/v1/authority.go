package v1

import (
	"context"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
)

type AuthorityGetter interface {
	Authority() Authority
}

type Authority interface {
	SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error
}

type authority struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func NewAuthority(app *KubeManage) *authority {
	return &authority{app: app, factory: app.Factory}
}

func (a *authority) SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error {
	return a.factory.Authority().SetMenuAuthority(ctx, auth)
}
