package main

import (
	"accbase/srv/user/model"
	_ "accbase/srv/user/model"
	user "accbase/srv/user/proto"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro"
	"log"
	"time"
)

type User struct{
	UserName string `json:"userName"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Confirm string `json:"confirm"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Email string `json:"email"`
}

func (u *User) Register(ctx context.Context, req *user.RegisterRequest, rsp *user.RegisterResponse) error{
	log.Print("收到了 User.Register 用户注册请求！")
	userName := req.UserName
	mobile := req.Mobile
	email := req.Email
	password := req.Password

	result, _ := model.FindOneByName(userName)

	if userName == result.UserName {
		rsp.Message = "账号重复！"
		rsp.Status = false
		return nil
	}
	fmt.Print(result)
	if mobile == result.Mobile {
		rsp.Message = "手机号重复！"
		rsp.Status = false
		return nil
	}

	user, err := model.Create(userName, mobile, email, password)

	if err != nil {
		return err
	}

	fmt.Print(user)

	token := CreateToken(userName, password)
	rsp.Token = token
	rsp.Status = true
	rsp.Message = "注册成功！"
	rsp.Authority = "guest"
	return nil
}

func (u *User) Login(ctx context.Context, req *user.LoginRequest, rsp *user.LoginResponse) error {
	log.Print("收到了 User.Login 用户登录请求！")
	userName := req.UserName
	password := req.Password

	result, _ :=model.FindOneByName(userName)

	fmt.Print(result)

	if userName != result.UserName || password != result.Password {
		fmt.Println("账号或密码错误！")
		rsp.Message = "账号或密码错误！"
		rsp.Status = false
		return nil
	}

	token := CreateToken(userName, password)
	rsp.Status = true
	rsp.Token = token
	rsp.Message = "登录成功！"
	rsp.Authority = "admin"
	return nil
}

func (u *User) Account(ctx context.Context, req *user.AccountRequest, rsp *user.AccountResponse) error {
	log.Print("收到了 User.Account 根据token查用户信息请求！")
	token := req.Token
	if token == "" {
		fmt.Println("请检查token是否为空！")
		return nil
	}
	b, name := VerifyToken(token)
	fmt.Print(name)
	if b {
		rsp.UserName = name
	}

	rsp.Status = true
	rsp.Message = "查询成功！"
	rsp.Authority = "admin"

	return nil
}

// 生成 token
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

// 解析 token
func VerifyToken(tokenString string) (bool, string) {
	var name string
	hmacSampleSecret := []byte("my_secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil{
		log.Print("请检查 token 参数！")
		return false, name
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["nbf"])
		name = claims["username"].(string)
	} else {
		fmt.Println(err)
		return false, name
	}

	return true, name
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