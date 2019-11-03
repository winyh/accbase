package main

import (
	auth "accbase/srv/auth/proto"
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"log"
	"strings"
)

type Auth struct {
	Client auth.AuthService
}

func (a *Auth) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Auth.Register API request")

	name, ok := req.Post["userName"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "UserName cannot be blank")
	}

	password, _p := req.Post["password"]
	if !_p || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "Password cannot be blank")
	}

	verifyPassword, _v := req.Post["verifyPassword"]
	if !_v || len(verifyPassword.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "VerifyPassword cannot be blank")
	}

	if strings.Join(verifyPassword.Values, "") != strings.Join(password.Values, "")  {
		return errors.BadRequest("go.micro.api.auth", "Password and VerifyPassword cannot be different")
	}

	response, err := a.Client.Register(ctx, &auth.RegisterRequest{
		UserName:  "winyh",
		Password: "123456",
		VerifyPassword:"123456",
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message":"注册成功！",
		"token": response.Token,
	})
	rsp.Body = string(b)

	return nil
}

func (a *Auth) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Auth.Login API request")

	// name, ok := req.Get["userName"] // Get 请求参数获取方式
	name, ok := req.Post["userName"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "UserName cannot be blank")
	}

	password, ok := req.Post["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "Password cannot be blank")
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
		"message":"登录成功！",
		"token": response.Token,
	})
	rsp.Body = string(b)

	return nil
}

func (a *Auth) UserInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Auth.UserInfo API request")
	token, ok := req.Get["token"]
	if !ok || len(token.Values) == 0 {
		return errors.BadRequest("go.micro.api.auth", "Token cannot be blank")
	}

	response, err := a.Client.UserInfo(ctx, &auth.UserInfoRequest{
		Token:  strings.Join(token.Values, ""),
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message":"请求成功！",
		"userName": response.UserName,
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