package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		if AlwaysAllowPath.Has(c.Request.URL.Path) {
			return
		}
		// GET 请求不记录
		if c.Request.Method == http.MethodGet {
			return
		}
		var (
			err    error
			userId int
			body   []byte
		)
		//如果请求不是get请求，从body中获取数据
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				logger.LG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			if len(m) > 0 {
				body, err = json.Marshal(&m)
				if err != nil {
					logger.LG.Error("marshal body error:", zap.Error(err))
					body = []byte{}
				}
			}
		}
		claims, err := utils.GetClaims(c)
		if err != nil {
			logger.LG.Error("get claims from token err:", zap.Error(err))
			return
		}
		userId = claims.ID
		record := model.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		//if len(record.Body) > 1024 {
		//	// 截断
		//	newBody := respPool.Get().([]byte)
		//	copy(newBody, record.Body)
		//	record.Body = string(newBody)
		//	defer respPool.Put(newBody[:0])
		//}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()
		//
		//if len(record.Resp) > 1024 {
		//	// 截断
		//	newBody := respPool.Get().([]byte)
		//	copy(newBody, record.Resp)
		//	record.Resp = string(newBody)
		//	defer respPool.Put(newBody[:0])
		//}

		if err := v1.CoreV1.System().Operation().CreateOperationRecord(c, &record); err != nil {
			logger.LG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
