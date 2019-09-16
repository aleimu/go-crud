package router

import (
	"fmt"
	"github.com/robfig/cron"
	"go-crud/server"
	"time"
)

func StartCron() {
	c := cron.New()
	//err := c.AddFunc("*/50 * * * * *", func() { fmt.Println("Every 50 Seconds Run Once!", time.Now()) })
	//err := c.AddFunc("* 50 * * * *", func() { fmt.Println("Every 5 Minutes Run Once!", time.Now()) })
	//err := c.AddFunc("0 0 * * * *", CtrHourJob)
	err := c.AddFunc("15 * * * * *", CtrHourJob)
	if err != nil {
		panic(err.Error())
	}
	c.Start()
}

func CtrHourJob() {
	fmt.Println("cron job start!")
	// 每小时执行一次的点击量/曝光量统计任务
	now := time.Now().Hour()
	if now < 1 { // 当天零点,结算上一天的数据,初始化今天的数据结构
		server.StorageDb()
	}
	server.CtrCronJob()
	fmt.Println("cron job end!")
}

/* 也可以参考cron的自测用例学习如何使用
Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
字段名	是否必须	允许的值	允许的特定字符
秒(Seconds)	是	0-59	* / , -
分(Minutes)	是	0-59	* / , -
时(Hours)	是	0-23	* / , -
日(Day of month)	是	1-31	* / , – ?
月(Month)	是	1-12 or JAN-DEC	* / , -
星期(Day of week)	否	0-6 or SUM-SAT	* / , – ?

*/
