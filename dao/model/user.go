package model

import "database/sql"

type UserModel struct {
	ID       int           `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserName string        `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt     string        `json:"salt" gorm:"column:salt" description:"盐"`
	Password string        `json:"password" gorm:"column:password" description:"密码"`
	Status   sql.NullInt64 `json:"status" gorm:"column:status" description:"登录状态"`
	CommonModel
}

func (u *UserModel) TableName() string {
	return "t_user"
}
