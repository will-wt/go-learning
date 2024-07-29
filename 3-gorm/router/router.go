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
	engine.GET("/api/listUser", ListUser)       // query 参数
	engine.GET("/api/getUser/:userId", GetUser) // path 参数
	engine.POST("/api/updateUser", UpdateUser)  // body 参数
}
