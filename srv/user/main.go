package main

import (
	"accbase/srv/user/model"
	_ "accbase/srv/user/model"
	user "accbase/srv/user/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	"time"
)

type User struct{
	model.User
}

func (u *User) Register(ctx context.Context, req *user.RegisterRequest, rsp *user.RegisterResponse) error{
	log.Print("收到了 User.Register 用户注册请求！")
	rsp.Token = "218210190339039"
	rsp.Authority = "guest"
	return nil
}

func (u *User) Login(ctx context.Context, req *user.LoginRequest, rsp *user.LoginResponse) error {
	log.Print("收到了 User.Login 用户登录请求！")
	result, _ :=model.Find()
	fmt.Print(result.UserName)
	rsp.Token = "218210190339039"
	rsp.Authority = "admin"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	user.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}