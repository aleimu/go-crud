package main

import "fmt"

//const (
//	mutexLocked = 1 << iota // mutex is locked
//	mutexWoken
//	mutexStarving
//	mutexWaiterShift      = iota
//	starvationThresholdNs = 1e6
//)
//在Go中使用另一个常量iota计数器,只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0(const内部的第一行之前)，
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。使用iota能简化定义，在定义枚举时很有用。

const (
	mutexLocked           = 1 << iota // mutex is locked 0
	test0                 = 123
	mutexWoken            = 1 << iota << iota //2,2 iota实际上是遍历const块的索引，每行中即便多次使用iota，其值也不会递增。
	mutexStarving         = 1 << iota         //3
	test1                 = 123
	test2                 = 324
	mutexWaiterShift      = iota //6
	starvationThresholdNs = 1e6
)

func main() {
	//chan1 := make(chan int)
	//chan2 := make(chan int)
	//
	//go func() {
	//	close(chan1)	//关闭的chan也可以读取
	//}()
	//
	//go func() {
	//	close(chan2)
	//}()
	//
	//select {
	//case <-chan1:
	//	fmt.Println("chan1 ready.")
	//case <-chan2:
	//	fmt.Println("chan2 ready.")
	//}
	//
	//fmt.Println("main exit.")
	//test_func1()
	//UpdateTable()
	//UpdateTable_true()
	test_list()

}

func test_func1() {
	fmt.Println(test0, test1, test2)
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	fmt.Println(starvationThresholdNs)

	fmt.Println(1 << 0)
	fmt.Println(1 << 1 << 3)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
}

/* UpdateTable
导致recover()失效（永远返回nil）。
以下三个条件会让recover()返回nil:
1. panic时指定的参数为nil；（一般panic语句如panic("xxx failed...")）
2. 当前协程没有发生panic；
3. recover没有被defer方法直接调用；

recover() 调用栈为“defer （匿名）函数” --> IsPanic() --> recover()。
也就是说，recover并没有被defer方法直接调用。符合第3个条件，所以recover() 永远返回nil。
*/

func UpdateTable() {
	// defer中决定提交还是回滚
	defer func() {
		fmt.Println("UpdateTable defer start!")
		if IsPanic() {
			fmt.Println("Rollback transaction!")
			// Rollback transaction
		} else {
			fmt.Println("Commit transaction!")

			// Commit transaction
		}
		fmt.Println("UpdateTable defer end!")
	}()

	fmt.Println("UpdateTable start!")
	// Database update operation...
	panic("sql err!")
	fmt.Println("UpdateTable end!")
}

func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
		return true
	}

	return false
}

func UpdateTable_true() {
	// defer中决定提交还是回滚,recover直接被defer调用
	defer func() {
		fmt.Println("UpdateTable defer start!")
		if err := recover(); err != nil {
			fmt.Println("Rollback transaction!")
		} else {
			fmt.Println("Commit transaction!")
			// Commit transaction
		}
		fmt.Println("UpdateTable defer end!")
	}()

	fmt.Println("UpdateTable start!")
	// Database update operation...
	//panic(nil)	//情况1
	panic("sql err!")
	fmt.Println("UpdateTable end!")
}

func test_list() {
	alist := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range alist {
		fmt.Println(k, v)
		alist[4] = 100
		alist = append(alist, 12)
		fmt.Println(alist)
	}

}
