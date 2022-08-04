package dto

type ServiceCreateInput struct {
	Name          string            `json:"name"`
	NameSpace     string            `json:"name_space"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Label         map[string]string `json:"label"`
}
