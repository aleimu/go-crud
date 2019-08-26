package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-crud/model"
	"net/http"
	"strconv"
)

func GetGroup(c *gin.Context) {
	id := c.Query("id")
	group, err := FindGroupById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": group})
	}

}
func AddGroup(c *gin.Context) {
	name := c.PostForm("name")
	DB.Create(&Group{Name: name, Status: 1})
	c.JSON(http.StatusOK, gin.H{"code": OK, "mgs": "OK"})
}

func SetGroup(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	status := c.DefaultPostForm("status", "1")
	fmt.Sprintln("id:%s, name:%s, status:%s", id, name, status)
	statusTemp, err := strconv.Atoi(status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "err": err})
	}
	group, err := FindGroupById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	group.Name = name
	group.Status = statusTemp
	DB.Save(&group)
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": group})

	//DB.Model(&group).Where("id = ?", id).Update("name", name)
}

func DelGroup(c *gin.Context) {
	id := c.Query("id")
	//status := c.DefaultQuery("status", "1")
	//strconv.Atoi(status)
	//DB.Delete(&email)
	DB.Where("id = ?", id).Delete(Group{})
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "OK"})
}
