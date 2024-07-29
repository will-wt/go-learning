package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MyRouteConfig web路由配置
func MyRouteConfig(engine *gin.Engine) {
	// 设置自定义404页面
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 - 请求地址无效!")
	})

	// 加载模版页面
	engine.LoadHTMLGlob("1-hello/template/*")

	engine.GET("/", Hello)
	engine.POST("/doWelcome", DoWelcome)

	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code":    200,
			"message": "success",
		})
	})
}
