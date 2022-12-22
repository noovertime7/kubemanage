package model

import (
	"context"
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	ID           int            `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UUID         uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"` // 用户UUID
	DepartmentID uint           `json:"DeptId" gorm:"index;comment:部门ID"`
	UserName     string         `json:"userName" gorm:"index;comment:用户登录名"`                                               // 用户登录名
	Password     string         `json:"-"  gorm:"comment:用户登录密码"`                                                          // 用户登录密码
	NickName     string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                         // 用户昵称
	SideMode     string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                       // 用户侧边主题
	Avatar       string         `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor    string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                        // 基础颜色
	ActiveColor  string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                   // 活跃颜色
	AuthorityId  uint           `json:"authorityId" gorm:"default:2222;comment:用户角色ID"`                                    // 用户角色ID
	Authority    SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities  []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone        sql.NullString `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email        sql.NullString `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable       int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Status       sql.NullInt64  `gorm:"column:status;type:int(11);comment:0离线;1在线" json:"status"`
	CommonModel
}

func init() {
	RegisterInitializer(SysUserOrder, &SysUser{})
}

func (u *SysUser) TableName() string {
	return "sys_users"
}

func (u *SysUser) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := u.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	if err := db.WithContext(ctx).Create(SysUserEntities).Error; err != nil {
		return err
	}
	//管理员用户
	if err := db.Model(&SysUserEntities[0]).Association("Authorities").Replace(SysAuthorityEntities); err != nil {
		return err
	}
	//普通用户
	if err := db.Model(&SysUserEntities[1]).Association("Authorities").Replace(SysAuthorityEntities[:1]); err != nil {
		return err
	}
	return nil
}

// IsInitData 判断是否admin初始化成功
func (u *SysUser) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	adminOk, err := u.isAdminInit(ctx, db)
	if err != nil {
		return false, nil
	}
	return adminOk, nil
}

func (u *SysUser) isAdminInit(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *SysUser
	if err := db.WithContext(ctx).Where("user_name = 'admin' ").Find(&out).Error; err != nil {
		return false, err
	}
	return out.ID != 0, nil
}

func (u *SysUser) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&u)
}

func (u *SysUser) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(u)
}
