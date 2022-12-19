package sys

import (
	"context"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

type SystemServiceGetter interface {
	SystemService() SystemService
}

type SystemService interface {
	GetSystemState(ctx context.Context) (*dto.Server, error)
}

type systemService struct{}

func NewSystemService() SystemService {
	return &systemService{}
}

var _ SystemService = &systemService{}

func (s *systemService) GetSystemState(ctx context.Context) (*dto.Server, error) {
	return utils.InitServer()
}
