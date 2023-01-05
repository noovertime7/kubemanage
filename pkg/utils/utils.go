package utils

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GormExist(err error) bool {
	return !errors.Is(gorm.ErrRecordNotFound, err)
}
