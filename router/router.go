package router

import (
	"go-crud/middleware"
	"go-crud/server"
	"go-crud/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	// 给表单限制上传大小 (默认 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 服务内部的重定向
	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"herllo": "world"})
	})
	r.StaticFS("/image", http.Dir("./upload/")) // 静态文件目录,不适合单个文件查看
	//r.StaticFile("/image/:filename", "./upload/")
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
		// 上传图片 curl -X POST http://localhost:3000/v1/image/upload -F "file=@a.jpg" -H "Content-Type: multipart/form-data"
		v2.POST("/image/upload", server.Upload)
		// 新增图片
		// 上传图片 curl -v -X POST "127.0.0.1:3000/v1/image" -d "group_id=1&name=aaa&url=bbb&status=1" -H "Content-Type: multipart/form-data"
		// 上传图片 curl -v -X POST "127.0.0.1:3000/v1/image" -d "{\"group_id\":1,\"name\":\"aaa\",\"url\":\"bbb\",\"status\":1}" -H "Content-Type: application/json"
		v2.POST("/image", server.AddImage)
		// 查询图片详情
		v2.GET("/image", server.GetImage2)
		// 删除图片
		v2.DELETE("/image", server.DelImage)
		// 删除图片
		v2.PUT("/image", server.UpdateImage)
		// 查看图片 127.0.0.1:3000/v1/image/a.jpg
		v2.GET("/upload/:filename", func(c *gin.Context) { // 这样的模糊匹配都允许和其他路径重复 如:/v1/image/:filename 就不允许
			filename := c.Param("filename") // 路径参数
			c.File(util.FilePath + filename)
		})
		// curl "127.0.0.1:3000/v1/group?id=2"
		v2.GET("/group", server.GetGroup)
		// curl -X POST "127.0.0.1:3000/v1/group" -d"name=aaa"
		v2.POST("/group", server.AddGroup)
		// curl -X PUT "127.0.0.1:3000/v1/group" -d"id=1&name=bbbb"
		v2.PUT("/group", server.SetGroup)
		// curl -X DELETE "127.0.0.1:3000/v1/group?id=1"
		v2.DELETE("/group", server.DelGroup)

		// curl "127.0.0.1:3000/v1/style?id=2"
		v2.GET("/style", server.GetStyle)
		// curl -X POST "127.0.0.1:3000/v1/style" -d"name=aaa"
		v2.POST("/style", server.AddStyle)
		// curl -X PUT "127.0.0.1:3000/v1/style" -d"id=1&name=bbbb"
		v2.PUT("/style", server.SetStyle)
		// curl -X DELETE "127.0.0.1:3000/v1/style?id=1"
		v2.DELETE("/style", server.DelStyle)
		// curl "127.0.0.1:3000/v1/styles?id=2"
		v2.GET("/styles", server.StyleList)
	}
	return r
}
