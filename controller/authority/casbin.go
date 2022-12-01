package authority

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto"
	"github.com/noovertime7/kubemanage/middleware"
	v1 "github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/wonderivan/logger"
)

// GetPolicyPathByAuthorityId
// @Tags      Casbin
// @Summary   获取权限列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      dto.CasbinInReceive                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  middleware.Response{data=dto.CasbinInfo,msg=string}  "获取权限列表,返回包括casbin详情列表"
// @Router    /casbin/getPolicyPathByAuthorityId [post]
func (c *casbinController) GetPolicyPathByAuthorityId(ctx *gin.Context) {
	rule := &dto.CasbinInReceive{}
	if err := rule.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败", err.Error())
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	fmt.Println("rule = ", rule)
	middleware.ResponseSuccess(ctx, v1.CoreV1.CasbinService().GetPolicyPathByAuthorityId(rule.AuthorityId))
}
