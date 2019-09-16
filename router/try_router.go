package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Try() {
	// 自测分组,尝试新姿势
	v0 := r.Group("v1/test")
	{
		// 服务内部的重定向
		v0.GET("/test", func(c *gin.Context) {
			c.Request.URL.Path = "/v1/test/test2"
			r.HandleContext(c)
		})
		v0.GET("/test2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": "world"})
		})
		// 重定向到外部
		v0.GET("/test3", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "wwww.baidu.com")
		})
		// 静态文件目录,不适合单个文件查看
		v0.StaticFS("/image", http.Dir("./upload/"))
		//r.StaticFile("/image/:filename", "./upload/")
		//1. 异步
		v0.GET("/async", func(c *gin.Context) {
			// goroutine 中只能使用只读的上下文 c.Copy()
			cCp := c.Copy()
			go func() {
				time.Sleep(5 * time.Second)
				// 注意使用只读上下文
				logger.Println("Done! in path " + cCp.Request.URL.Path)
			}()
			c.JSON(http.StatusOK, gin.H{"user": gin.H{"a": gin.H{"b": "b",}, "Number": 123}, "Message": "hey", "Number": 123})
		})
		// 两个请求本质是两个goroutine,他们之间能相互传递信息吗? --- 单向可以,chan1请求发送后,chan2请求也发送了,chan1才会有返回;双向不行,类似于死锁
		ch1 := make(chan string)
		//ch2 := make(chan string)
		v0.GET("/chan1", func(c *gin.Context) {
			ch1 <- c.Query("ch")
			c.JSON(http.StatusOK, gin.H{"hello": "world"})
		})
		v0.GET("/chan2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"hello": <-ch1})
		})
	}
}
