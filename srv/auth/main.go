package main

import (
	auth "accbase/srv/auth/proto"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro"
	"log"
	"time"
)

type Auth struct{}

func (a *Auth)Register(ctx context.Context, req *auth.RegisterRequest, rsp *auth.RegisterResponse) error{
	userName := req.UserName //  proto 里的字段在这里首字母大写
	password := req.Password
	verifyPassword := req.VerifyPassword

	if userName != "" || password != "" {
		fmt.Println("账号或密码不能为空！")
	}

	if password != verifyPassword {
		fmt.Println("两次输入密码不一致！")
	}

	tokenString := CreateToken(userName, password)
	rsp.Token = tokenString
	return nil
}

func (a *Auth)Login(ctx context.Context, req *auth.LoginRequest, rsp *auth.LoginResponse) error{
	userName := req.UserName
	password := req.Password
	if userName != "winyh" || password != "123456" {
		fmt.Println("账号或密码错误！")
	}

	tokenString := CreateToken(userName, password)
	rsp.Token = tokenString
	return nil
}

func CreateToken(userName string, password string) string {

	hmacSampleSecret := []byte("my_secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userName,
		"password": password,
		"nbf": time.Date(2019, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
	return tokenString
}

func main(){
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
	)

	service.Init()

	_ = auth.RegisterAuthHandler(service.Server(), new(Auth))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
