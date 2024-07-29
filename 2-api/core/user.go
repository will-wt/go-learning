package core

import (
	"github.com/gin-gonic/gin"
	"go-learning/2-api/common"
	"net/http"
	"strconv"
)

// User 定义用户结构体
type User struct {
	UserId   uint32 `json:"userId"`
	Nickname string `json:"name"`
	Age      uint8  `json:"age"`
	Address  string `json:"address"`
}

func ListUser(context *gin.Context) {
	corpId := context.Query("corpId")
	status, _ := strconv.Atoi(context.Query("status"))

	if corpId == "" {
		context.JSON(200, gin.H{
			"code":    406,
			"message": "parameter missing",
		})
		return
	}

	if status != 1 {
		context.JSON(200, common.Result{
			Code:    200,
			Message: "success",
			Data: []User{
				{UserId: 1, Nickname: "ZhangSan", Age: 20},
				{UserId: 2, Nickname: "LiSi", Age: 30},
				{UserId: 3, Nickname: "WangWu", Age: 25},
			},
		})
	} else {
		context.JSON(200, common.Result{
			Code:    200,
			Message: "success",
			Data: []User{
				{UserId: 4, Nickname: "QianLiu", Age: 30},
				{UserId: 5, Nickname: "ZhaoBa", Age: 30},
			},
		})
	}
}

func GetUser(ctx *gin.Context) {
	// http://xxx/api/user/123
	userId64, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)

	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    406,
			"message": "parameter missing",
		})
		return
	}

	userId := uint32(userId64)
	user := User{
		UserId:   userId,
		Nickname: "xxx",
		Age:      30,
	}

	ctx.JSON(200, common.OfSuccessResult(user))
}

func UpdateUser(ctx *gin.Context) {
	// 获取单个body参数
	//userId := ctx.PostForm("userId")
	//userName := ctx.PostForm("userName")

	// 绑定参数到结构体，可批量获取参数
	var user User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	if user.UserId <= 0 || user.Nickname == "" || user.Age <= 0 {
		ctx.JSON(200, gin.H{
			"code":    406,
			"message": "parameter missing",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})

}
