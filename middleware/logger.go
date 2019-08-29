package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-crud/conf"
	"os"
	"path"
	"time"
)

func Logger() *logrus.Logger {
	//日志文件
	fileName := path.Join(conf.LogFilePath, conf.LogFileName)
	fmt.Println("log file path:", fileName, conf.LogFilePath, conf.LogFileName)
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err.Error())
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	//设置日志格式
	logger := Logger()
	return func(c *gin.Context) {
		fmt.Println("log start !!!!")
		startTime := time.Now() // 处理请求
		c.Next()
		endTime := time.Now() // 结束时间
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		logger.Infof("| %3d | %3v | %3s | %s | %s |",
			c.Writer.Status(), // 状态码
			latencyTime,
			c.ClientIP(),         // 请求IP
			c.Request.Method,     // 请求方式
			c.Request.RequestURI, // 请求路由
		)
		fmt.Println("log end !!!!")
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

/*
不要把Context放在结构体中，要以参数的方式传递
以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
Context是线程安全的，可以放心的在多个goroutine中传递
*/

/* 日志分割什么的,暂时不启用 https://www.cnblogs.com/xinliangcoder/p/11212573.html
func Logger() gin.HandlerFunc {
	logClient := logrus.New()

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	apiLogPath := "api.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}
*/
