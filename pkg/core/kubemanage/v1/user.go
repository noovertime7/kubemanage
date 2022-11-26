package v1

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
}

type userService struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func NewUserService(app *KubeManage) *userService {
	return &userService{app: app, factory: app.Factory}
}

var _ UserService = &userService{}

func (u *userService) Login(ctx *gin.Context, userInfo *dto.AdminLoginInput) (string, error) {
	user, err := u.factory.User().Find(ctx, &model.SysUser{UserName: userInfo.UserName})
	if err != nil {
		return "", err
	}
	if !loginCheck(&checkInfo{salt: pkg.Salt, inputPwd: userInfo.Password, dbPwd: user.Password}) {
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

type checkInfo struct {
	salt     string
	inputPwd string
	dbPwd    string
}

func loginCheck(info *checkInfo) bool {
	encryptInputPwd := pkg.GenSaltPassword(info.salt, info.inputPwd)
	return encryptInputPwd == info.dbPwd
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
	menus, err := CoreV1.Menu().GetMenu(ctx, aid)
	if err != nil {
		return nil, err
	}
	var outRules []string
	rules := CoreV1.CasbinService().GetPolicyPathByAuthorityId(aid)
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
	check := &checkInfo{
		salt:     pkg.Salt,
		inputPwd: info.OldPwd,
		dbPwd:    user.Password,
	}
	if !loginCheck(check) {
		return errors.New("原密码错误,请重新输入")
	}
	//生成新密码
	user.Password = pkg.GenSaltPassword(pkg.Salt, info.NewPwd)
	return u.factory.User().Updates(ctx, user)
}

func (u *userService) ResetPassword(ctx *gin.Context, uid int) error {
	user := &model.SysUser{ID: uid, Password: pkg.GenSaltPassword(pkg.Salt, "kubemanage")}
	return u.factory.User().Updates(ctx, user)
}
