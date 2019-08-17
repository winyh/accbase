package main

import (
	"fmt"
	"github.com/casbin/casbin"
	//"github.com/gin-gonic/gin"
)

func main(){
	e := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.
	// 这三条数据可以从数据库动态获取到

	if e.Enforce(sub, obj, act) == true {
		fmt.Println("进入了")
	} else {
		fmt.Println("拒绝了")
	}

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}