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
	fmt.Println("收到请求！")
	userName := req.UserName
	password := req.Password
	if userName != "winyh" || password != "123456" {
		fmt.Println("账号或密码错误！")
	}

	tokenString := CreateToken(userName, password)
	rsp.Token = tokenString
	return nil
}

func (a *Auth)UserInfo(ctx context.Context, req *auth.UserInfoRequest, rsp *auth.UserInfoResponse) error{
	fmt.Println("收到请求！")
	token := req.Token
	if token != "" {
		fmt.Println("请检查token是否为空！")
	}
	b, name := VerifyToken(token)

	if b {
		rsp.UserName = name
	}

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

func VerifyToken(tokenString string) (bool, string) {
	var name string
	hmacSampleSecret := []byte("my_secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		//}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
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
