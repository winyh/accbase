package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Post(c *gin.Context) {
	c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
		"title": "文章页面",
	})
}