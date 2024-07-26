package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/1-hello/core"
)

func main() {
	// 初始化 gin 框架引擎
	engine := gin.Default()

	// 请求路由配置
	core.MyRouteConfig(engine)

	// 监听服务端口
	engine.Run(":7001")
}
