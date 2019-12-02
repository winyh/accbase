package main

import (
	_ "accbase/database"
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

	cl := user.NewUserService("go.micro.srv.user", service.Client())

	// Make request
	rsp, err := cl.Login(context.Background(), &user.LoginRequest{
		UserName: "John",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Token)
}
