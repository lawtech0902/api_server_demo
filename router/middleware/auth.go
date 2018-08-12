package middleware

import (
	"github.com/gin-gonic/gin"
	"go_projects/api_server/pkg/token"
	"go_projects/api_server/handler"
	"go_projects/api_server/pkg/errno"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午9:44'
*/

// 用户认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析jwt
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			return
		}

		c.Next()
	}
}
