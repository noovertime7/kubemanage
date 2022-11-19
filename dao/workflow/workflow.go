package workflow

import (
	"context"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type WorkFlowInterface interface {
	Save(ctx context.Context, obj *model.Workflow) error
	Updates(ctx context.Context, obj *model.Workflow) error
	Find(ctx context.Context, id int) (*model.Workflow, error)
	FindList(ctx context.Context, search *model.Workflow) ([]*model.Workflow, error)
	PageList(ctx context.Context, params *kubernetes.WorkFlowListInput) ([]*model.Workflow, int, error)
	Delete(ctx context.Context, wid int) error
}

type workflow struct {
	db *gorm.DB
}

func NewWorkFlow(db *gorm.DB) WorkFlowInterface {
	return &workflow{db: db}
}

func (w *workflow) PageList(ctx context.Context, params *kubernetes.WorkFlowListInput) ([]*model.Workflow, int, error) {
	var total int64 = 0
	var list []*model.Workflow
	offset := (params.Page - 1) * params.Limit
	query := w.db.WithContext(ctx).Table(model.GetWorkflowTableName())
	query.Find(&list).Count(&total)
	if params.FilterName != "" {
		query = query.Where("( name like ?)", "%"+params.FilterName+"%")
	}
	if err := query.Limit(params.Limit).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}

func (w *workflow) Save(ctx context.Context, obj *model.Workflow) error {
	now := time.Now()
	obj.UpdatedAt = now
	return w.db.WithContext(ctx).Save(obj).Error
}

func (w *workflow) Updates(ctx context.Context, obj *model.Workflow) error {
	if obj.ID == 0 {
		return errors.New("id no set")
	}
	return w.db.WithContext(ctx).Updates(obj).Error
}

func (w *workflow) Find(ctx context.Context, id int) (*model.Workflow, error) {
	out := &model.Workflow{}
	return out, w.db.WithContext(ctx).Where("id = ?", id).First(out).Error
}

func (w *workflow) FindList(ctx context.Context, search *model.Workflow) ([]*model.Workflow, error) {
	var res []*model.Workflow
	return res, w.db.WithContext(ctx).Where(&search).Find(&res).Error
}

func (w *workflow) Delete(ctx context.Context, wid int) error {
	return w.db.WithContext(ctx).Where("id = ?", wid).Delete(&model.Workflow{}).Error
}
