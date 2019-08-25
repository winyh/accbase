package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
		"title": "后台首页",
	})
}

func FileUpload(c *gin.Context){

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	filename := filepath.Base(file.Filename)
	dst := "./public/upload/" + filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"message": "上传成功",
		"data":dst,
	})
}
