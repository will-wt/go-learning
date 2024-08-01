package dao

import (
	"go-learning/3-gorm/global"
	"gorm.io/gorm"
	"time"
)

type UserDO struct {
	Id          uint32    `json:"id"`
	GmtCreate   time.Time `json:"gmtCreate"`
	GmtModified time.Time `json:"gmtModified"`
	UserId      string    `json:"user_id"`
	Nickname    string    `json:"nickname"`
	Status      uint8     `json:"status"`
	TenantId    string    `json:"tenantId"`
}

// TableName 自定义表名
func (UserDO) TableName() string {
	return "sys_user"
}

type UserDAO struct {
}

// Get 查询用户
func (userDAO *UserDAO) Get(userId string) (*UserDO, error) {
	var item *UserDO
	// 查询所有字段，并写入到结构体中
	err := global.DB.Where("user_id = ?", userId).First(&item).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return item, err
}

// Add 增加用户
func (userDAO *UserDAO) Add(userDO UserDO) struct {
	id  uint32
	Row int64
	Err error
} {
	now := time.Now()
	userDO.GmtCreate = now
	userDO.GmtModified = now

	tx := global.DB.Create(&userDO)
	return struct {
		id  uint32
		Row int64
		Err error
	}{id: userDO.Id, Row: tx.RowsAffected, Err: tx.Error}
}

// UpdateInfo 修改用户
func (userDAO *UserDAO) UpdateInfo(userDO UserDO) struct {
	Row int64
	Err error
} {
	// 自定义sql
	sql := "update sys_user " +
		"set nickname = ?, " +
		"  	 gmt_modified = now() " +
		"where user_id = ? "
	tx := global.DB.Exec(sql, userDO.Nickname, userDO.UserId)

	//tx := global.DB.Model(&userDO).Where("user_Id = ?", userDO.UserId).Updates(UserDO{Nickname: userDO.Nickname, GmtModified: time.Now()})
	return struct {
		Row int64
		Err error
	}{Row: tx.RowsAffected, Err: tx.Error}
}

// UpdateStatus 修改用户状态
func (userDAO *UserDAO) UpdateStatus(userId string, status uint8) error {
	// 自定义sql
	sql := "update sys_user " +
		"set status = ?, " +
		"  	 gmt_modified = now() " +
		"where user_id = ? "
	tx := global.DB.Exec(sql, status, userId)
	return tx.Error
}

// Delete 删除用户
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
