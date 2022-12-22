package sys

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"strconv"

	"github.com/noovertime7/kubemanage/dao"
	"github.com/noovertime7/kubemanage/dao/model"
)

type DepartmentServiceGetter interface {
	Department() DepartmentService
}

type DepartmentService interface {
	GetDepartmentTree(ctx context.Context) ([]model.Department, error)
	GetDeptUsers(ctx context.Context, did uint) ([]model.SysUser, error)
	PageList(ctx *gin.Context, info dto.PageListDeptInput) (dto.PageListDeptOut, error)
}

type departmentService struct {
	factory dao.ShareDaoFactory
}

func NewDepartmentService(factory dao.ShareDaoFactory) DepartmentService {
	return &departmentService{factory: factory}
}

func (d *departmentService) PageList(ctx *gin.Context, info dto.PageListDeptInput) (dto.PageListDeptOut, error) {
	list, total, err := d.factory.Department().PageList(ctx, &info)
	if err != nil {
		return dto.PageListDeptOut{}, err
	}
	return dto.PageListDeptOut{
		Total: total,
		List:  list,
	}, err
}

func (d *departmentService) GetDeptUsers(ctx context.Context, did uint) ([]model.SysUser, error) {
	in := &model.Department{DeptId: did}
	dept, err := d.factory.Department().FindListWithUsers(ctx, in)
	if err != nil {
		return nil, err
	}
	return dept.SysUsers, err
}

func (d *departmentService) GetDepartmentTree(ctx context.Context) ([]model.Department, error) {
	treeMap, err := d.getDepartmentTreeMap(ctx)
	if err != nil {
		return nil, err
	}
	deployments := treeMap["0"]
	for i := 0; i < len(deployments); i++ {
		if err := d.getDeptChildrenList(&deployments[i], treeMap); err != nil {
			return nil, err
		}
	}
	return deployments, nil
}

func (d *departmentService) getDepartmentTreeMap(ctx context.Context) (treeMap map[string][]model.Department, err error) {
	var department *model.Department
	treeMap = make(map[string][]model.Department)
	departments, err := d.factory.Department().FindList(ctx, department)
	if err != nil {
		return nil, err
	}
	for _, v := range departments {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

func (d *departmentService) getDeptChildrenList(department *model.Department, treeMap map[string][]model.Department) (err error) {
	department.Children = treeMap[strconv.Itoa(int(department.DeptId))]
	for i := 0; i < len(department.Children); i++ {
		err = d.getDeptChildrenList(&department.Children[i], treeMap)
	}
	return err
}
