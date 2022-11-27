package v1

import (
	"context"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

type MenuGetter interface {
	Menu() MenuService
}

type MenuService interface {
	GetMenu(ctx context.Context, authorityId uint) ([]model.SysMenu, error)
	AddBaseMenu(ctx context.Context, in *dto.AddSysMenusInput) error
	AddMenuAuthority(ctx context.Context, menus []model.SysBaseMenu, authorityId uint) error
}

type menuService struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func NewMenuService(app *KubeManage) *menuService {
	return &menuService{app: app, factory: app.Factory}
}

func (m *menuService) GetMenu(ctx context.Context, authorityId uint) ([]model.SysMenu, error) {
	menuTree, err := m.getMenuTree(ctx, authorityId)
	if err != nil {
		return nil, err
	}
	//parent_id = 0 ,代表所有跟路由
	menus := menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = m.getChildrenList(&menus[i], menuTree)
	}
	return menus, nil
}

func (m *menuService) getMenuTree(ctx context.Context, authorityId uint) (map[string][]model.SysMenu, error) {
	var allMenus []model.SysMenu
	treeMap := make(map[string][]model.SysMenu)
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
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, nil
}

func (m *menuService) getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) error {
	// treeMap中包含所有路由
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		if err := m.getChildrenList(&menu.Children[i], treeMap); err != nil {
			return err
		}
	}
	return nil
}

// AddBaseMenu 添加基础路由
func (m *menuService) AddBaseMenu(ctx context.Context, in *dto.AddSysMenusInput) error {
	menuInfo := &model.SysBaseMenu{
		ParentId: in.ParentId,
		Name:     in.Name,
		Path:     in.Path,
		Hidden:   in.Hidden,
		Sort:     in.Sort,
		Meta:     in.Meta,
	}
	menu, err := m.factory.BaseMenu().Find(ctx, menuInfo)
	if !errors.Is(err, gorm.ErrRecordNotFound) && menu.ID != 0 {
		return errors.New("存在重复名称菜单，请修改菜单名称")
	}
	return m.factory.BaseMenu().Save(ctx, menuInfo)
}

// AddMenuAuthority 为角色增加menu树
func (m *menuService) AddMenuAuthority(ctx context.Context, menus []model.SysBaseMenu, authorityId uint) error {
	auth := &model.SysAuthority{AuthorityId: authorityId, SysBaseMenus: menus}
	return CoreV1.Authority().SetMenuAuthority(ctx, auth)
}
