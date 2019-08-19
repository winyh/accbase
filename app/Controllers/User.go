package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
		"title": "用户页面",
	})
}