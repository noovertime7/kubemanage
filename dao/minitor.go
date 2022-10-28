package dao

import (
	"time"
)

type K8SDB struct {
	Id        int       `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Config    string    `gorm:"column:config;type:text" json:"config"`
	IsDeleted int       `gorm:"column:is_deleted;type:int(11);NOT NULL" json:"is_deleted"`
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;NOT NULL" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;NOT NULL" json:"update_at"`
}

func (k *K8SDB) TableName() string {
	return "t_k8s"
}

func (k *K8SDB) Find(search *K8SDB) (*K8SDB, error) {
	out := &K8SDB{}
	return out, Gorm.Where(search).Find(out).Error
}

func (k *K8SDB) Save() error {
	return Gorm.Save(k).Error
}
