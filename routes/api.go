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
		v1.POST("/user", Controllers.UserInfo)

		// 后台用户的操作
		v1.GET("/admin/user/:id", Controllers.AdminUserInfo)
		v1.POST("/admin/user", Controllers.AdminUserCreate)
		v1.DELETE("/admin/user/:id", Controllers.AdminUserDestroy)
		v1.PATCH("/admin/user", Controllers.AdminUserUpdate)
		v1.GET("/admin/users", Controllers.AdminUserFindAll)


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