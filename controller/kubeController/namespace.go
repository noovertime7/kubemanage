package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubeDto"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

var NameSpace namespace

type namespace struct{}

// CreateNameSpace 创建namespace
// ListPage godoc
// @Summary      创建namespace
// @Description  创建namespace
// @Tags         NameSpace
// @ID           /api/k8s/namespace/create
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "创建成功}"
// @Router       /api/k8s/namespace/create [put]
func (n *namespace) CreateNameSpace(ctx *gin.Context) {
	params := &kubeDto.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.NameSpace.CreateNameSpace(params.Name); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "创建成功")
}

// DeleteNameSpace 删除namespace
// ListPage godoc
// @Summary      删除namespace
// @Description  删除namespace
// @Tags         NameSpace
// @ID           /api/k8s/namespace/del
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/namespace/del [delete]
func (n *namespace) DeleteNameSpace(ctx *gin.Context) {
	params := &kubeDto.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.NameSpace.DeleteNameSpace(params.Name); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// GetNameSpaceList 获取NameSpace列表
// ListPage godoc
// @Summary      获取NameSpace列表
// @Description  获取NameSpace列表
// @Tags         NameSpace
// @ID           /api/k8s/namespace/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": service.NameSpaceResp}"
// @Router       /api/k8s/namespace/list [get]
func (n *namespace) GetNameSpaceList(ctx *gin.Context) {
	params := &kubeDto.NameSpaceListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.NameSpace.GetNameSpaces(params.FilterName, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetNameSpaceDetail 获取NameSpace详情
// ListPage godoc
// @Summary      获取NameSpace详情
// @Description  获取NameSpace详情
// @Tags         NameSpace
// @ID           /api/k8s/namespace/detail
// @Accept       json
// @Produce      json
// @Param        name  query  string  true  "namespace名称"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":data }"
// @Router       /api/k8s/namespace/detail [get]
func (n *namespace) GetNameSpaceDetail(ctx *gin.Context) {
	params := &kubeDto.NameSpaceNameInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.NameSpace.GetNameSpacesDetail(params.Name)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
