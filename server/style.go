package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud/cache"
	"go-crud/model"
	. "go-crud/util"
	"net/http"
	"time"
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
	var sql model.StyleForm
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
	go FreshRedis(sql.GroupId) // 同步下

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
	if result.RowsAffected != ok || result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "data": "update style err!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": sql})
	go FreshRedis(sql.GroupId) // 同步下
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
	// 依据分组查询此分组下的style并加入缓存中,二次查询尝试从缓存中获取数据
	id := c.Query("system")
	groupName := "system:" + id
	rds, err := cache.RedisClient.Get(groupName).Result()
	if err != nil || rds == "" {
		rds1 := GetFreshStyle(id)
		c.JSON(http.StatusOK, gin.H{"code": OK, "source": "db", "data": rds1})
		rds = Map2Str(rds1)
		go AsyncRedis(groupName, rds)
		// FIXME Str2Map不能解析[{map}],应该使用Str2Slice
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": Str2Slice(rds)})

}

func GetFreshStyle(group_id interface{}) []model.Style {
	data := make(map[string]interface{})
	data["group_id"] = group_id
	fmt.Println("data:", data)
	styles, err := model.GetStyleList(data, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetFreshStyle err: " + err.Error())
	}
	return styles
}

func AsyncRedis(groupName string, data interface{}) {
	// 按分组group同步修改缓存中的信息
	_, err := cache.RedisClient.Set(groupName, data, OneDay).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AsyncRedis err: " + err.Error())
	}
}

func FreshRedis(group_id interface{}) {
	styles := GetFreshStyle(group_id)
	groupName := fmt.Sprintf("system:%s", group_id)
	AsyncRedis(groupName, Map2Str(styles))
}

func FreshAllRedis(c *gin.Context) {
	// 分别刷新每个分组的广告到缓存中
	groups := model.FindGroups()
	for i, v := range groups {
		fmt.Println("i:", i, v)
		FreshRedis(v.ID)
		InitDay(Int2Str(v.ID))
	}
	CtrCronJob()
	c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": nil})
}

func CtrCronJob() {
	// 遍历当前所有上线的广告,并统计缓存中的数据
	data := make(map[string]interface{})
	data["status"] = 1
	styles, err := model.GetStyleList(data, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetFreshStyle err: " + err.Error())
	}
	for k, v := range styles {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
		StackDay(Int2Str(v.ID))
	}
}

func InitDay(id string) {
	// 初始化一天的每小时
	_, err := cache.RedisClient.RPush(DayShowKey+id, DayValueStrList).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("DayShowKey err: " + err.Error())
	}
	_, err = cache.RedisClient.RPush(DayClickKey+id, DayValueStrList).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("DayClickKey err: " + err.Error())
	}
	StackDay(id)
}

func StackDay(id string) {
	// 将hour:show:key:X 的值依据当前小时移入day:click:key:X
	scount := GetHourShow(id)
	ccount := GetHourClick(id)
	fmt.Println("scount:", scount, "ccount:", ccount)
	SetHourShow(id, scount)
	SetHourClick(id, ccount)
	CleanHour(id)
}

func CleanHour(id string) {
	// 清理小时计数
	_, err := cache.RedisClient.Set(HourShowKey+id, 0, OneDay).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AddHourClick err: " + err.Error())
	}
	_, err = cache.RedisClient.Set(HourClickKey+id, 0, OneDay).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AddHourClick err: " + err.Error())
	}
}

func SetHourShow(id, count string) {
	// 设置某广告的某小时的曝光量
	hourNow := int64(time.Now().Hour())
	_, err := cache.RedisClient.LSet(DayShowKey+id, hourNow, count).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("SetHourShow err: " + err.Error())
	}
}

func SetHourClick(id, count string) {
	// 设置某广告的某小时的点击量
	hourNow := int64(time.Now().Hour())
	_, err := cache.RedisClient.RPush(DayClickKey+id, hourNow, count).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("SetHourClick err: " + err.Error())
	}
}

func AddHourShow(c *gin.Context) {
	// 增加曝光量
	count := c.DefaultQuery("count", "1")
	id := c.DefaultQuery("id", "1")
	_, err := cache.RedisClient.IncrBy(HourShowKey+id, Str2Int64(count)).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AddHourShow err: " + err.Error())
	}
	//c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": nil}) // 不返回也没关系,看情况吧
}

func AddHourClick(c *gin.Context) {
	// 增加点击量
	count := c.DefaultQuery("count", "1")
	id := c.DefaultQuery("id", "1")
	_, err := cache.RedisClient.IncrBy(HourClickKey+id, Str2Int64(count)).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AddHourClick err: " + err.Error())
	}
	//c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": nil})
}

func GetHourShow(id string) string {
	// 获取曝光量
	count, err := cache.RedisClient.Get(HourShowKey + id).Result()
	fmt.Println("GetHourShow:", count, err)
	if err == cache.RedisNil {
		fmt.Println(HourShowKey + id + " does not exist")
		return "0"
	} else if err != nil {
		fmt.Println("GetHourShow err: " + err.Error())
		panic("GetHourShow err: " + err.Error())
	} else {
		fmt.Println("count:", count)
	}
	return count
}

func GetHourClick(id string) string {
	// 获取点击量
	count, err := cache.RedisClient.Get(HourClickKey + id).Result()
	fmt.Println("GetHourClick:", count)
	if err == cache.RedisNil {
		fmt.Println(HourClickKey + id + " does not exist")
		return "0"
	} else if err != nil {
		fmt.Println("GetHourClick err: " + err.Error())
		panic("GetHourClick err: " + err.Error())
	} else {
		fmt.Println("count:", count)
	}
	return count
}

func GetDayShow(id string) []string {
	// 获取曝光量
	count, err := cache.RedisClient.LRange(DayShowKey+id, 0, 23).Result()
	fmt.Println("GetDayShow:", count)
	if err == cache.RedisNil {
		fmt.Println(DayShowKey + id + " does not exist")
		return DayValueStrList
	} else if err != nil {
		fmt.Println("GetDayShow err: " + err.Error())
		panic("GetDayShow err: " + err.Error())
	} else {
		fmt.Println("count:", count)
	}
	return count
}

func GetDayClick(id string) []string {
	// 获取点击量
	count, err := cache.RedisClient.LRange(DayClickKey+id, 0, 23).Result()
	fmt.Println("GetDayClick:", count)
	if err == cache.RedisNil {
		fmt.Println(DayClickKey + id + " does not exist")
		return DayValueStrList
	} else if err != nil {
		fmt.Println("GetDayClick err: " + err.Error())
		panic("GetDayClick err: " + err.Error())
	} else {
		fmt.Println("count:", count)
	}
	return count
}
