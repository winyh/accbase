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
		v1.POST("/user/create", Controllers.UserCreate)
		v1.POST("/user/delete", Controllers.UserDestroy)
		v1.POST("/user/update", Controllers.UserUpdate)
		v1.GET("/users", Controllers.UserFindAll)
	}

}