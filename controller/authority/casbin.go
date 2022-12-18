package authority

import (
	"github.com/gin-gonic/gin"

	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/sys"
	"github.com/noovertime7/kubemanage/pkg/globalError"
)

// GetPolicyPathByAuthorityId
// @Tags      Authority
// @Summary   获取权限列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.CasbinInReceive                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  middleware.Response{data=dto.CasbinInfo,msg=string}  "获取权限列表,返回包括casbin详情列表"
// @Router    /api/authority/getPolicyPathByAuthorityId [get]
func (a *authorityController) GetPolicyPathByAuthorityId(ctx *gin.Context) {
	rule := &dto.CasbinInReceive{}
	if err := rule.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	middleware.ResponseSuccess(ctx, v1.CoreV1.System().CasbinService().GetPolicyPathByAuthorityId(rule.AuthorityId))
}

// UpdateCasbinByAuthorityId
// @Tags      Authority
// @Summary   通过角色更新接口权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.UpdateCasbinInput                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  middleware.Response{msg=string}  "通过角色更新接口权限"
// @Router    /api/authority/updateCasbinByAuthority [post]
func (a *authorityController) UpdateCasbinByAuthorityId(ctx *gin.Context) {
	params := &dto.UpdateCasbinInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		v1.Log.ErrorWithCode(globalError.ParamBindError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := v1.CoreV1.System().CasbinService().UpdateCasbin(params.AuthorityId, convertToCasbinRules(params.CasbinInfo)); err != nil {
		v1.Log.ErrorWithCode(globalError.UpdateError, err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.UpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

func convertToCasbinRules(info []dto.CasbinInfo) []sys.CasbinRule {
	out := make([]sys.CasbinRule, len(info))
	for i := range info {
		out[i] = info[i]
	}
	return out
}
