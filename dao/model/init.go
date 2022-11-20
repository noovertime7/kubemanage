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
	WorkFlowOrder
)

// SysUserEntities 用户初始化数据
var (
	SysUserEntities = []*SysUser{
		{
			UUID:        uuid.NewV4(),
			UserName:    "admin",
			Password:    "29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389",
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
			Password:    "29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389",
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
			Password:    "29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389",
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
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 9, Meta: Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 7, Meta: Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: Meta{Title: "excel导入导出", Icon: "takeaway-box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: Meta{Title: "断点续传", Icon: "upload-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: Meta{Title: "客户列表（资源示例）", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: Meta{Title: "系统工具", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: Meta{Title: "系统配置", Icon: "operation"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: Meta{Title: "字典管理", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: Meta{Title: "字典详情-${id}", Icon: "list", ActiveName: "dictionary"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: Meta{Title: "操作历史", Icon: "pie-chart"}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: Meta{Title: "断点续传（插件版）", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: Meta{Title: "官方网站", Icon: "home-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: true, ParentId: "14", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoPkg", Name: "autoPkg", Component: "view/systemTools/autoPkg/autoPkg.vue", Sort: 0, Meta: Meta{Title: "自动化package", Icon: "folder"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 6, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "https://plugin.gin-vue-admin.com/", Name: "https://plugin.gin-vue-admin.com/", Component: "https://plugin.gin-vue-admin.com/", Sort: 0, Meta: Meta{Title: "插件市场", Icon: "shop"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "installPlugin", Name: "installPlugin", Component: "view/systemTools/installPlugin/index.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "autoPlug", Name: "autoPlug", Component: "view/systemTools/autoPlug/autoPlug.vue", Sort: 2, Meta: Meta{Title: "插件模板", Icon: "folder"}},
		{MenuLevel: 0, Hidden: false, ParentId: "27", Path: "plugin-email", Name: "plugin-email", Component: "plugin/email/view/index.vue", Sort: 3, Meta: Meta{Title: "邮件插件", Icon: "message"}},
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
	{Ptype: "p", V0: pkg.AdminDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
	{Ptype: "p", V0: pkg.AdminDefaultAuthStr, V1: "/api/user/loginout ", V2: "GET"},
	{Ptype: "p", V0: pkg.AdminDefaultAuthStr, V1: "/api/menu/get_menus", V2: "GET"},

	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/user/loginout ", V2: "GET"},
	{Ptype: "p", V0: pkg.UserDefaultAuthStr, V1: "/api/menu/get_menus", V2: "GET"},

	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/user/login", V2: "POST"},
	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/user/loginout ", V2: "GET"},
	{Ptype: "p", V0: pkg.UserSubDefaultAuthStr, V1: "/api/menu/get_menus", V2: "GET"},
}
