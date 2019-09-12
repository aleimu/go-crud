## 常用的几种连接数据库的方式
db, err := sql.Open("mysql", "user:password@unix(/tmp/mysql.sock)/test")
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/test")   //指定IP和端口
db, err := sql.Open("mysql", "user:password@/test")  //默认方式

## curl上传图片的方式
curl -X POST http://localhost:3000/v1/image/upload -F "file=@a.jpg" -H "Content-Type: multipart/form-data"




## TODO
-  能不能让连个请求之间发生关系,每个请求都是一个go协程,能像普通协程之间一样通过chan影响行为吗?

    

## 反射三定律：
	反射可以将“接口类型变量”转换为“反射类型对象”。
	反射可以将“反射类型对象”转换为“接口类型变量”。
	如果要修改“反射类型对象”，其值必须是“可写的”（settable）。

## context的使用

    https://segmentfault.com/a/1190000006190038
    
    常见应用场景:
        由一个请求衍生出的各个goroutine之间需要满足一定的约束关系，以实现一些诸如有效期，中止routine树，传递请求全局变量之类的功能。
        使用context实现上下文功能约定需要在你的方法的传入参数的第一个传入一个context.Context类型的变量
    
    使用方法:
        在导入这个包之后，初始化Context对象，在每个资源访问方法中都调用它，然后在使用时检查Context对象是否已经被Cancel，如果是就释放绑定的资源。
        在中间过程传递，在资源相关操作时判断是否ctx.Done()传出了值.
    
        Context 的调用应该是链式的，通过WithCancel，WithDeadline，WithTimeout或WithValue派生出新的 Context。当父 Context 被取消时，其派生的所有 Context 都将取消。
        通过context.WithXXX都将返回新的 Context 和 CancelFunc。调用 CancelFunc 将取消子代，移除父代对子代的引用，并且停止所有定时器。
        未能调用 CancelFunc 将泄漏子代，直到父代被取消或定时器触发。go vet工具检查所有流程控制路径上使用 CancelFuncs。
    
        Done()，返回一个channel。当times out或者调用cancel方法时，将会close掉。
        Err()，返回一个错误。该context为什么被取消掉。
        Deadline()，返回截止时间和ok。
        Value()，返回值。
    
    所有方法:
        func Background() Context
        func TODO() Context
    
        func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
        func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
        func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
        func WithValue(parent Context, key, val interface{}) Context
    
    context的使用规范:
        1.不要把context存储在结构体中，而是要显式地进行传递
        2.把context作为第一个参数，并且一般都把变量命名为ctx :func DoSomething（ctx context.Context，arg Arg）error { // ... use ctx ... }
        3.就算是程序允许，也不要传入一个nil的context，如果不知道是否要用context的话，用context.TODO()来替代
        4.context.WithValue()只用来传递request作用域的数据(过渡进程和 Api 的请求范围的数据)，不要用它来传递可选参数，该数据必须是线程安全的。
        5.Context 对象是线程安全的，你可以把一个 Context 对象传递给任意个数的 gorotuine
    
    
    链接：
        https://segmentfault.com/a/1190000006744213 # 官方文档翻译,值得参考
        https://www.jianshu.com/p/0dc7596ba90a
        https://deepzz.com/post/golang-context-package-notes.html
        https://studygolang.com/articles/5131

