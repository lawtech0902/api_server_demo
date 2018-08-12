package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午5:16'
*/

// 全局中间件
// 在请求的 Header 中插入 X-Request-Id（X-Request-Id 值为 32 位的 UUID，用于唯一标识一次 HTTP 请求）
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
