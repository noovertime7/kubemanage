package model

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID       int           `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserName string        `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt     string        `json:"salt" gorm:"column:salt" description:"盐"`
	Password string        `json:"password" gorm:"column:password" description:"密码"`
	Status   sql.NullInt64 `json:"status" gorm:"column:status" description:"登录状态"`
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
		UserName: "admin",
		Salt:     "admin",
		Password: "29c09a3c055e47f704fb7c6df5b530e25f80ee3ab2a3ce44858284f929157389",
		Status:   sql.NullInt64{Int64: 0, Valid: true},
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
