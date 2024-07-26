package core

import (
	"github.com/gin-gonic/gin"
)

// MyRouteConfig web路由配置
func MyRouteConfig(engine *gin.Engine) {
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
