package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about/index.tmpl", gin.H{
		"title": "关于我们",
	})
}