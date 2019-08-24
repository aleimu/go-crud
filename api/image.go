package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)



func Upload(c *gin.Context) {

	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": OK, "msg": fmt.Sprintf("get form err: %s", err.Error())})
		return
	}
	//更换图片名防止重复
	filename := fmt.Sprintf("%s_%s_%s", strconv.FormatInt(time.Now().Unix(), 10), GetRandomString(10), filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, FilePath+filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": OK, "msg": fmt.Sprintf("upload file err: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "upload file success!", "image": filename})
	return

}

