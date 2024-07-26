package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/2-login/router"
)

func main() {
	// 初始化 gin 框架引擎
	engine := gin.Default()

	// 配置请求路由
	router.MyRouter(engine)

	// 启动，并监听端口
	engine.Run(":7001")
}
