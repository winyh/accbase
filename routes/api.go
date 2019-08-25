package routes

import (
	"github.com/gin-gonic/gin"
	"accbase/app/Controllers"
)

func InitApi(r *gin.Engine){

	// 简单的api路由组: v1
	v1 := r.Group("/api")
	{
		v1.GET("/ping", Controllers.Ping)
		v1.POST("/token", Controllers.GetToken)
		v1.POST("/user/info", Controllers.UserInfo)
		v1.POST("/user/create", Controllers.UserCreate)
		v1.POST("/user/delete", Controllers.UserDestroy)
		v1.POST("/user/update", Controllers.UserUpdate)
		v1.GET("/users", Controllers.UserFindAll)

		v1.GET("/roles", Controllers.GetAllRoles) // 获取所有角色
		v1.POST("/addRoleForUser", Controllers.AddRoleForUser) // 为用户添加角色
		v1.POST("/addNamedPolicy", Controllers.AddNamedPolicy) //向当前命名策略添加授权规则
	}

}