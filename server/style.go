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
	fmt.Println("Query:", id, status, system)
	if id != "" {
		data["id"] = id
	}
	if status != "" {
		data["status"] = status
	}
	if system != "" {
		data["system"] = system
	}
	fmt.Println("data:", data)
	style, err := model.GetStyleList(data, 0, 10, " id desc")
	count := model.GetStyleTotal(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": style, "count": count})
	}

}

func AddStyle(c *gin.Context) {
	var styleform model.StyleForm
	err := c.ShouldBind(&styleform) // form 必须这样绑定,json 也可以用这个方式校验,看源码可以看出是依据c.Request.Method, c.ContentType()推断出合适的类型
	// err := c.ShouldBindJSON(&newImage)	// 只有json可以
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln("styleform", styleform)
	err = model.AddNewStyle(styleform)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "add success!"})

}

func UpdateStyle(c *gin.Context) {
	var sql model.StyleForm
	var ok int64 = 1
	err := c.ShouldBind(sql)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Sprintln("sql", sql)
	result := model.DB.Model(&model.Style{}).Where("id = ?", sql.Id).Updates(&model.Style{ImageName: sql.ImageName, Url: sql.Url, GroupId: sql.GroupId, Status: sql.Status}) // model式批量更新
	fmt.Println("result:", result, result.Error, result.RowsAffected)
	if result.RowsAffected == ok || result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": sql})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": sql})
}

func DelStyle(c *gin.Context) {
	id := c.Query("id")
	//status := c.DefaultQuery("status", "1")
	//strconv.Atoi(status)
	//DB.Delete(&email)
	model.DB.Where("id = ?", id).Delete(model.Style{})
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "OK"})
}
