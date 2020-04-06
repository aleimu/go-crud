package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3) //设置计数器，数值即为goroutine的个数
	go func() {
		//Do some work
		time.Sleep(1 * time.Second)

		fmt.Println("Goroutine 1 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	go func() {
		//Do some work
		time.Sleep(2 * time.Second)

		fmt.Println("Goroutine 2 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()
	go other_go(&wg)
	wg.Wait() //主goroutine阻塞等待计数器变为0
	fmt.Printf("All Goroutine finished!")
}

func other_go(wg *sync.WaitGroup) { //WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址
	time.Sleep(3 * time.Second)
	go other_go1() //goroutine又继续派生新的goroutine，这种情况下使用WaitGroup就不太容易
	go other_go1()
	go other_go1()

	fmt.Println("Goroutine 3 finished!")
	wg.Done() //goroutine执行结束后将计数器减1

}
func other_go1() {
	time.Sleep(3 * time.Second)
	fmt.Println("Goroutine 4 finished!")
}
