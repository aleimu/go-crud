package main

import (
	"go-crud/conf"
	"go-crud/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	// 启动定时任务
	router.StartCron()
	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
