package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Avatar   string `gorm:"avatar"`
	Role     string `gorm:"role"`
	Status   int    `gorm:"status"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

func (u *User) Login(param models.UserParam) int64 {
	query := map[string]interface{}{
		"username": param.Username,
		"password": param.Password,
	}
	return global.Db.Where(query).First(&User{}).RowsAffected
}

func (u *User) Update(param models.UserParam) int64 {
	user := User{
		Id:     param.Id,
		Role:   param.Role,
		Status: param.Status,
	}
	return global.Db.Model(&user).Updates(user).RowsAffected
}

func (u *User) ChangePwd(param models.UserParam) int64 {
	return 1
}

func (u *User) Delete(id uint) int64 {
	return global.Db.Delete(&User{}, id).RowsAffected
}

func (u *User) GetList(page models.Page) (uList []models.UserList) {
	userList := make([]models.UserList, 0)
	common.RestPage(page, "user", &User{}, &userList, &User{})
	return userList
}
