package dao

import (
	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/api"
	"github.com/noovertime7/kubemanage/dao/authority"
	"github.com/noovertime7/kubemanage/dao/dept"
	"github.com/noovertime7/kubemanage/dao/menu"
	"github.com/noovertime7/kubemanage/dao/operation"
	"github.com/noovertime7/kubemanage/dao/user"
	"github.com/noovertime7/kubemanage/dao/workflow"
)

// ShareDaoFactory 数据库抽象工厂 包含所有数据操作接口
type ShareDaoFactory interface {
	GetDB() *gorm.DB
	WorkFlow() workflow.WorkFlowInterface
	Department() dept.Department
	User() user.User
	Api() api.APi
	Authority() authority.Authority
	AuthorityMenu() authority.AuthorityMenu
	BaseMenu() menu.BaseMenu
	Opera() operation.Operation
}

func NewShareDaoFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{db: db}
}

var _ ShareDaoFactory = &shareDaoFactory{}

type shareDaoFactory struct {
	db *gorm.DB
}

func (s *shareDaoFactory) GetDB() *gorm.DB {
	return s.db
}

func (s *shareDaoFactory) WorkFlow() workflow.WorkFlowInterface {
	return workflow.NewWorkFlow(s.db)
}

func (s *shareDaoFactory) Department() dept.Department {
	return dept.NewDepartment(s.db)
}

func (s *shareDaoFactory) User() user.User {
	return user.NewUser(s.db)
}

func (s *shareDaoFactory) Api() api.APi {
	return api.NewApi(s.db)
}

func (s *shareDaoFactory) Authority() authority.Authority {
	return authority.NewAuthority(s.db)
}

func (s *shareDaoFactory) AuthorityMenu() authority.AuthorityMenu {
	return authority.NewAuthorityMenu(s.db)
}

func (s *shareDaoFactory) BaseMenu() menu.BaseMenu {
	return menu.NewBaseMenu(s.db)
}

func (s *shareDaoFactory) Opera() operation.Operation {
	return operation.NewOperation(s.db)
}
