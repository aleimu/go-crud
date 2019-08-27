package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud/model"
	. "go-crud/util"
	"net/http"
)

func GetStyle(c *gin.Context) {
	id := c.Query("id")
	style, err := model.GetStyleById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": style})
	}

}

func StyleList(c *gin.Context) {
	//var data map[string]interface{}
	data := make(map[string]interface{})
	id := c.Query("id")
	status := c.Query("status")
	system := c.Query("system")
	if id != "" {
		data["id"] = id
	}
	if status != "" {
		data["status"] = status
	}
	if system != "" {
		data["system"] = system
	}
	style, err := model.GetStyleList(data, 1, 10, " id desc")
	count := model.GetStyleTotal(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": style, "count": count})
	}

}

func AddStyle(c *gin.Context) {
	var styleform model.StyleForm
	//err := c.ShouldBind(styleform)
	err := c.ShouldBindQuery(styleform)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Sprintln("styleform", styleform)
	err = model.AddNewStyle(styleform)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "add success!"})

}

func SetStyle(c *gin.Context) {
	id := c.Query("id")
	style, err := model.GetStyleById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	var styleform model.StyleForm
	err = c.ShouldBind(styleform)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Sprintln("styleform", styleform)
	style.ImageName = styleform.ImageName
	//style.Status = styleform.Status
	model.DB.Save(&style)
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": style})

	//DB.Model(&group).Where("id = ?", id).Update("name", name)
}

func DelStyle(c *gin.Context) {
	id := c.Query("id")
	//status := c.DefaultQuery("status", "1")
	//strconv.Atoi(status)
	//DB.Delete(&email)
	model.DB.Where("id = ?", id).Delete(model.Style{})
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "OK"})
}
