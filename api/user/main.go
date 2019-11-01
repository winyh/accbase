package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"log"
)

type Say struct{}

var (
	cl hello.SayService
)


// 用户注册分发token
func register(){

}

// 用户登录分发token
func login(){

}

// 用户退出登录
func logout(){

}

// 刷新令牌
func refreshToken(){

}

// 	获取用户信息,根据token获取
func getUserInfo(){

}

// 	获取用户角色集合
func getUserRoles(){

}


func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response, err := cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main()  {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.Init()

	// setup Greeter Server Client
	cl = hello.NewSayService("go.micro.srv.user", client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/api/user", say.Anything)
	router.GET("/api/user/:id", say.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
