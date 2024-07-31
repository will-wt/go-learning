package dao

import (
	"go-learning/3-gorm/global"
	"gorm.io/gorm"
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

type UserDAO struct {
}

// TableName 自定义表名
func (UserDO) TableName() string {
	return "sys_user"
}

func (userDAO *UserDAO) Get(userId string) (*UserDO, error) {
	var item *UserDO
	err := global.DB.Where("user_id = ?", userId).First(&item).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return item, err
}

func (userDAO *UserDAO) Add(userDO UserDO) struct {
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

func (userDAO *UserDAO) Update(userDO UserDO) struct {
	Row int64
	Err error
} {
	tx := global.DB.Model(&userDO).Where("user_Id = ?", userDO.UserId).Updates(UserDO{Nickname: userDO.Nickname, GmtModified: time.Now()})
	return struct {
		Row int64
		Err error
	}{Row: tx.RowsAffected, Err: tx.Error}
}

func (userDAO *UserDAO) Delete(userId string) struct {
	Row int64
	Err error
} {
	tx := global.DB.Where("user_id = ?", userId).Delete(&UserDO{})
	return struct {
		Row int64
		Err error
	}{Row: tx.RowsAffected, Err: tx.Error}
}
