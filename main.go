package main

import (
	"fmt"
	"go-crud/conf"
	"go-crud/router"
)

func main() {
	// 从配置文件读取配置
	test()
	conf.Init()
	// 启动定时任务
	router.StartCron()
	// 装载路由
	r := router.NewRouter()

	r.Run(":3001")
}

func test() {
	var a = 123;
	var b = 345;
	fmt.Println("--------------------------------------")
	fmt.Println("%0.2f:", (float32(a) / float32(b)))
	fmt.Println("--------------------------------------")

}
