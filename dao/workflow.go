package dao

import (
	"github.com/noovertime7/kubemanage/dto"
	"gorm.io/gorm"
	"time"
)

type WorkflowResp struct {
	Items []*Workflow
	Total int
}
type Workflow struct {
	ID          uint       `json:"id" gorm:"pk"`
	Name        string     `json:"name"`
	NameSpace   string     `json:"namespace"`
	Replicas    int32      `json:"replicas"`
	Deployment  string     `json:"deployment"`
	Service     string     `json:"service"`
	Ingress     string     `json:"ingress"`
	ServiceType string     `json:"service_type" gorm:"column:service_type"`
	IsDeleted   uint       `json:"is_deleted"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (w *Workflow) TableName() string {
	return "t_workflow"
}

func (w *Workflow) PageList(params *dto.WorkFlowListInput) ([]*Workflow, int, error) {
	var total int64 = 0
	var list []*Workflow
	offset := (params.Page - 1) * params.Limit
	query := Gorm
	query.Find(&list).Count(&total)
	query = query.Table(w.TableName()).Where("is_deleted=0")
	if params.FilterName != "" {
		query = query.Where("( name like ?)", "%"+params.FilterName+"%")
	}
	if err := query.Limit(params.Limit).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, int(total), nil
}

func (w *Workflow) Save() error {
	return Gorm.Save(w).Error
}

func (w *Workflow) Find(search *Workflow) (*Workflow, error) {
	out := &Workflow{}
	err := Gorm.Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (w *Workflow) DeleteById() error {
	w.IsDeleted = 1
	return Gorm.Table(w.TableName()).Where("id = ?", w.ID).Updates(map[string]interface{}{
		"status":     w.IsDeleted,
		"deleted_at": time.Now(),
	}).Error
}
