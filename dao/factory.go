package dao

import (
	"github.com/noovertime7/kubemanage/dao/workflow"
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
	WorkFlow() workflow.WorkFlowInterface
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
