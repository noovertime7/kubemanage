package model

import (
	"context"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(DepartmentOrder, &Department{})
}

// TODO 是否删除多余字段

type Department struct {
	DeptId   int          `json:"deptId" gorm:"primaryKey;autoIncrement;"` //部门编码
	ParentId int          `json:"parentId" gorm:""`                        //上级部门
	DeptName string       `json:"deptName"  gorm:"size:128;"`              //部门名称
	Sort     int          `json:"sort" gorm:"size:4;"`                     //排序
	Leader   string       `json:"leader" gorm:"size:128;"`                 //负责人
	Status   int          `json:"status" gorm:"size:4;"`                   //状态
	Children []Department `json:"children" gorm:"-"`
	Users    []SysUser    `json:"users" gorm:"foreignKey:DeptId;references:DeptId"`
	CommonModel
}

func (d *Department) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(d)
}

func (d *Department) InitData(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(DepartmentInitData).Error
}

func (d *Department) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *Department
	// TODO 验证方式统一优化
	if err := db.WithContext(ctx).Where("dept_name = 'Kubemanage' ").Find(&out).Error; err != nil {
		return false, nil
	}
	return out.DeptId != 0, nil
}

func (d *Department) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&d)
}

func (*Department) TableName() string {
	return "sys_dept"
}
