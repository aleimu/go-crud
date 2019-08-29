package router

import (
	"github.com/gin-gonic/gin"
	"go-crud/middleware"
	"go-crud/server"
	"go-crud/util"
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
	r.Use(middleware.LoggerToFile())  // Logging to a file.
	r.Use(middleware.BeforeRequest()) // flask.g
	r.Use(middleware.Recovery())      // 处理各种异常,保障返回为json CatchExpection
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
		v0.GET("/test3", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "wwww.baidu.com")
		})
		v0.StaticFS("/image", http.Dir("./upload/")) // 静态文件目录,不适合单个文件查看
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
		// 上传图片 curl -X POST http://localhost:3000/v1/upload -F "file=@a.jpg" -H "Content-Type: multipart/form-data"
		v2.POST("/upload", server.Upload)
		// 新增图片
		// 上传图片form curl -v -X POST "127.0.0.1:3000/v1/image" -d "group_id=1&name=aaa&url=bbb&status=1" -H "Content-Type: multipart/form-data"
		// 上传图片json curl -v -X POST "127.0.0.1:3000/v1/image" -d "{\"group_id\":1,\"name\":\"aaa\",\"url\":\"bbb\",\"status\":1}" -H "Content-Type: application/json"
		v2.POST("/image", server.AddImage)
		// 查询图片详情
		v2.GET("/image", server.GetImage2)
		// 删除图片
		v2.DELETE("/image", server.DelImage)
		// 删除图片
		v2.PUT("/image", server.UpdateImage)
		// 查看图片 127.0.0.1:3000/v1/image/a.jpg
		v2.GET("/image/:filename", func(c *gin.Context) {
			filename := c.Param("filename") // 路径参数
			c.File(util.FilePath + filename)
		})
		// curl "127.0.0.1:3000/v1/group?id=2"
		v2.GET("/group", server.GetGroup)
		// curl -X POST "127.0.0.1:3000/v1/group" -d"name=aaa"
		v2.POST("/group", server.AddGroup)
		// curl -X PUT "127.0.0.1:3000/v1/group" -d"id=1&name=bbbb"
		v2.PUT("/group", server.UpdateGroup)
		// curl -X DELETE "127.0.0.1:3000/v1/group?id=1"
		v2.DELETE("/group", server.DelGroup)

		// curl "127.0.0.1:3000/v1/style?id=2"
		v2.GET("/style", server.GetStyle)
		// curl -v -X POST "127.0.0.1:3001/v1/style" -d "group_id=1&image_id=1&image_url=www&image_name=qqqq&url=qqwqw&oper_id=123&oper_name=321"
		v2.POST("/style", server.AddStyle)
		// curl -v -X PUT "127.0.0.1:3001/v1/style" -d "id=1&group_id=2&image_id=2&image_url=www&image_name=qqqq&url=qqwqw&oper_id=123&oper_name=321"
		v2.PUT("/style", server.UpdateStyle)
		// curl -X DELETE "127.0.0.1:3000/v1/style?id=1"
		v2.DELETE("/style", server.DelStyle)
		// curl "127.0.0.1:3000/v1/styles?id=2"
		v2.GET("/styles", server.StyleList)

		// 任务
		v2.GET("/tasks", server.GetStyle)        // 任务列表
		v2.GET("/tasks/:id", server.AddStyle)    // 任务详情
		v2.PUT("/tasks", server.UpdateStyle)     // 派发任务
		v2.DELETE("/tasks/:id", server.DelStyle) // 删除任务
	}
	return r
}

func go404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": "404", "msg": "page not found!", "data": nil})
	return
}
