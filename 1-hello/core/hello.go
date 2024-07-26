package core

import (
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	ctx.HTML(200, "hello.html", nil)
}

func DoWelcome(ctx *gin.Context) {
	message := ctx.PostForm("message")

	if message != "" {
		ctx.HTML(200, "welcome.html", gin.H{
			"message": message,
		})
	} else {
		ctx.HTML(200, "500.html", gin.H{
			"message": "用户名或密码错误",
		})
	}

}
