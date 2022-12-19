package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg/globalError"
	"github.com/pkg/errors"
	"net/http"
)

type ResponseCode int

type Response struct {
	Code    ResponseCode `json:"code"`
	Msg     string       `json:"msg"`
	RealErr string       `json:"real_err"`
	Data    interface{}  `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: http.StatusOK, Msg: "", Data: data}
	tempMsg, ok := data.(string)
	if ok {
		resp.Msg = tempMsg
	}
	if ok && tempMsg == "" {
		resp.Msg = "操作成功"
		resp.Data = "操作成功"
	}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

func ResponseError(c *gin.Context, err error) {
	//判断错误类型
	// As - 获取错误的具体实现
	var code ResponseCode
	var myError = new(globalError.GlobalError)
	if errors.As(err, &myError) {
		code = ResponseCode(myError.Code)
	}
	resp := &Response{Code: code, Msg: err.Error(), RealErr: myError.RealErrorMessage, Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
