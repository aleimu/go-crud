package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud/model"
	. "go-crud/util"
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

func GetImage(c *gin.Context) {
	var image = model.Image{}
	image.ID = model.Str2Int(c.Query("id"))
	result, err := model.GetImage(image);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "date": result})
}

// 比较了一下还是这种方式好用些
func GetImage2(c *gin.Context) {
	result, err := model.GetImage2(map[string]interface{}{"id": c.Query("id")});
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "date": result})
}

func AddImage(c *gin.Context) {
	var newImage model.ImageForm
	// err := c.ShouldBindQuery(&newImage) // 无效
	err := c.ShouldBind(&newImage) // form 必须这样绑定,json 也可以用这个方式校验,看源码可以看出是依据c.Request.Method, c.ContentType()推断出合适的类型
	// err := c.ShouldBindJSON(&newImage)	// 只有json可以
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newImage", newImage)
	err = model.AddNewImage(newImage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "add success!"})
	return
}

func DelImage(c *gin.Context) {
	id := c.Query("id") //字符串参数 welcome?firstname=Jane&lastname=Doe  c.Request.URL.Query().Get("lastname") 的一种快捷方式
	fmt.Println("Query id", id)
	id = c.Param("id")
	fmt.Println("Param id", id)
	id = c.PostForm("id")
	fmt.Println("PostForm id", id)
	err := model.DelImage(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "del success!"})
	return
}

func UpdateImage(c *gin.Context) {
	var updateImage model.ImageForm
	err := c.ShouldBind(&updateImage) // form 必须这样绑定,json 也可以用这个方式校验,看源码可以看出是依据c.Request.Method, c.ContentType()推断出合适的类型
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("newImage", updateImage)
	err = model.UpdateImage2(updateImage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "add success!"})
	return
}
