package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

var IngressController ingressController

type ingressController struct{}

// CreateIngress 创建ingress
// ListPage godoc
// @Summary      创建ingress
// @Description  创建ingress
// @Tags         ingress
// @ID           /api/k8s/ingress/create
// @Accept       json
// @Produce      json
// @Param        body  body  kubernetes.IngressCreteInput  true  "body"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "新增成功}"
// @Router       /api/k8s/ingress/create [post]
func (i *ingressController) CreateIngress(ctx *gin.Context) {
	params := &kubernetes.IngressCreteInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Ingress.CreateIngress(params); err != nil {
		v1.Log.ErrorWithCode(globalError.CreateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.CreateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "新增成功")
}

// DeleteIngress 删除ingress
// ListPage godoc
// @Summary      删除ingress
// @Description  删除ingress
// @Tags         ingress
// @ID           /api/k8s/ingress/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "ingress名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/ingress/del [delete]
func (i *ingressController) DeleteIngress(ctx *gin.Context) {
	params := &kubernetes.IngressNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Ingress.DeleteIngress(params.NameSpace, params.Name); err != nil {
		v1.Log.ErrorWithCode(globalError.DeleteError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.DeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateIngress 更新ingress
// ListPage godoc
// @Summary      更新ingress
// @Description  更新ingress
// @Tags         ingress
// @ID           /api/k8s/ingress/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "ingress名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/ingress/update [put]
func (i *ingressController) UpdateIngress(ctx *gin.Context) {
	params := &kubernetes.IngressUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := kube.Ingress.UpdateIngress(params.NameSpace, params.Content); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetIngressList 查看ingress列表
// ListPage godoc
// @Summary      查看ingress列表
// @Description  查看ingress列表
// @Tags         ingress
// @ID           /api/k8s/ingress/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data":""  }"
// @Router       /api/k8s/ingress/list [get]
func (i *ingressController) GetIngressList(ctx *gin.Context) {
	params := &kubernetes.IngressListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Ingress.GetIngressList(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetIngressDetail 获取ingress详情
// ListPage godoc
// @Summary      获取ingress详情
// @Description  获取ingress详情
// @Tags         ingress
// @ID           /api/k8s/ingress/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "ingress名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":""  }"
// @Router       /api/k8s/ingress/detail [get]
func (i *ingressController) GetIngressDetail(ctx *gin.Context) {
	params := &kubernetes.IngressNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := kube.Ingress.GetIngressDetail(params.NameSpace, params.Name)
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetIngressNumPreNp 根据命名空间获取ingress数量
// ListPage godoc
// @Summary      根据命名空间获取ingress数量
// @Description  根据命名空间获取ingress数量
// @Tags         ingress
// @ID           /api/k8s/ingress/numnp
// @Accept       json
// @Produce      json
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data":"" }"
// @Router       /api/k8s/ingress/numnp [get]
func (i *ingressController) GetIngressNumPreNp(ctx *gin.Context) {
	data, err := kube.Ingress.GetIngressNp()
	if err != nil {
		v1.Log.ErrorWithCode(globalError.GetError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.GetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
