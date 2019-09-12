package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud/cache"
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
	var sql model.Style
	err := c.ShouldBind(&sql) // form 必须这样绑定,json 也可以用这个方式校验,看源码可以看出是依据c.Request.Method, c.ContentType()推断出合适的类型
	// err := c.ShouldBindJSON(&newImage)	// 只有json可以
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln("styleform", sql)
	err = model.AddNewStyle(sql)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "add success!"})
	go FreshRedis(sql.ID) // 同步下

}

func UpdateStyle(c *gin.Context) {
	var sql model.Style
	var ok int64 = 1
	err := c.ShouldBind(sql)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Sprintln("sql", sql)
	result := model.DB.Model(&model.Style{}).Where("id = ?", sql.ID).Updates(&model.Style{ImageName: sql.ImageName, Url: sql.Url, GroupId: sql.GroupId, Status: sql.Status}) // model式批量更新
	fmt.Println("result:", result, result.Error, result.RowsAffected)
	if result.RowsAffected != ok || result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "data": "update style err!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": sql})
	go FreshRedis(sql.ID) // 同步下
}

func DelStyle(c *gin.Context) {
	id := c.Query("id")
	//status := c.DefaultQuery("status", "1")
	//strconv.Atoi(status)
	//DB.Delete(&email)
	model.DB.Where("id = ?", id).Delete(model.Style{})
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "OK"})
}

func SearchStyle(c *gin.Context) {
	// 依据system查询style并加入缓存中,二次查询尝试从缓存中获取数据
	system := c.Query("system")
	systemName := "system:" + system
	rds, err := cache.RedisClient.Get(systemName).Result()
	if err != nil || rds == "" {
		rds1 := GetSystemStyle(system)
		c.JSON(http.StatusOK, gin.H{"code": OK, "source": "db", "data": rds1})
		rds = Map2Str(rds1)
		go AsyncRedis(systemName, rds)
		// FIXME Str2Map不能解析[{map}],应该使用Str2Slice
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": Str2Slice(rds)})

}

func GetSystems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": model.GetSystems3()})
}

/*
Go语言中byte和rune实质上就是uint8和int32类型。
Go的字符称为rune，等价于C中的char，可直接与整数转换
rune实际是整型，必需先将其转换为string才能打印出来，否则打印出来的是一个整数
*/
