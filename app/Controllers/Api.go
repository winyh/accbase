package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"accbase/app/Models"
)

type Admins struct {
	gorm.Model
	Name string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserCreate(c *gin.Context) {
	var json  Models.Admins
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	id, err := json.Insert()

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", id)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"id":id,
		"message":"创建成功",
	})
}

func UserDestroy(c *gin.Context)  {
	var json  Models.Admins
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}

	json.Destroy()

	c.JSON(200, gin.H{
		"status": true,
		"message":"删除成功",
	})
}

func UserUpdate(c *gin.Context) {
	var json  Models.Admins
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	id, err := json.Update(3)

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", id)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"id":id,
		"message":"更新成功",
	})
}

func UserFindAll(c *gin.Context)  {
	var json  Models.Admins
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	result, err := json.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": result,
		"message":"查询成功",
	})
}
