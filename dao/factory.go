package dao

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/noovertime7/kubemanage/dao/api"
	"github.com/noovertime7/kubemanage/dao/authority"
	"github.com/noovertime7/kubemanage/dao/cmdb"
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
	CMDB() cmdb.CMDBFactory
	Transactioner
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

func (s *shareDaoFactory) CMDB() cmdb.CMDBFactory {
	return cmdb.NewCMDBFactory(s.db)
}

type Transactioner interface {
	Begin(opts ...*sql.TxOptions)
	Commit()
	Rollback()
}

// 存放之前的db实例，一定要传递引用，否则开启事务后，oldTx会被修改掉
var oldTx gorm.DB

func (s *shareDaoFactory) Begin(opts ...*sql.TxOptions) {
	oldTx = *s.db
	tx := s.db.Begin(opts...)
	*s.db = *tx
}

func (s *shareDaoFactory) Commit() {
	s.db.Commit()
	*s.db = oldTx
}

func (s *shareDaoFactory) Rollback() {
	s.db.Rollback()
	*s.db = oldTx
}
