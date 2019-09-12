package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crud/util"
	"net/http"
)

func BeforeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("BeforeRequest Start !!!!")
		// 在gin上下文中定义变量,相当于flask中g的概念
		c.Set("example", "12345")
		// 以上为真正处理请求前
		c.Next() //处理请求,Next很特殊,Next之前的语句都是在进入handler之前执行的before_request,之后的语句是after_request执行的
		// server处理请求后
		fmt.Println("BeforeRequest End !!!!")
	}
}

// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": util.ERR, "msg": "server err awsl!", "data": fmt.Sprintf("内部错误:%s", err)})
			}
		}()
		c.Next()
	}
}
