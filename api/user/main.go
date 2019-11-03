package main

import (
	auth "accbase/srv/auth/proto"
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"log"
)

type Auth struct {
	Client auth.AuthService
}

func (a *Auth) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Auth.Login API request")


	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "Name cannot be blank")
	}

	response, err := a.Client.Login(ctx, &auth.LoginRequest{
		UserName:  "winyh",
		Password: "123456",
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Token,
	})
	rsp.Body = string(b)

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