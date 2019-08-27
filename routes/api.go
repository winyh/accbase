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

		v1.POST("/upload", Controllers.FileUpload) // 文件上传

		v1.GET("/subjects", Controllers.GetAllSubjects) // 获取主题列表 [包含用户和角色]
		v1.GET("/objects", Controllers.GetAllObjects) // 获取对象列表 [资源 或操作 对象]
		v1.GET("/roles", Controllers.GetAllRoles) // 获取所有角色
		v1.GET("/policy", Controllers.GetPolicy) // 获取策略中的所有授权规则
		v1.POST("/addRoleForUser", Controllers.AddRoleForUser) // 为用户添加角色
		v1.POST("/addNamedPolicy", Controllers.AddNamedPolicy) //向当前命名策略添加授权规则

		v1.POST("/deleteUser", Controllers.DeleteUser) // 删除用户
		v1.POST("/deleteRole", Controllers.DeleteRole) // 删除角色
		v1.POST("/deletePermissionForUser", Controllers.DeletePermissionForUser) // 删除权限
	}

}