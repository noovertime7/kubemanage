package dao

import (
	"github.com/noovertime7/kubemanage/model"
	"gorm.io/gorm"
	"time"
)

type workflowResp struct {
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

func (w *Workflow) PageList(params *model.WorkFlowListInput) ([]Workflow, int, error) {
	var total int64 = 0
	var list []Workflow
	offset := (params.PageNo - 1) * params.PageSize
	query := Gorm
	query.Find(&list).Count(&total)
	query = query.Table(w.TableName()).Where("is_delete=0")
	if params.Info != "" {
		query = query.Where("( name like ?)", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
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
