package core

import (
	"github.com/gin-gonic/gin"
	"go-learning/2-api/common"
	"net/http"
)

func MyRouter(engine *gin.Engine) {
	// 设置自定义404页面
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 - 请求地址无效!")
	})

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, common.OfResult(400, "API无效", nil))
	})

	// user api
	engine.GET("/api/user/list", ListUser)      // query 参数
	engine.GET("/api/user/:userId", GetUser)    // path 参数
	engine.POST("/api/user/update", UpdateUser) // body 参数
}
