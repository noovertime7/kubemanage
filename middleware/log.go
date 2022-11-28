package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/pkg/logger"
	"go.uber.org/zap"
	"time"
)

// Logger 接收gin框架默认的日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.LG.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("host", c.Request.Host),
			zap.Duration("cost", cost),
		)
	}
}
