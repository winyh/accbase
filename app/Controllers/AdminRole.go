package Controllers

import (
	"accbase/app/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Role struct {
	Models.Role
}


func RoleInfo(c *gin.Context)  {
	var json  Models.Role
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	role, err := json.FindOne(id)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": role,
		"message":"查询成功",
	})
}

func RoleCreate(c *gin.Context) {
	var json  Models.Role
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	res := json.Insert()

	if res != nil {
		fmt.Printf("database error %v", err)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"message":"创建成功",
	})
}

func RoleDestroy(c *gin.Context)  {
	var json  Models.Role
	var msg string
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := json.Destroy(id)

	if err != nil{
		msg = "删除失败"
	} else {
		msg = "删除成功"
	}

	c.JSON(200, gin.H{
		"status": true,
		"message": msg,
	})
}

func RoleUpdate(c *gin.Context) {
	var json  Models.Role
	var msg string
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	c.BindJSON(&json)

	err := json.Update(id)

	if err != nil{
		msg = "更新失败"
	} else {
		msg = "更新成功"
	}

	if err != nil {
		fmt.Printf("database error %v", err)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"message": msg,
	})
}

func RoleFindAll(c *gin.Context)  {
	var json  Models.Role

	admin, err := json.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": admin,
		"message":"查询成功",
	})
}