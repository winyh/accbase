package routes

import (
	"github.com/gin-gonic/gin"
	"accbase/app/Controllers"
)

func InitRouter()  *gin.Engine{
	r := gin.Default()

	r.Static("/public", "./public") // 静态文件服务
	r.LoadHTMLGlob("views/**/*") // 载入html模板目录

	// web路由
	r.GET("/", Controllers.Home)
	r.GET("/about", Controllers.About)
	r.GET("/post", Controllers.Post)
	r.GET("/user", Controllers.User)

	// api 路由
	InitApi(r)

	return r
}