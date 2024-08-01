package router

import (
	"github.com/gin-gonic/gin"
	"go-learning/3-gorm/api"
	"go-learning/3-gorm/web"
	"net/http"
)

func MyRouter(engine *gin.Engine) {
	// 设置自定义404页面
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 - 请求地址无效!")
	})

	apiRouter(engine)

	webRouter(engine)
}

func apiRouter(engine *gin.Engine) {
	// user module
	engine.POST("/api/user/add", api.AddUser)
	engine.GET("/api/user/get", api.GetUser)
	engine.POST("/api/user/update_info", api.UpdateUserInfo)
	engine.POST("/api/user/update_status", api.UpdateUserStatus)
}

func webRouter(engine *gin.Engine) {
	// 加载模版页面
	engine.LoadHTMLGlob("3-gorm/template/*")

	// login module
	engine.GET("/", web.Login)
	engine.GET("/page/login", web.Login)
	engine.POST("/page/doLogin", web.DoLogin)
}
