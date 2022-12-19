package v1

import (
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/sys"
)

type SystemGetter interface {
	System() SystemInterface
}

// SystemInterface 顶层抽象 包括系统相关接口
type SystemInterface interface {
	sys.UserServiceGetter
	sys.MenuGetter
	sys.CasbinServiceGetter
	sys.AuthorityGetter
	sys.OperationServiceGetter
	sys.APIServiceGetter
	sys.SystemServiceGetter
	sys.DepartmentServiceGetter
}

var _ SystemInterface = &system{}

func NewSystem(app *KubeManage) SystemInterface {
	return &system{app: app, factory: app.Factory}
}

type system struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func (s *system) User() sys.UserService {
	return sys.NewUserService(s.factory)
}

func (s *system) Menu() sys.MenuService {
	return sys.NewMenuService(s.factory)
}

func (s *system) CasbinService() sys.CasbinService {
	return sys.NewCasbinService(s.factory)
}

func (s *system) Authority() sys.Authority {
	return sys.NewAuthority(s.factory)
}

func (s *system) Operation() sys.OperationService {
	return sys.NewOperationService(s.factory)
}

func (s *system) Api() sys.APIService {
	return sys.NewApiService(s.factory)
}

func (s *system) SystemService() sys.SystemService {
	return sys.NewSystemService()
}

func (s *system) Department() sys.DepartmentService {
	return sys.NewDepartmentService(s.factory)
}
