package server

import (
	"fmt"
	"go-crud/model"
	"time"
)

func StorageDb() {
	// 入库

}

// 查列表包含今天的就返回数据库 库中的列表，单单只查今天的就从redis中查询

func InitTodayDb() {
	// 初始化今天的数据
	ctr := model.Ctr{StyleId: 1, Show: 0, Click: 0, Crt: 0.0, ShowDay: DayShowVlue, ClickDay: DayClickVlue, CreateDate: time.Now()}
	err := model.AddCtr(ctr)
	if err != nil {
		panic("InitTodayDb err:" + err.Error())
	}
}

func StorageYesterdayDb() {
	// 归档昨天的数据
	filer:=make(map[string]interface{})
	filer["style_id"]=1
	filer["create_date"]=Yesterday()
	model.GetCtr(filer)

}

func Today() time.Time {
	// 获取今天零点
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)

	fmt.Println(t.Unix() + 1)
	fmt.Println(t2.AddDate(0, 0, 1).Unix())
	return t
}

func Yesterday() time.Time {
	// 获取昨天零点
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	//t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	t.AddDate(0, 0, -1)
	return t
}
