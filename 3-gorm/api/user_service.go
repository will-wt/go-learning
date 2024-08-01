package api

import (
	"github.com/gin-gonic/gin"
	"go-learning/3-gorm/common/response"
	"go-learning/3-gorm/dal"
	"go-learning/3-gorm/dal/dao"
)

// UserDTO 定义用户请求入参结构体
type UserDTO struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname"`
	TenantId string `json:"tenantId"`
	Status   uint8  `json:"status"`
	Age      uint8  `json:"age"`
	Address  string `json:"address"`
}

// AddUser 新增用户
func AddUser(context *gin.Context) {
	var user UserDTO

	if err := context.Bind(&user); err != nil {
		response.Error(context, 400, err.Error())
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

	response.OkWithData(ctx, userDO)
}

// UpdateUserInfo 修改某个用户信息
func UpdateUserInfo(ctx *gin.Context) {
	// 绑定参数到结构体，可批量获取参数
	var user UserDTO
	if err := ctx.Bind(&user); err != nil {
		response.Error(ctx, 400, err.Error())
		return
	}

	if user.UserId == "" {
		response.Error(ctx, 406, "parameter missing")
		return
	}

	userDO := dao.UserDO{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}

	r := dal.UserDAO.UpdateInfo(userDO)
	if r.Err != nil {
		response.Error(ctx, 400, "更新失败")
		return
	}

	response.Ok(ctx)
}

// UpdateUserStatus 更新用户状态
func UpdateUserStatus(context *gin.Context) {
	var user UserDTO
	if err := context.Bind(&user); err != nil {
		response.Error(context, 400, err.Error())
		return
	}

	if user.UserId == "" {
		response.Error(context, 406, "parameter missing")
		return
	}

	err := dal.UserDAO.UpdateStatus(user.UserId, user.Status)
	if err != nil {
		response.Error(context, 400, "更新失败")
		return
	}

	response.Ok(context)
}
