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

func (user *User) Create() (u User, err error) {
	result := DB.Create(&user)
	fmt.Print(result)
	return
}

func Find() (result User,err error) {
	var user User
	DB.Where("user_name = ?", "snowman").First(&user)
	result = user
	return
}
