package model

type Workflow struct {
	CommonModel
	Name        string `json:"name" gorm:"column:name"`
	NameSpace   string `json:"namespace" gorm:"column:namespace"`
	Replicas    int32  `json:"replicas" gorm:"column:replicas"`
	Deployment  string `json:"deployment" gorm:"column:deployment"`
	Service     string `json:"service" gorm:"column:service"`
	Ingress     string `json:"ingress" gorm:"column:ingress"`
	ServiceType string `json:"service_type" gorm:"column:service_type"`
}

func (w *Workflow) TableName() string {
	return "t_workflow"
}

func GetWorkflowTableName() string {
	temp := &Workflow{}
	return temp.TableName()
}
