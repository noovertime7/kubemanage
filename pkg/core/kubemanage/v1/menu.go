package v1

import (
	"context"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"strconv"
)

type MenuGetter interface {
	Menu() MenuService
}

type MenuService interface {
	GetMenuTree(ctx context.Context, authorityId uint) ([]model.SysMenu, error)
}

type menuService struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func NewMenuService(app *KubeManage) *menuService {
	return &menuService{app: app, factory: app.Factory}
}

func (m *menuService) GetMenuTree(ctx context.Context, authorityId uint) ([]model.SysMenu, error) {
	var allMenus []model.SysMenu
	SysAuthorityMenu := &model.SysAuthorityMenu{AuthorityId: strconv.Itoa(int(authorityId))}
	authorityMenus, err := m.factory.AuthorityMenu().FindList(ctx, SysAuthorityMenu)
	if err != nil {
		return nil, err
	}
	var MenuIds []string
	for i := range authorityMenus {
		MenuIds = append(MenuIds, authorityMenus[i].MenuId)
	}
	baseMenus, err := m.factory.BaseMenu().FindIn(ctx, MenuIds)
	if err != nil {
		return nil, err
	}
	for i := range baseMenus {
		allMenus = append(allMenus, model.SysMenu{
			SysBaseMenu: *baseMenus[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(baseMenus[i].ID),
		})
	}
	return allMenus, nil
}
