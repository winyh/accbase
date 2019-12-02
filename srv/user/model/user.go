package model

import (
	. "accbase/database"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"userName"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Email string `json:"email"`
}

// 新建
func Create(userName string,  mobile string, email string, password string) (user User, err error) {
	DB.Create(&User{UserName:userName, Password:password, Mobile: mobile, Email:email})
	fmt.Print(user)
	return
}

// 根据id删除
func Destroy(id int64)(user User, err error){
	DB.Where("id ?", id).Delete(&user)
	return
}

// 根据id更新
func Update(id int64, name string, mobile string, email string)(user User){
	DB.First(&user, id)
	DB.Model(&user).Updates(User{UserName: name, Mobile: mobile, Email: email})
	return
}

// 查询所有
func FindAll() (user []User,err error) {
	DB.Find(&user)
	return
}

// 根据名称查询
func FindOneByName(name string)(user User, err error)  {
	DB.Where("user_name = ?", name).First(&user)
	return
}
