package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
		"title": "后台首页",
	})
}
