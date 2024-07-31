package api

import (
	"github.com/gin-gonic/gin"
	"go-learning/3-gorm/common/result"
	"go-learning/3-gorm/dal"
	"go-learning/3-gorm/dal/dao"
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

// AddUser 新增用户
func AddUser(context *gin.Context) {
	var user User

	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, result.Error(400, err.Error()))
		return
	}

	userDO := dao.UserDO{
		UserId:   user.UserId,
		Nickname: user.Nickname,
		TenantId: user.TenantId,
	}

	dal.UserDAO.Add(userDO)
}

// GetUser 查询某个用户
func GetUser(ctx *gin.Context) {
	userId := ctx.Query("userId")

	userDO, _ := dal.UserDAO.Get(userId)

	ctx.JSON(200, result.OkWithData(userDO))
}

// UpdateUser 修改某个用户信息
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

	userDO := dao.UserDO{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}

	r := dal.UserDAO.Update(userDO)
	if r.Err != nil {
		ctx.JSON(200, result.Error(400, "更新失败"))
		return
	}

	ctx.JSON(200, result.Ok())

}
