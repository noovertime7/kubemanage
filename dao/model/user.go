package model

import (
	"context"
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID          int           `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UUID        uuid.UUID     `json:"uuid" gorm:"index;comment:用户UUID"`                                                  // 用户UUID
	UserName    string        `json:"userName" gorm:"index;comment:用户登录名"`                                               // 用户登录名
	Password    string        `json:"-"  gorm:"comment:用户登录密码"`                                                          // 用户登录密码
	NickName    string        `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                         // 用户昵称
	SideMode    string        `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                       // 用户侧边主题
	Avatar      string        `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string        `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                        // 基础颜色
	ActiveColor string        `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                   // 活跃颜色
	AuthorityId uint          `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                     // 用户角色ID
	Authority   SysAuthority  `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Phone       string        `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email       string        `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable      int           `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Status      sql.NullInt64 `gorm:"column:status;type:int(11);comment:0离线;1在线" json:"status"`
	CommonModel
}

func init() {
	RegisterInitializer(&UserModel{})
}

func (u *UserModel) TableName() string {
	return "t_user"
}

func (u *UserModel) InitData(ctx context.Context, db *gorm.DB) error {
	// 初始化admin
	datas := []*UserModel{{
		UUID:        uuid.NewV4(),
		UserName:    "admin",
		Password:    "29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389",
		NickName:    "admin",
		SideMode:    "dark",
		Avatar:      "https://qmplusimg.henrongyi.top/gva_header.jpg",
		BaseColor:   "#fff",
		ActiveColor: "#1890ff",
		AuthorityId: 1,
		Phone:       "12345678901",
		Email:       "test@qq.com",
		Enable:      1,
		Status:      sql.NullInt64{Int64: 0, Valid: true},
		CommonModel: CommonModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}}
	for _, data := range datas {
		if err := db.WithContext(ctx).Create(data).Error; err != nil {
			return err
		}
	}
	return nil
}

// IsInitData 判断是否admin初始化成功
func (u *UserModel) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	adminOk, err := u.isAdminInit(ctx, db)
	if err != nil {
		return false, err
	}
	return adminOk, nil
}

func (u *UserModel) isAdminInit(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *UserModel
	if err := db.WithContext(ctx).Table("t_user").Where("user_name = 'admin' ").Find(&out).Error; err != nil {
		return false, err
	}
	return out.ID != 0, nil
}

func (u *UserModel) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&u)
}

func (u *UserModel) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(u)
}
