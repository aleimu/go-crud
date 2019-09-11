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

func GetSystemStyle(system interface{}) []model.Style {
	data := make(map[string]interface{})
	data["system"] = system
	data["status"] = 1
	fmt.Println("data:", data)
	styles, err := model.GetStyleList(data, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetSystemStyle err: " + err.Error())
	}
	return styles
}

func AsyncRedis(systemName string, data interface{}) {
	// 按分组group同步修改缓存中的信息
	_, err := cache.RedisClient.Set(systemName, data, OneDay).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AsyncRedis err: " + err.Error())
	}
}

func FreshAllRedis(c *gin.Context) {
	// 分别刷新每个分组的广告到缓存中
	groups := model.FindGroups()
	for i, v := range groups {
		fmt.Println("i:", i, v)
		FreshRedis(v.ID)
		InitDay(Int2Str(v.ID))
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": nil})
}

func FreshRedis(system interface{}) {
	styles := GetSystemStyle(system)
	systemName := fmt.Sprintf("system:%s", system)
	AsyncRedis(systemName, Map2Str(styles))
}

func CtrCronJob() {
	// 遍历当前所有上线的广告,并统计缓存中的数据
	data := make(map[string]interface{})
	data["status"] = 1
	styles, err := model.GetStyleList(data, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetSystemStyle err: " + err.Error())
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
