package sys

import "github.com/noovertime7/kubemanage/dao"

type DepartmentService interface{}

type departmentService struct {
	factory dao.ShareDaoFactory
}
