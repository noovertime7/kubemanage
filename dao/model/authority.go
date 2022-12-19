package model

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func init() {
	RegisterInitializer(SysAuthorityOrder, &SysAuthority{})
}

type SysAuthority struct {
	CommonModel
	AuthorityId     uint            `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        uint            `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (s *SysAuthority) MigrateTable(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).AutoMigrate(&s)
}

func (s *SysAuthority) InitData(ctx context.Context, db *gorm.DB) error {
	ok, err := s.IsInitData(ctx, db)
	if err != nil || ok {
		return err
	}
	if err := db.Model(&SysAuthorityEntities[0]).Association("DataAuthorityId").Replace(
		[]*SysAuthority{
			{AuthorityId: 111},
			{AuthorityId: 222},
			{AuthorityId: 2221},
		}); err != nil {
		return errors.Wrapf(err, "%s表数据初始化失败!",
			db.Model(&SysAuthorityEntities[0]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}
	if err := db.Model(&SysAuthorityEntities[1]).Association("DataAuthorityId").Replace(
		[]*SysAuthority{
			{AuthorityId: 222},
			{AuthorityId: 2221},
		}); err != nil {
		return errors.Wrapf(err, "%s表数据初始化失败!",
			db.Model(&SysAuthorityEntities[1]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}
	return nil
}

func (s *SysAuthority) IsInitData(ctx context.Context, db *gorm.DB) (bool, error) {
	var out *SysAuthority
	if err := db.WithContext(ctx).Where("authority_id = '111' ").Find(&out).Error; err != nil {
		return false, nil
	}
	return out.AuthorityId != 0, nil
}

func (s *SysAuthority) TableCreated(ctx context.Context, db *gorm.DB) bool {
	return db.WithContext(ctx).Migrator().HasTable(&s)
}

func (*SysAuthority) TableName() string {
	return "sys_authorities"
}
