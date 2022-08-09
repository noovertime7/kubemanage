package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: http.StatusOK, Msg: "", Data: data}
	tempMsg, ok := data.(string)
	if ok {
		resp.Msg = tempMsg
	}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

func ResponseError(c *gin.Context, code int, err error) {
	resp := &Response{Code: code, Msg: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
