package router

import (
	"fmt"
	"github.com/cron"
	"time"
)

func StartCron() {
	c := cron.New()
	//err := c.AddFunc("*/50 * * * * *", func() { fmt.Println("Every 50 Seconds Run Once!", time.Now()) })
	//err := c.AddFunc("* 50 * * * *", func() { fmt.Println("Every 5 Minutes Run Once!", time.Now()) })
	err := c.AddFunc("0 0 * * * *", func() { fmt.Println("Run once an hour!", time.Now()) })
	if err != nil {
		panic(err.Error())
	}
	c.Start()
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
