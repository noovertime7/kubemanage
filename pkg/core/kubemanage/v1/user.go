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
}

type userService struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func NewUserService(app *KubeManage) *userService {
	return &userService{app: app, factory: app.Factory}
}

var _ UserService = &userService{}

func (u userService) Login(ctx *gin.Context, userInfo *dto.AdminLoginInput) (string, error) {
	user, err := u.factory.User().Find(ctx, &model.UserModel{UserName: userInfo.UserName})
	if err != nil {
		return "", err
	}
	if !loginCheck(&checkInfo{salt: user.Salt, inputPwd: userInfo.Password, dbPwd: user.Password}) {
		return "", errors.New("密码错误,请重新输入")
	}
	token, err := pkg.JWTToken.GenerateToken(&user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u userService) LoginOut(ctx *gin.Context, uid int) error {
	user := &model.UserModel{ID: uid, Status: sql.NullInt64{Int64: 0, Valid: true}}
	return u.factory.User().Updates(ctx, user)
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
