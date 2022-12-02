package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto"
)

type OperationServiceGetter interface {
	Operation() OperationService
}

type OperationService interface {
	CreateOperationRecord(ctx *gin.Context, record *model.SysOperationRecord) error
	DeleteRecord(ctx *gin.Context, id int) error
	GetPageList(ctx *gin.Context, in *dto.OperationListInput) (*dto.OperationListOutPut, error)
}

type operationService struct {
	app     *KubeManage
	factory dao.ShareDaoFactory
}

func NewOperationService(app *KubeManage) *operationService {
	return &operationService{app: app, factory: app.Factory}
}

func (o *operationService) CreateOperationRecord(ctx *gin.Context, record *model.SysOperationRecord) error {
	return o.factory.Opera().Save(ctx, record)
}

func (o *operationService) DeleteRecord(ctx *gin.Context, id int) error {
	record := &model.SysOperationRecord{ID: id}
	return o.factory.Opera().Delete(ctx, record)
}

func (o *operationService) GetPageList(ctx *gin.Context, in *dto.OperationListInput) (*dto.OperationListOutPut, error) {
	list, total, err := o.factory.Opera().PageList(ctx, in)
	if err != nil {
		return nil, err
	}
	return &dto.OperationListOutPut{OperationList: list, Total: total}, nil
}
