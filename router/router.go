package router

import (
	"github.com/gin-gonic/gin"
	"go_projects/api_server/handler/sd"
	"go_projects/api_server/handler/user"
	"go_projects/api_server/router/middleware"
	"net/http"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/9 下午6:07'
*/

// Load 导入 middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// g.Use()为每一个请求设置header
	g.Use(gin.Recovery())     // 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(middleware.NoCache) // 强制浏览器不使用缓存
	g.Use(middleware.Options) // 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Secure)  // 一些安全设置
	g.Use(mw...)

	// 404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	// 用户路由分组
	u := g.Group("/v1/user")
	{
		u.POST("/:username", user.Create)
	}

	// 健康检查路由分组，类似Flask蓝图形式，加前缀
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
