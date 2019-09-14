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
	filter := make(map[string]interface{})
	filter["system"] = system
	filter["status"] = 1
	fmt.Println("filter:", filter)
	styles, err := model.GetStyleList(filter, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetSystemStyle err: " + err.Error())
	}
	return styles
}

// 查列表包含今天的就返回数据库 库中的列表，单单只查今天的就从redis中查询

func StorageDb() {
	// 遍历当前所有上线的广告,并统计缓存中的数据
	filter := make(map[string]interface{})
	filter["status"] = 1
	styles, err := model.GetStyleList(filter, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetSystemStyle err: " + err.Error())
	}
	for k, v := range styles {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
		fmt.Println("v.ID:", v.ID)
		InitTestRds(Int2Str(v.ID)) // 制造些数据测试
		StorageYesterdayDb(v.ID)
		InitTodayDb(v.ID)
		InitTodayRds(Int2Str(v.ID))
	}
}

func InitTodayDb(id int) {
	// 初始化今天的数据 mysql
	ctr := model.Ctr{StyleId: id, Show: 0, Click: 0, Crt: 0.0, ShowDay: DayShowValue, ClickDay: DayClickValue, CreateDate: Today()}
	err := model.AddCtr(ctr)
	if err != nil {
		panic("InitTodayDb err:" + err.Error())
	}
}

func InitTodayRds(id string) {
	// 初始化一天的每小时 redis
	cache.RedisClient.Del(DayShowKey + id) // 不删除不行啊,覆盖不了原来的list
	cache.RedisClient.Del(DayClickKey + id)
	_, err := cache.RedisClient.RPush(DayShowKey+id, DayValueStrList).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("InitTodayRds err: " + err.Error())
	}
	_, err = cache.RedisClient.RPush(DayClickKey+id, DayValueStrList).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("InitTodayRds err: " + err.Error())
	}
}

func InitTestRds(id string) {
	fmt.Println("--------------------------------------------------------------------------------------------------------")
	cache.RedisClient.Del(DayShowKey + id) // 不删除不行啊,覆盖不了原来的list
	cache.RedisClient.Del(DayClickKey + id)
	// 初始化一天的每小时 redis
	DayValueStrList := RandStringInt(24)
	_, err := cache.RedisClient.RPush(DayShowKey+id, DayValueStrList).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("InitTestRds err: " + err.Error())
	}
	_, err = cache.RedisClient.RPush(DayClickKey+id, DayValueStrList).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("InitTestRds err: " + err.Error())
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------")
}

func InitYesterdayDb(id int) model.Ctr {
	// 初始化今天的数据
	ctr := model.Ctr{StyleId: id, Show: 0, Click: 0, Crt: 0.0, ShowDay: DayShowValue, ClickDay: DayClickValue, CreateDate: Yesterday()}
	fmt.Println("ctr------:", ctr)
	//err := model.AddCtr(ctr)
	//if err != nil {
	//	panic("InitYesterdayDb err:" + err.Error())
	//}
	return ctr
}

// redis中都是字符串,转换类型到go中,很麻烦啊
func StorageYesterdayDb(id int) {
	// 归档昨天的数据
	filer := make(map[string]interface{})
	filer["style_id"] = id
	filer["create_date"] = Yesterday()
	result, err := model.GetCtr(filer)
	if err != nil {
		if err.Error() == "record not found" { // 昨天的记录缺失,需要补全
			result = InitYesterdayDb(id)
		} else {
			panic("StorageYesterdayDb err:" + err.Error())
		}
	}
	fmt.Println("old result:", result)
	str_id := Int2Str(id)
	//tmp_show := Str2Int(GetHourShow(str_id))
	//tmp_click := Str2Int(GetHourClick(str_id))

	tmp_click_day := GetDayClick(str_id)
	tmp_show_day := GetDayShow(str_id)
	tmp_click_sum := ListStrSum(tmp_click_day)
	tmp_show_sum := ListStrSum(tmp_show_day)
	fmt.Println("--------------------------------------")
	fmt.Println(tmp_click_sum, tmp_show_sum)
	fmt.Println("crt-------", float64(tmp_click_sum)/float64(tmp_show_sum))
	result.Click = tmp_click_sum
	result.Show = tmp_show_sum
	result.Crt = Float2(float64(tmp_click_sum) / float64(tmp_show_sum))
	result.ClickDay = ListStr2Str(tmp_click_day)
	result.ShowDay = ListStr2Str(tmp_show_day)
	fmt.Println("new result:", result)
	model.DB.Save(&result)

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
	systems := model.GetSystems3()
	for i, v := range systems {
		fmt.Println("i:", i, v)
		FreshRedis(v)
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": nil})
}

//刷新system的广告
func FreshRedis(system interface{}) {
	styles := GetSystemStyle(system)
	systemName := fmt.Sprintf("system:%s", system)
	AsyncRedis(systemName, Map2Str(styles))
}

func CtrCronJob() {
	// 遍历当前所有上线的广告,并统计缓存中的数据
	systems := model.GetSystems3()
	for i, v := range systems {
		fmt.Println("i:", i, v)
		FreshRedis(v)
	}
	filter := make(map[string]interface{})
	filter["status"] = 1
	styles, err := model.GetStyleList(filter, 0, 0, " id desc")
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
	fmt.Print("CleanHour:", id)
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

func RecoverNoKey(id string) {
	if err := recover(); err != nil {
		fmt.Println("RecoverNoKey:", id)
		fmt.Println(err)
		InitTodayRds(id)
	}
}

func SetHourShow(id, count string) {
	// 设置某广告的某小时的曝光量
	hourNow := int64(time.Now().Hour())
	defer RecoverNoKey(id)
	tmp, err := cache.RedisClient.LIndex(DayShowKey+id, hourNow).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("LIndex HourShow err: " + err.Error())
	}
	_, err = cache.RedisClient.LSet(DayShowKey+id, hourNow, StrSumStr(tmp, count)).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("SetHourShow err: " + err.Error())
	}
}

func SetHourClick(id, count string) {
	// 设置某广告的某小时的点击量
	hourNow := int64(time.Now().Hour())
	defer RecoverNoKey(id)
	tmp, err := cache.RedisClient.LIndex(DayClickKey+id, hourNow).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("LIndex HourClick err: " + err.Error())
	}
	_, err = cache.RedisClient.LSet(DayClickKey+id, hourNow, StrSumStr(tmp, count)).Result()
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
