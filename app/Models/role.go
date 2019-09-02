package Models

import (
	. "accbase/database"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `json"roleName" binding:"required"` // 这里约定的是前台传值
	Description string `json"description" binding:"required"`
	Status int `json"status"`
}

func init()  {
	DB.AutoMigrate(&Role{})
}

func (role *Role) Insert() (err error){

	result := DB.Create(&role)
	if result.Error != nil {
		err = result.Error
	}
	return

}


func (role *Role) Destroy(id int) (err error){
	result := DB.Where(`id = ?`, id).Delete(&role)
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (role *Role) Update(id int) (err error){
	result := DB.Model(&role).Where("id = ?", id).Updates(&role)
	if result.Error != nil {
		err = result.Error
	}
	return
}


// FindOne 查询指定id admin用户
func (role *Role) FindOne(id int) (item Role, err error) {

	result := DB.First(&role, id)
	item = *role
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// FindAll 查询所有admin用户
func (role *Role) FindAll() (roles []Role, err error) {

	result := DB.Find(&role)

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}