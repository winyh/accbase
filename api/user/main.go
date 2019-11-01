package main

import (
	auth "accbase/srv/auth/proto"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"log"
)

type Auth struct {
	Client auth.AuthService
}

func (a *Auth) Hello(ctx context.Context, req *auth.LoginRequest, rsp *auth.LoginResponse) error {
	log.Print("Received Auth.Login API request")

	name := req.GetUserName()
	if len(name) == 0 {
		return errors.BadRequest("go.micro.api.auth", "Name cannot be blank")
	}

	response, err := a.Client.Login(ctx, &auth.LoginRequest{
		UserName: "winyh",
		Password: "123456",
	})

	log.Print(response)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.auth"),
	)

	service.Init()

	_ = service.Server().Handle(
		service.Server().NewHandler(
			&Auth{Client: auth.NewAuthService("go.micro.srv.auth", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}