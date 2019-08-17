package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/") // Your driver and data source.
	e, _ := casbin.NewEnforcer("./casbin/model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	if bool, _ := e.Enforce("alice", "data1", "read"); bool == true {
		fmt.Println("进入了")
		e.AddPolicy("eve", "data2", "read")
	} else {
		fmt.Println("拒绝了")
	}


	//if e.Enforce("alice", "data1", "read") == true {
	//	fmt.Println("进入了")
	//} else {
	//	fmt.Println("拒绝了")
	//}

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}