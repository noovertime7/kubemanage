package dao

import (
	"github.com/noovertime7/kubemanage/dao/authority"
	"github.com/noovertime7/kubemanage/dao/menu"
	"github.com/noovertime7/kubemanage/dao/user"
	"github.com/noovertime7/kubemanage/dao/workflow"
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
	WorkFlow() workflow.WorkFlowInterface
	User() user.User
	Authority() authority.Authority
	AuthorityMenu() authority.AuthorityMenu
	BaseMenu() menu.BaseMenu
}

func NewShareDaoFactory(db *gorm.DB) ShareDaoFactory {
	return &shareDaoFactory{db: db}
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (s *shareDaoFactory) WorkFlow() workflow.WorkFlowInterface {
	return workflow.NewWorkFlow(s.db)
}

func (s *shareDaoFactory) User() user.User {
	return user.NewUser(s.db)
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
