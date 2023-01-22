package sys

import (
	"fmt"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg"
)

const (
	activeUserID = 1
	lockUserID   = 2
	lockAction   = "lock"
	unlockAction = "unlock"
)

type UserServiceGetter interface {
	User() UserService
}

type UserService interface {
	Login(ctx *gin.Context, userInfo *dto.AdminLoginInput) (string, error)
	LoginOut(ctx *gin.Context, uid int) error
	GetUserInfo(ctx *gin.Context, uid int, aid uint) (*dto.UserInfoOut, error)
	RegisterUser(ctx *gin.Context, userInfo dto.UserInfoInput) error
	SetUserAuth(ctx *gin.Context, uid int, auths []uint) error
	UpdateUser(ctx *gin.Context, userInfo dto.UserInfoInput) error
	LockUser(ctx *gin.Context, uid int, action string) error
	DeleteUser(ctx *gin.Context, uid int) error
	DeleteUsers(ctx *gin.Context, uids []int) error
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

	if user.Enable == lockUserID {
		return "", errors.New("用户已被锁定")
	}

	if !pkg.CheckPassword(userInfo.Password, user.Password) {
		return "", errors.New("密码错误,请重新输入")
	}

	// 登录成功 修改登录状态
	user.Status = sql.NullInt64{
		Int64: 1,
		Valid: true,
	}

	u.factory.Begin()
	defer u.factory.Commit()

	if err := u.factory.User().Updates(ctx, user); err != nil {
		u.factory.Rollback()
		return "", err
	}

	token, err := pkg.JWTToken.GenerateToken(pkg.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.UserName,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	})
	if err != nil {
		u.factory.Rollback()
		return "", err
	}
	return token, nil
}

func (u *userService) LoginOut(ctx *gin.Context, uid int) error {
	user := &model.SysUser{ID: uid, Status: sql.NullInt64{Int64: 2, Valid: true}}
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

// SetUserAuth 绑定用户角色
func (u *userService) SetUserAuth(ctx *gin.Context, uid int, auths []uint) error {
	var authorities []model.SysAuthority
	for _, aid := range auths {
		auth := model.SysAuthority{AuthorityId: aid}
		authorities = append(authorities, auth)
	}
	user := &model.SysUser{ID: uid}
	return u.factory.User().ReplaceAuthorities(ctx, user, authorities)
}

func (u *userService) DeleteUser(ctx *gin.Context, uid int) error {
	userInfo := &model.SysUser{ID: uid}
	user, err := u.factory.User().Find(ctx, userInfo)
	if err != nil {
		return err
	}
	// 清理角色菜单绑定关系
	if err := u.factory.User().RemoveAuthorities(ctx, user, user.Authorities); err != nil {
		return err
	}

	if err := u.factory.User().Delete(ctx, userInfo); err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteUsers(ctx *gin.Context, uids []int) error {
	for _, uid := range uids {
		if err := u.DeleteUser(ctx, uid); err != nil {
			return err
		}
	}
	return nil
}

// LockUser 锁定或解锁定用户
func (u *userService) LockUser(ctx *gin.Context, uid int, action string) error {
	switch action {
	case lockAction:
		user := &model.SysUser{ID: uid, Enable: lockUserID}
		return u.factory.User().Updates(ctx, user)
	case unlockAction:
		user := &model.SysUser{ID: uid, Enable: activeUserID}
		return u.factory.User().Updates(ctx, user)
	default:
		return fmt.Errorf("expect action %s or %s ,but got %s ", lockAction, unlockAction, action)
	}
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
			Phone:          user.Phone.String,
			Email:          user.Email.String,
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

func (u *userService) RegisterUser(ctx *gin.Context, userInfo dto.UserInfoInput) error {
	// 查询用户名是否存在
	tempUser, err := u.factory.User().Find(ctx, &model.SysUser{UserName: userInfo.UserName})
	if err != nil {
		return err
	}
	if tempUser.ID != 0 {
		return fmt.Errorf("%s已注册", userInfo.UserName)
	}
	// 添加用户
	encryptPwd, err := pkg.GenSaltPassword(userInfo.Password)
	if err != nil {
		return err
	}
	user := &model.SysUser{
		UUID:         uuid.NewV4(),
		DepartmentID: userInfo.DepartmentID,
		UserName:     userInfo.UserName,
		Password:     encryptPwd,
		NickName:     userInfo.NickName,
		SideMode:     "dark",
		Avatar:       "https://qmplusimg.henrongyi.top/gva_header.jpg",
		BaseColor:    "#fff",
		ActiveColor:  "#1890ff",
		AuthorityId:  userInfo.AuthorityId,
		Phone:        sql.NullString{String: userInfo.Phone, Valid: true},
		Email:        sql.NullString{String: userInfo.Email, Valid: true},
		Enable:       userInfo.Enable,
		Status:       sql.NullInt64{Int64: 2, Valid: true},
	}
	if err := u.factory.User().Save(ctx, user); err != nil {
		return err
	}
	// 设置用户权限
	var auths []model.SysAuthority
	for _, aid := range userInfo.Authorities {
		auth := model.SysAuthority{AuthorityId: aid}
		auths = append(auths, auth)
	}
	return u.factory.User().ReplaceAuthorities(ctx, user, auths)
}

// UpdateUser 更新用户信息
func (u *userService) UpdateUser(ctx *gin.Context, userInfo dto.UserInfoInput) error {
	// 查询用户名是否存在
	tempUser, err := u.factory.User().Find(ctx, &model.SysUser{UserName: userInfo.UserName})
	if err != nil {
		return err
	}
	user := &model.SysUser{
		ID:           tempUser.ID,
		DepartmentID: userInfo.DepartmentID,
		NickName:     userInfo.NickName,
		Phone:        sql.NullString{String: userInfo.Phone, Valid: true},
		Email:        sql.NullString{String: userInfo.Email, Valid: true},
		Enable:       userInfo.Enable,
	}
	if err := u.factory.User().Updates(ctx, user); err != nil {
		return err
	}
	// 设置用户权限
	var auths []model.SysAuthority
	for _, aid := range userInfo.Authorities {
		auth := model.SysAuthority{AuthorityId: aid}
		auths = append(auths, auth)
	}
	return u.factory.User().ReplaceAuthorities(ctx, user, auths)
}
