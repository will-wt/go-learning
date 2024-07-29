package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/3-gorm/config"
	"go-learning/3-gorm/global"
	"go-learning/3-gorm/initialize"
	"go-learning/3-gorm/router"
)

func main() {
	// 初始化 gin 框架引擎
	engine := gin.Default()

	// 配置请求路由
	router.MyRouter(engine)

	// 启动，并监听端口
	engine.Run(":7001")
}

// 初始化自启动函数
func init() {
	// 加载、解析配置项
	config.LoadConfig()
	fmt.Println("dsn: ", config.Config.Datasource.Dsn())

	global.DB = initialize.InitGorm() // 初始化数据库连接池
}
