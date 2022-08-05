package dto

import nwV1 "k8s.io/api/networking/v1"

type IngressCreteInput struct {
	Name      string                 `json:"name"`
	NameSpace string                 `json:"namespace"`
	Label     map[string]string      `json:"label"`
	Hosts     map[string][]*HttpPath `json:"hosts"`
}

type HttpPath struct {
	Path        string        `json:"path"`
	PathType    nwV1.PathType `json:"path_type"`
	ServiceName string        `json:"service_name"`
	ServicePort int32         `json:"service_port"`
}
