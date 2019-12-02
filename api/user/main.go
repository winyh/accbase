package main

import (
	user "accbase/srv/user/proto"
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"log"
	"strings"
)

type User struct {
	Client user.UserService
}

type Response struct {
	Status bool
	Message string
	Token string
	Data string
}

func (u *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 User.Register API 请求")

	name, ok := req.Post["userName"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "UserName cannot be blank")
	}

	mobile, ok := req.Post["mobile"]
	if !ok || len(mobile.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Mobile cannot be blank")
	}

	email, ok := req.Post["email"]
	if !ok || len(email.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Email cannot be blank")
	}

	password, _p := req.Post["password"]
	if !_p || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Password cannot be blank")
	}

	confirm, _v := req.Post["confirm"]
	if !_v || len(confirm.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Confirm cannot be blank")
	}

	if strings.Join(confirm.Values, "") != strings.Join(password.Values, "")  {
		return errors.BadRequest("go.micro.api.user", "Password and Confirm cannot be different")
	}

	response, err := u.Client.Register(ctx, &user.RegisterRequest{
		UserName:  strings.Join(name.Values, ""),
		Mobile: strings.Join(mobile.Values, ""),
		Email: strings.Join(email.Values, ""),
		Password: strings.Join(password.Values, ""),
		Confirm: strings.Join(confirm.Values, ""),
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	data := Response{
		Status:response.Status,
		Message:response.Message,
		Token: response.Token,
	}

	b, _ := json.Marshal(data)

	rsp.Body = string(b)

	return nil
}

func (u *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Auth.Login API 请求")

	// name, ok := req.Get["userName"] // Get 请求参数获取方式
	name, ok := req.Post["userName"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "UserName cannot be blank")
	}

	password, ok := req.Post["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Password cannot be blank")
	}

	response, err := u.Client.Login(ctx, &user.LoginRequest{
		UserName:  strings.Join(name.Values, ""),
		Password: strings.Join(password.Values, ""),
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200

	data := Response{
		Status:response.Status,
		Message:response.Message,
		Token: response.Token,
	}

	b, _ := json.Marshal(data)
	rsp.Body = string(b)

	return nil
}

func (u *User) Account(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Auth.Account API 请求")
	token, ok := req.Get["token"]
	if !ok || len(token.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Token cannot be blank")
	}

	response, err := u.Client.Account(ctx, &user.AccountRequest{
		Token:  strings.Join(token.Values, ""),
	})

	if err != nil {
		return err
	}

	rsp.StatusCode = 200

	data := Response{
		Status:response.Status,
		Message:response.Message,
		Data: response.UserName,
	}

	b, _ := json.Marshal(data)
	rsp.Body = string(b)
	rsp.Body = string(b)

	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	service.Init()

	_ = service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserService("go.micro.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}