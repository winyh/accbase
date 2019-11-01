package main

import (
	"github.com/micro/go-micro"
)

func main(){
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	service.Init()

	auth.RegisterAuthHandler(service.Server(), &Auth{
		customers: loadCustomerData("data/customers.json"),
	})

	service.Run()
}
