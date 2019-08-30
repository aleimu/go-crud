package conf

import (
	"go-crud/cache"
	"go-crud/model"
	"os"

	"github.com/joho/godotenv"
)

var (
	LogFilePath = "./logs/"
	LogFileName = "my.log"
	FilePath    = "./upload/"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err == nil {
		LogFilePath = os.Getenv("LogFilePath")
		LogFileName = os.Getenv("LogFileName")
		FilePath = os.Getenv("UPLOADFILE") + string(os.PathSeparator)
	}

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"), true)
	cache.Redis()
}
