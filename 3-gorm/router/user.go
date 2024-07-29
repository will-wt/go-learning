package router

import (
	"github.com/gin-gonic/gin"
	"go-learning/3-gorm/common/result"
	"go-learning/3-gorm/dal"
	"net/http"
)

// User 定义用户结构体
type User struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname"`
	TenantId string `json:"tenantId"`
	Age      uint8  `json:"age"`
	Address  string `json:"address"`
}

func AddUser(context *gin.Context) {
	var user User

	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, result.Error(400, err.Error()))
		return
	}

	userDO := dal.UserDO{
		UserId:   user.UserId,
		Nickname: user.Nickname,
		TenantId: user.TenantId,
	}

	dal.Add(userDO)
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Query("userId")

	userDO := dal.Get(userId)

	ctx.JSON(200, result.OkWithData(userDO))
}

func UpdateUser(ctx *gin.Context) {
	// 绑定参数到结构体，可批量获取参数
	var user User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, result.Error(400, err.Error()))
		return
	}

	if user.UserId == "" {
		ctx.JSON(200, result.Error(406, "parameter missing"))
		return
	}

	userDO := dal.UserDO{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}

	r := dal.Update(userDO)
	if r.Err != nil {
		ctx.JSON(200, result.Error(400, "更新失败"))
		return
	}

	ctx.JSON(200, result.Ok())

}
