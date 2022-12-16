package model

import (
	"database/sql"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/noovertime7/kubemanage/pkg"
	uuid "github.com/satori/go.uuid"
)

// 初始化顺序
const (
	SysUserOrder = iota
	MenuAuthorityOrder
	SysBaseMenuOrder
	SysAuthorityOrder
	CasbinInitOrder
	OperatorationOrder
	WorkFlowOrder
)

// SysUserEntities 用户初始化数据
var (
	SysUserEntities = []*SysUser{
		{
			UUID:        uuid.NewV4(),
			UserName:    "admin",
			Password:    "$2a$14$Zfb6w0UDBFMN0.nJeVXCUO3zH/iWKGtbBYyIzDDRnC..EgTS0Et0S",
			NickName:    "admin",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: pkg.AdminDefaultAuth,
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
		{
			UUID:        uuid.NewV4(),
			UserName:    "chenteng",
			Password:    "$2a$14$yLCxKYP46M2NRnXujYe3mOfNe00GtBtjpaLM2eIzYCzYKQXqzsuka",
			NickName:    "chenteng",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: pkg.UserDefaultAuth,
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
		{
			UUID:        uuid.NewV4(),
			UserName:    "chentengsub",
			Password:    "$2a$14$MPINiht5QO2wlR3DynizXOtuqcNAOrNZdrSUKXrbjqcKbK.jcfyAW",
			NickName:    "chentengsub",
			SideMode:    "dark",
			Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
			BaseColor:   "#fff",
			ActiveColor: "#1890ff",
			AuthorityId: pkg.UserSubDefaultAuth,
			Phone:       "12345678901",
			Email:       "test@qq.com",
			Enable:      1,
			Status:      sql.NullInt64{Int64: 0, Valid: true},
		},
	}
)

// SysBaseMenuEntities 菜单初始化数据
var (
	SysBaseMenuEntities = []SysBaseMenu{
		// 根菜单
		{MenuLevel: 0, Hidden: false, Disabled: true, ParentId: "0", Path: "dashboard", Name: "仪表盘", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "cmdb", Name: "资产中心", Sort: 3, Meta: Meta{Title: "资产中心", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "kubernetes", Name: "容器管理", Sort: 4, Meta: Meta{Title: "容器管理", Icon: "menu"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "devops", Name: "应用发布", Sort: 5, Meta: Meta{Title: "应用发布", Icon: "compass"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "0", Path: "setting", Name: "系统设置", Sort: 6, Meta: Meta{Title: "系统设置", Icon: "setting"}},
		//子菜单 ParentId对应跟菜单顺序 且不需要icon
		// 资产中心子菜单
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "2", Path: "host", Name: "主机管理", Sort: 0, Meta: Meta{Title: "主机管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "2", Path: "secret", Name: "授权管理", Sort: 1, Meta: Meta{Title: "授权管理"}},
		// 容器管理子菜单
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "cluster", Name: "集群管理", Sort: 0, Meta: Meta{Title: "集群管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "deployment", Name: "工作负载", Sort: 1, Meta: Meta{Title: "工作负载"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "service", Name: "服务发现", Sort: 2, Meta: Meta{Title: "服务发现"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "node", Name: "节点管理", Sort: 3, Meta: Meta{Title: "节点管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "config", Name: "配置中心", Sort: 4, Meta: Meta{Title: "配置中心"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "3", Path: "events", Name: "事件中心", Sort: 5, Meta: Meta{Title: "事件中心"}},
		// 系统设置子菜单
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "5", Path: "authority", Name: "角色管理", Sort: 1, Meta: Meta{Title: "角色管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "5", Path: "user", Name: "用户管理", Sort: 2, Meta: Meta{Title: "用户管理"}},
		{MenuLevel: 0, Hidden: false, Disabled: false, ParentId: "5", Path: "operation", Name: "操作历史", Sort: 3, Meta: Meta{Title: "操作历史"}},
	}
)

// SysAuthorityEntities 角色初始化数据
var (
	SysAuthorityEntities = []SysAuthority{
		{
			AuthorityId:   pkg.AdminDefaultAuth,
			AuthorityName: "管理员",
			DefaultRouter: "dashboard",
			ParentId:      0,
		},
		{
			AuthorityId:   pkg.UserDefaultAuth,
			AuthorityName: "普通用户",
			DefaultRouter: "dashboard",
			ParentId:      0,
		},
		{
			AuthorityId:   pkg.UserSubDefaultAuth,
			AuthorityName: "普通用户子角色",
			DefaultRouter: "dashboard",
			ParentId:      222,
		},
	}
)

var CasbinApi = []adapter.CasbinRule{
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/user/loginout", V2: "GET"},
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/menu/get_menus", V2: "GET"},
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/user/getinfo", V2: "GET"},
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/user/:id/change_pwd", V2: "POST"},

	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/user/loginout", V2: "GET"},
	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/menu/get_menus", V2: "GET"},
	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/user/getinfo", V2: "GET"},
	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/user/:id/change_pwd", V2: "POST"},
}
