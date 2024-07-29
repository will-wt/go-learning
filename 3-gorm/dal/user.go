package dal

import (
	"go-learning/3-gorm/global"
	"time"
)

type UserDO struct {
	Id          int       `json:"id"`
	GmtCreate   time.Time `json:"gmtCreate"`
	GmtModified time.Time `json:"gmtModified"`
	UserId      string    `json:"user_id"`
	Nickname    string    `json:"nickname"`
	Status      int       `json:"status"`
	TenantId    string    `json:"tenantId"`
}

// TableName 自定义表名
func (UserDO) TableName() string {
	return "sys_user"
}

var item UserDO
var items []UserDO

func Get(userId string) UserDO {
	global.DB.Where("user_id = ?", userId).Find(&item)
	return item
}

func Add(userDO UserDO) struct {
	id  int
	Row int64
	Err error
} {
	now := time.Now()
	userDO.GmtCreate = now
	userDO.GmtModified = now

	tx := global.DB.Create(&userDO)
	return struct {
		id  int
		Row int64
		Err error
	}{id: userDO.Id, Row: tx.RowsAffected, Err: tx.Error}
}

func Update(userDO UserDO) struct {
	Row int64
	Err error
} {
	tx := global.DB.Model(&userDO).Where("user_Id = ?", userDO.UserId).Updates(UserDO{Nickname: userDO.Nickname, GmtModified: time.Now()})
	return struct {
		Row int64
		Err error
	}{Row: tx.RowsAffected, Err: tx.Error}
}

func Delete(userId string) struct {
	Row int64
	Err error
} {
	tx := global.DB.Where("user_id = ?", userId).Delete(&UserDO{})
	return struct {
		Row int64
		Err error
	}{Row: tx.RowsAffected, Err: tx.Error}
}
