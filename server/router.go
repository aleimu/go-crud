package server

import (
	"go-crud/api"
	"go-crud/middleware"
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
	// 用户路由
	v1 := r.Group("/v1/user")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("/register", api.UserRegister)

		// 用户登录
		v1.POST("/login", api.UserLogin)

		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("/me", api.UserMe)
			v1.DELETE("/logout", api.UserLogout)
		}
	}
	// 广告管理路由
	v2 := r.Group("/v1")
	{
		// 上传图片 curl -X POST http://localhost:3000/v1/image/upload -F "file=@a.jpg" -H "Content-Type: multipart/form-data"
		v2.POST("image/upload", api.Upload)

		// 查询图片详情
		v2.GET("image/get", api.UserLogin)

	}
	return r
}
