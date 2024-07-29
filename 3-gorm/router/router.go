package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyRouter(engine *gin.Engine) {
	// 设置自定义404页面
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 - 请求地址无效!")
	})

	// 加载模版页面
	engine.LoadHTMLGlob("3-gorm/template/*")

	engine.GET("/", Login)
	engine.GET("/login", Login)
	engine.POST("/doLogin", DoLogin)

	// user api
	engine.POST("/api/user/add", AddUser)
	engine.GET("/api/user/get", GetUser)
	engine.POST("/api/user/update", UpdateUser)
}
