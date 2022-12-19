package sys

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
)

type AuthorityGetter interface {
	Authority() Authority
}

type Authority interface {
	CreateAuthority(ctx context.Context, aid uint, name string) error
	DeleteAuthority(ctx context.Context, aid uint) error
	UpdateAuthority(ctx context.Context, aid uint, name string) error
	SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error
	PageAuthority(ctx context.Context, pageInfo dto.PageInfo) (*dto.AuthorityList, error)
}

type authority struct {
	factory dao.ShareDaoFactory
	menu    MenuService
	casbin  CasbinService
}

func NewAuthority(factory dao.ShareDaoFactory) Authority {
	return &authority{factory: factory, menu: NewMenuService(factory), casbin: NewCasbinService(factory)}
}

func (a *authority) SetMenuAuthority(ctx context.Context, auth *model.SysAuthority) error {
	return a.factory.Authority().SetMenuAuthority(ctx, auth)
}

func (a *authority) CreateAuthority(ctx context.Context, aid uint, name string) error {
	authDB := &model.SysAuthority{AuthorityId: aid}
	auth, err := a.factory.Authority().Find(ctx, authDB)
	if err != nil {
		return err
	}
	if auth.AuthorityId != 0 {
		return errors.New(fmt.Sprintf("存在重复角色ID:%d", aid))
	}
	if err := a.factory.Authority().Save(ctx, &model.SysAuthority{AuthorityId: aid, AuthorityName: name}); err != nil {
		return err
	}
	// 添加基础权限 dashboard
	dashboardMenu := model.SysBaseMenuEntities[:1]
	if err := a.menu.AddMenuAuthority(ctx, dashboardMenu, aid); err != nil {
		return err
	}
	// 添加基础APi权限
	return a.casbin.AddCasbin(aid, ConvertTOCasbinRules(model.BasicApiRule))
}

func ConvertTOCasbinRules(in []model.BasicCasbinInfo) []CasbinRule {
	out := make([]CasbinRule, len(in))
	for i := range in {
		out[i] = in[i]
	}
	return out
}

func (a *authority) UpdateAuthority(ctx context.Context, aid uint, name string) error {
	authDB := &model.SysAuthority{AuthorityId: aid}
	auth, err := a.factory.Authority().Find(ctx, authDB)
	if err != nil {
		return err
	}
	auth.AuthorityName = name
	return a.factory.Authority().Updates(ctx, auth)
}

func (a *authority) DeleteAuthority(ctx context.Context, aid uint) error {
	//查询当前角色是否还有用户
	authDB := &model.SysAuthority{AuthorityId: aid}
	auth, err := a.factory.Authority().FindAllInfo(ctx, authDB)
	if err != nil {
		return err
	}
	if len(auth.Users) != 0 {
		return errors.New("当前角色还有未解除绑定的用户")
	}
	//删除接口关联
	a.casbin.RemoveCasbinByAuthority(aid)
	// 删除菜单关联
	if len(auth.SysBaseMenus) > 0 {
		if err := a.factory.Authority().DeleteAuthorityMenu(ctx, auth, auth.SysBaseMenus); err != nil {
			return err
		}
	}
	return a.factory.Authority().Delete(ctx, auth)
}

func (a *authority) PageAuthority(ctx context.Context, pageInfo dto.PageInfo) (*dto.AuthorityList, error) {
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
