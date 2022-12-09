package kube

import (
	"context"
	"encoding/json"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Secret secret

type secret struct {
}

type SecretResp struct {
	Total int             `json:"total"`
	Items []coreV1.Secret `json:"items"`
}

type SecretNp struct {
	NameSpace string `json:"namespace"`
	SecretNum int    `json:"secret_num"`
}

func (d *secret) toCells(secrets []coreV1.Secret) []DataCell {
	cells := make([]DataCell, len(secrets))
	for i := range secrets {
		cells[i] = secretCell(secrets[i])
	}
	return cells
}

func (d *secret) FromCells(cells []DataCell) []coreV1.Secret {
	secrets := make([]coreV1.Secret, len(cells))
	for i := range cells {
		secrets[i] = coreV1.Secret(cells[i].(secretCell))
	}
	return secrets
}

func (d *secret) GetSecrets(filterName, namespace string, limit, page int) (*SecretResp, error) {
	SecretsList, err := K8s.ClientSet.CoreV1().Secrets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	selectableData := &dataSelector{
		GenericDataList: d.toCells(SecretsList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginatite: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filterd := selectableData.Filter()
	total := len(filterd.GenericDataList)
	data := filterd.Sort().Paginate()
	secrets := d.FromCells(data.GenericDataList)
	return &SecretResp{
		Total: total,
		Items: secrets,
	}, nil
}

func (d *secret) GetSecretsDetail(name, namespace string) (*coreV1.Secret, error) {
	data, err := K8s.ClientSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *secret) DeleteSecrets(name, namespace string) error {
	return K8s.ClientSet.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (d *secret) UpdateSecrets(content, namespace string) error {
	var secret = &coreV1.Secret{}
	if err := json.Unmarshal([]byte(content), secret); err != nil {
		return err
	}
	if _, err := K8s.ClientSet.CoreV1().Secrets(namespace).Update(context.TODO(), secret, metaV1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}
