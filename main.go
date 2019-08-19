package main

import (
	"accbase/config"
	"fmt"
	// "github.com/casbin/casbin"
	_ "accbase/config"
	"accbase/routes"
)

func main(){

	fmt.Println(config.PORT)

	//e := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	//
	//sub := "alice" // the user that wants to access a resource.
	//obj := "data1" // the resource that is going to be accessed.
	//act := "write" // the operation that the user performs on the resource.
	//
	//if e.Enforce(sub, obj, act) == true {
	//	fmt.Println("进入了")
	//} else {
	//	fmt.Println("拒绝了")
	//}

	r := routes.InitRouter()

	r.Run(":" + config.PORT) // 监听并在 0.0.0.0:8050 上启动服务
}