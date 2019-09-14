package router

import (
	"github.com/gin-gonic/gin"
	"go-crud/conf"
	"go-crud/middleware"
	"go-crud/server"
	"net/http"
	"os"
	"time"
)

// 全局的logger实例
var logger = middleware.Logger()

func NewApp() *gin.Engine {
	r := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiBr

	// 使用自定义中间件, 使用顺序需要注意
	r.Use(middleware.Recovery())      // 处理各种异常,保障返回为json CatchExpection
	r.Use(middleware.LoggerToFile())  // Logging to a file.
	r.Use(middleware.BeforeRequest()) // flask.g
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	r.NoRoute(go404)
	return r

}

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := NewApp()
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
		r.GET("/async", func(c *gin.Context) {
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

	// 用户路由
	v1 := r.Group("/v1/user")
	{
		v1.POST("ping", server.Ping)
		// 用户登录
		v1.POST("/register", server.UserRegister)
		// 用户登录
		v1.POST("/login", server.UserLogin)
		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("/me", server.UserMe)
			v1.DELETE("/logout", server.UserLogout)
		}
	}
	// 广告管理路由
	v2 := r.Group("/v1")
	{
		// 上传图片
		v2.POST("/upload", server.Upload) // 上传图片 curl -X POST http://localhost:3000/v1/upload -F "file=@a.jpg" -H "Content-Type: multipart/form-data"
		// 新增图片信息
		v2.POST("/image", server.AddImage) // 上传图片form curl -v -X POST "127.0.0.1:3000/v1/image" -d "group_id=1&name=aaa&url=bbb&status=1" -H "Content-Type: multipart/form-data"
		// 上传图片json curl -v -X POST "127.0.0.1:3000/v1/image" -d "{\"group_id\":1,\"name\":\"aaa\",\"url\":\"bbb\",\"status\":1}" -H "Content-Type: application/json"
		// 查询图片详情
		v2.GET("/image", server.GetImage2)
		// 删除图片
		v2.DELETE("/image", server.DelImage)
		// 删除图片
		v2.PUT("/image", server.UpdateImage)
		// 查看图片
		v2.GET("/image/:filename", func(c *gin.Context) { //127.0.0.1:3000/v1/image/a.jpg
			filename := c.Param("filename") // 路径参数
			c.File(conf.FilePath + filename)
		})
		// 查询分组详情
		v2.GET("/groups", server.GetGroups) // curl "127.0.0.1:3000/v1/group?id=2"
		// 查询分组详情
		v2.GET("/group", server.GetGroup) // curl "127.0.0.1:3000/v1/group?id=2"
		// 新增分组
		v2.POST("/group", server.AddGroup) // curl -X POST "127.0.0.1:3000/v1/group" -d"name=aaa"
		// 修改分组信息/状态
		v2.PUT("/group", server.UpdateGroup) //curl -X PUT "127.0.0.1:3000/v1/group" -d"id=1&name=bbbb"
		// 删除分组
		v2.DELETE("/group", server.DelGroup) //url -X DELETE "127.0.0.1:3000/v1/group?id=1"

		// 获取广告信息
		v2.GET("/style", server.GetStyle) //curl "127.0.0.1:3000/v1/style?id=2"
		// 新增广告模式
		v2.POST("/style", server.AddStyle) //curl -v -X POST "127.0.0.1:3001/v1/style" -d "group_id=1&image_id=1&image_url=www&image_name=qqqq&url=qqwqw&oper_id=123&oper_name=321"
		// 修改模式
		v2.PUT("/style", server.UpdateStyle) //curl -v -X PUT "127.0.0.1:3001/v1/style" -d "id=1&group_id=2&image_id=2&image_url=www&image_name=qqqq&url=qqwqw&oper_id=123&oper_name=321"
		// 删除
		v2.DELETE("/style", server.DelStyle) //curl -X DELETE "127.0.0.1:3000/v1/style?id=1"
		// 查询模式列表
		v2.GET("/styles", server.StyleList) //curl "127.0.0.1:3000/v1/styles?id=2"

		// todo
		v2.GET("/freshall", server.FreshAllRedis) // 获取广告列表
		v2.GET("/search", server.SearchStyle)     // 获取广告列表 curl "127.0.0.1:3000/v1/search"
		v2.GET("/show", server.AddHourShow)       // 增加对应广告的展示量 curl "127.0.0.1:3000/v1/show?id=2&count=10"
		v2.GET("/click", server.AddHourClick)     // 增加广告的点击量 curl "127.0.0.1:3000/v1/click?id=2&count=10"
		v2.GET("/system", server.GetSystems)      // 获取系统编号 curl "127.0.0.1:3000/v1/system"

	}
	return r
}

func go404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": "404", "msg": "page not found!", "data": nil})
	return
}
