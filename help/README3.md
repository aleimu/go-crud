## 死锁的几种情况

https://blog.csdn.net/benben_2015/article/details/84842486

- 同一个goroutine中，使用同一个 channel 读写。
```go
package main
func main(){
    ch:=make(chan int)  //这就是在main程里面发生的死锁情况
    ch<-6   //  这里会发生一直阻塞的情况，执行不到下面一句
    <-ch
}
//对于同一无缓冲通道，在接收者未准备好之前，发送操作是阻塞的。而此处的通道ch1就是缺少一个配对的接收者，因此造成了死锁。
//解决上面问题的方式有两种：第一种添加配对的接收者；第二种将默认的通道替换成缓冲通道。
```

- 2个 以上的go程中， 使用同一个 channel 通信。 读写channel 先于 go程创建。
```go
package main

func main(){
    ch:=make(chan int)
    ch<-666    //这里一直阻塞，运行不到下面
    go func (){
        <-ch  //这里虽然创建了子go程用来读出数据，但是上面会一直阻塞运行不到下面
    }()
}
//这里如果想不成为死锁那匿名函数go程就要放到ch<-666这条语句前面 
```

- 2个以上的go程中互相等对方造成死锁，使用多个 channel 通信。 A go 程 获取channel 1 的同时，尝试使用channel 2， 同一时刻，B go 程 获取channel 2 的同时，尝试使用channel 1
```go
package main
func main()  {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {    //匿名子go程
        for {
            select {    //这里互相等对方造成死锁
            case <-ch1:   //这里ch1有数据读出才会执行下一句
                ch2 <- 777
            }
        }
    }()
    for {         //主go程
        select {
        case <-ch2 : //这里ch2有数据读出才会执行下一句
            ch1 <- 999
        }
    }
}
```

- 在go语言中， channel 和 读写锁、互斥锁 尽量避免交叉混用。——“隐形死锁”。如果必须使用。推荐借助“条件变量”
```go
package main

import (
    "runtime"
    "math/rand"
    "time"
    "fmt"
    "sync"
)
// 使用读写锁
var rwMutex2 sync.RWMutex

func readGo2(idx int, in <-chan int)  {     // 读go程
    for {
        time.Sleep(time.Millisecond * 500)      // 放大实验现象// 一个go程可以读 无限 次。
        rwMutex2.RLock()    // 读模式加  读写锁
        num := <-in         // 从 公共的 channel 中获取数据
        fmt.Printf("%dth 读 go程，读到：%d\n", idx, num)
        rwMutex2.RUnlock()  // 解锁 读写锁
    }
}

func writeGo2(idx int, out chan<- int)  {
    for {                                   // 一个go程可以写 无限 次。
        // 生产一个随机数
        num := rand.Intn(500)
        rwMutex2.Lock()     // 写模式加  读写锁
        out <- num
        fmt.Printf("-----%dth 写 go程，写入：%d\n", idx, num)
        rwMutex2.Unlock()   // 解锁  读写锁

        //time.Sleep(time.Millisecond * 200)        // 放大实验现象
    }
}

func main()  {
    // 播种随机数种子。
    rand.Seed(time.Now().UnixNano())

    // 创建 模拟公共区的 channel
    ch := make(chan int, 5)

    for i:=0; i<5; i++ {        // 同时创建 N 个 读go程
            go readGo2(i+1, ch)
    }
    for i:=0; i<5; i++ {        // 同时创建 N 个 写go程
        go writeGo2(i+1, ch)
    }
    for {                       // 防止 主 go 程 退出
        runtime.GC()
    }
}

```
