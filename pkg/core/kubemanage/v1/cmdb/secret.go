package cmdb

import (
	"context"
	"fmt"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/utils"
	"github.com/noovertime7/kubemanage/runtime"
	"gorm.io/gorm"
)

type SecretService interface {
	CreateSecret(ctx context.Context, in *dto.CMDBSecretCreateInput) error
	UpdateSecret(ctx context.Context, in *dto.CMDBSecretUpdateInput) error
	PageSecret(ctx context.Context, pager runtime.Pager) (dto.PageCMDBSecretOut, error)
	GetSecretList(ctx context.Context) ([]model.CMDBSecret, error)
	DeleteSecret(ctx context.Context, instanceID string) error
	DeleteSecrets(ctx context.Context, instanceIDs []string) error
}

func NewSecretService(factory dao.ShareDaoFactory) SecretService {
	return &serctService{factory: factory}
}

type serctService struct {
	factory dao.ShareDaoFactory
}

func (s *serctService) CreateSecret(ctx context.Context, in *dto.CMDBSecretCreateInput) error {
	// 判断是否存在重名
	var (
		enHostPassword string
		enPrivateKey   string
		err            error
	)
	secretDB, err := s.factory.CMDB().Secret().Find(ctx, model.CMDBSecret{Name: in.Name})
	if utils.GormExist(err) {
		return fmt.Errorf("存在相同名称认证信息，请重新填写")
	}

	if secretDB.Id != 0 {
		return fmt.Errorf("存在相同名称认证信息，请重新填写")
	}

	if in.HostPassword != "" {
		enHostPassword, err = utils.Encrypt([]byte(in.HostPassword))
		if err != nil {
			return err
		}
	}

	if in.PrivateKey != "" {
		enPrivateKey, err = utils.Encrypt([]byte(in.PrivateKey))
		if err != nil {
			return err
		}
	}
	secret := model.CMDBSecret{
		InstanceID:   utils.GetSnowflakeID(),
		Name:         in.Name,
		Protocol:     in.Protocol,
		SecretType:   in.SecretType,
		HostUserName: in.HostUserName,
		HostPassword: enHostPassword,
		Content:      in.Content,
		PrivateKey:   enPrivateKey,
	}
	return s.factory.CMDB().Secret().Save(ctx, &secret)
}

func (s *serctService) UpdateSecret(ctx context.Context, in *dto.CMDBSecretUpdateInput) error {
	secretDB, err := s.factory.CMDB().Secret().Find(ctx, model.CMDBSecret{InstanceID: in.InstanceID})
	if err != nil {
		return err
	}
	secret := &model.CMDBSecret{
		InstanceID:   in.InstanceID,
		Name:         in.Name,
		Protocol:     in.Protocol,
		SecretType:   in.SecretType,
		HostUserName: in.HostUserName,
		HostPassword: in.HostPassword,
		Content:      in.Content,
		PrivateKey:   in.PrivateKey,
	}
	// 表示密码发生变化
	if in.HostPassword != secretDB.HostPassword && in.HostPassword != "" {
		enHostPassword, err := utils.Encrypt([]byte(in.HostPassword))
		if err != nil {
			return err
		}
		secret.HostPassword = enHostPassword
	}
	// 秘钥发生变化
	if in.PrivateKey != secretDB.PrivateKey && in.PrivateKey != "" {
		enPrivateKey, err := utils.Encrypt([]byte(in.PrivateKey))
		if err != nil {
			return err
		}
		secret.PrivateKey = enPrivateKey
	}
	return s.factory.CMDB().Secret().Updates(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("instanceID = ?", secret.InstanceID)
	}, secret)
}

func (s *serctService) PageSecret(ctx context.Context, pager runtime.Pager) (dto.PageCMDBSecretOut, error) {
	list, total, err := s.factory.CMDB().Secret().PageList(ctx, pager)
	if err != nil {
		return dto.PageCMDBSecretOut{}, err
	}
	return dto.PageCMDBSecretOut{Total: total, List: list, Page: pager.GetPage(), PageSize: pager.GetPageSize()}, nil
}

func (s *serctService) GetSecretList(ctx context.Context) ([]model.CMDBSecret, error) {
	return s.factory.CMDB().Secret().FindList(ctx, model.CMDBSecret{})
}

func (s *serctService) DeleteSecret(ctx context.Context, instanceID string) error {
	return s.factory.CMDB().Secret().Delete(ctx, model.CMDBSecret{InstanceID: instanceID}, false)
}

func (s *serctService) DeleteSecrets(ctx context.Context, instanceIDs []string) error {
	for _, ins := range instanceIDs {
		if err := s.DeleteSecret(ctx, ins); err != nil {
			return err
		}
	}
	return nil
}
