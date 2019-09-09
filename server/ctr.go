package server

import (
	"fmt"
	"go-crud/cache"
	"go-crud/model"
	. "go-crud/util"
)

func StorageDb() {
	// 入库
	RangeCtr()
}

// 查列表包含今天的就返回数据库 库中的列表，单单只查今天的就从redis中查询

func RangeCtr() {
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
