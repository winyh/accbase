package main

import (
	_ "accbase/database"
	auth "accbase/srv/auth/proto"
	user "accbase/srv/user/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := auth.NewAuthService("go.micro.srv.auth", service.Client())

	// 新加
	cluser := user.NewUserService("go.micro.srv.user", service.Client())

	// Make request
	rsp, err := cl.Login(context.Background(), &auth.LoginRequest{
		UserName: "John",
	})

	// 新加
	rsp1, err1 := cluser.Login(context.Background(), &user.LoginRequest{
		UserName: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println(rsp.Token)
	fmt.Println(rsp1.Token)
}
