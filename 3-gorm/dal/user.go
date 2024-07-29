package dal

import "go-learning/3-gorm/global"

type User struct {
	Id       int    `json:"id"`
	UserId   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "sys_user"
}

var item User
var items []User

func Get(userId string) User {
	global.DB.Where("user_id = ?", userId).Find(&item)
	return item
}

func Add(user User) struct {
	id  int
	row int64
	err error
} {
	tx := global.DB.Create(&user)
	return struct {
		id  int
		row int64
		err error
	}{id: user.Id, row: tx.RowsAffected, err: tx.Error}
}

func Update(user User) struct {
	row int64
	err error
} {
	tx := global.DB.Model(User{}).Where("user_Id = ?", user.UserId).Updates(user)
	return struct {
		row int64
		err error
	}{row: tx.RowsAffected, err: tx.Error}
}

func Delete(userId string) struct {
	row int64
	err error
} {
	tx := global.DB.Where("user_id = ?", userId).Delete(&User{})
	return struct {
		row int64
		err error
	}{row: tx.RowsAffected, err: tx.Error}
}
