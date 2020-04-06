package main

import (
	"fmt"
	"log"
	"time"
)

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}
func DelayFunction(ticker *time.Ticker) {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
		ticker.Stop()
	}
}

// TickerDemo 用于演示ticker基础用法
func TickerDemo(ticker *time.Ticker) {
	for range ticker.C {
		log.Println("Ticker tick.")
	}
}
func main() {
	ch := make(chan string)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	fmt.Println("start!")
	go TickerDemo(ticker)
	go WaitChannel(ch)
	DelayFunction(ticker)
	//time.Sleep(2 * time.Second)
	fmt.Println("end!")
	//ch <- "a"

}

// go build -ldflags "-s -w"  -o .\go_build_timeout_demo_go.exe timeout_demo.go 打包，并减小二级制文件大小