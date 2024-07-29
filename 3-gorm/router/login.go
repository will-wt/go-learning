package router

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.html", nil)
}

func DoLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "admin" && password == "admin" {
		ctx.HTML(200, "welcome.html", gin.H{
			"username": username,
		})
	} else {
		ctx.HTML(200, "500.html", gin.H{
			"errorMessage": "用户名或密码错误",
		})
	}

}
