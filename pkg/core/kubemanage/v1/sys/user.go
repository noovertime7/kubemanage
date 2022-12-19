package sys

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg"
	"github.com/pkg/errors"
)

type UserServiceGetter interface {
	User() UserService
}

type UserService interface {
	Login(ctx *gin.Context, userInfo *dto.AdminLoginInput) (string, error)
	LoginOut(ctx *gin.Context, uid int) error
	GetUserInfo(ctx *gin.Context, uid int, aid uint) (*dto.UserInfoOut, error)
	SetUserAuth(ctx *gin.Context, uid int, aid uint) error
	DeleteUser(ctx *gin.Context, uid int) error
	ChangePassword(ctx *gin.Context, uid int, info *dto.ChangeUserPwdInput) error
	ResetPassword(ctx *gin.Context, uid int) error

	PageList(ctx *gin.Context, did uint, info dto.PageUsersIn) (dto.PageUsers, error)
}

type userService struct {
	Menu    MenuService
	Casbin  CasbinService
	factory dao.ShareDaoFactory
}

func NewUserService(factory dao.ShareDaoFactory) UserService {
	return &userService{
		factory: factory,
		Menu:    NewMenuService(factory),
		Casbin:  NewCasbinService(factory),
	}
}

var _ UserService = &userService{}

func (u *userService) Login(ctx *gin.Context, userInfo *dto.AdminLoginInput) (string, error) {
	user, err := u.factory.User().Find(ctx, &model.SysUser{UserName: userInfo.UserName})
	if err != nil {
		return "", err
	}

	if !pkg.CheckPassword(userInfo.Password, user.Password) {
		return "", errors.New("密码错误,请重新输入")
	}

	token, err := pkg.JWTToken.GenerateToken(pkg.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.UserName,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *userService) LoginOut(ctx *gin.Context, uid int) error {
	user := &model.SysUser{ID: uid, Status: sql.NullInt64{Int64: 0, Valid: true}}
	return u.factory.User().Updates(ctx, user)
}

func (u *userService) GetUserInfo(ctx *gin.Context, uid int, aid uint) (*dto.UserInfoOut, error) {
	user, err := u.factory.User().Find(ctx, &model.SysUser{ID: uid})
	if err != nil {
		return nil, err
	}
	menus, err := u.Menu.GetMenuByAuthorityID(ctx, aid)
	if err != nil {
		return nil, err
	}
	var outRules []string
	rules := u.Casbin.GetPolicyPathByAuthorityId(aid)
	for _, rule := range rules {
		item := rule.Path + "," + rule.Method
		outRules = append(outRules, item)
	}
	return &dto.UserInfoOut{
		User:      *user,
		Menus:     menus,
		RuleNames: outRules,
	}, nil
}

func (u *userService) SetUserAuth(ctx *gin.Context, uid int, aid uint) error {
	user := &model.SysUser{ID: uid, AuthorityId: aid}
	return u.factory.User().Updates(ctx, user)
}

func (u *userService) DeleteUser(ctx *gin.Context, uid int) error {
	user := &model.SysUser{ID: uid}
	return u.factory.User().Delete(ctx, user)
}

func (u *userService) ChangePassword(ctx *gin.Context, uid int, info *dto.ChangeUserPwdInput) error {
	userDB := &model.SysUser{ID: uid}
	user, err := u.factory.User().Find(ctx, userDB)
	if err != nil {
		return err
	}

	if !pkg.CheckPassword(info.OldPwd, user.Password) {
		return errors.New("原密码错误,请重新输入")
	}

	//生成新密码
	user.Password, err = pkg.GenSaltPassword(info.NewPwd)
	if err != nil {
		return err
	}
	return u.factory.User().Updates(ctx, user)
}

func (u *userService) ResetPassword(ctx *gin.Context, uid int) error {
	newPwd, err := pkg.GenSaltPassword("kubemanage")
	if err != nil {
		return err
	}
	user := &model.SysUser{ID: uid, Password: newPwd}
	return u.factory.User().Updates(ctx, user)
}

func (u *userService) PageList(ctx *gin.Context, did uint, info dto.PageUsersIn) (dto.PageUsers, error) {
	users, total, err := u.factory.User().PageList(ctx, did, info)
	if err != nil {
		return dto.PageUsers{}, err
	}
	var out []dto.PageUserItem
	for _, user := range users {
		dept, err := u.factory.Department().Find(ctx, &model.Department{DeptId: user.DepartmentID})
		if err != nil {
			return dto.PageUsers{}, err
		}
		outItem := dto.PageUserItem{
			ID:             user.ID,
			DepartmentID:   user.DepartmentID,
			DepartmentName: dept.DeptName,
			UserName:       user.UserName,
			NickName:       user.NickName,
			Authorities:    user.Authorities,
			Phone:          user.Phone,
			Email:          user.Email,
			Enable:         user.Enable,
			Status:         user.Status.Int64,
		}
		out = append(out, outItem)
	}
	return dto.PageUsers{
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
		List:     out,
	}, nil
}
