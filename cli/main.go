package main

import (
	"context"
	"fmt"

	auth "accbase/srv/auth/proto"
	"github.com/micro/go-micro"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := auth.NewAuthService("go.micro.srv.auth", service.Client())

	// Make request
	rsp, err := cl.Login(context.Background(), &auth.LoginRequest{
		UserName: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Token)
}
