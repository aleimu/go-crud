package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-crud/model"
	"net/http"
	"strconv"
)

func GetStyle(c *gin.Context) {
	id := c.Query("id")
	group, err := FindStyleById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": group})
	}

}
func AddStyle(c *gin.Context) {
	name := c.PostForm("name")
	DB.Create(&Style{ImageName: name, Status: 1})
	c.JSON(http.StatusOK, gin.H{"code": OK, "mgs": "OK"})
}

func SetStyle(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	status := c.DefaultPostForm("status", "1")
	fmt.Sprintln("id:%s, name:%s, status:%s", id, name, status)
	statusTemp, err := strconv.Atoi(status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "err": err})
	}
	group, err := FindStyleById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	group.ImageName = name
	group.Status = statusTemp
	DB.Save(&group)
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": group})

	//DB.Model(&group).Where("id = ?", id).Update("name", name)
}

func DelStyle(c *gin.Context) {
	id := c.Query("id")
	//status := c.DefaultQuery("status", "1")
	//strconv.Atoi(status)
	//DB.Delete(&email)
	DB.Where("id = ?", id).Delete(Style{})
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "OK"})
}
