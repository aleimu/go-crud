# Golang Fast CRUD

## 目的

本项目采用了Golang中比较流行的组件，最初的设想是作为基础快速搭建Restful Web API,其中也实现了不少功能点作为实例,探索了一些最佳实践可供参考. 

## 参考项目 

   - https://github.com/bydmm/singo   本项目的起源
   - https://github.com/gin-gonic/examples  gin 的实例,帮助学习gin的特性
   - https://github.com/EDDYCJY/go-gin-example gin实例,目录比较复杂
   - https://github.com/crawlab-team/crawlab   基于Golang的分布式爬虫管理平台,涉及内容丰富,值得学习
   - https://github.com/EDDYCJY/blog 煎鱼的博客,一些go相关的示例
   - https://github.com/wangsongyan/wblog 基于gin+gorm开发的个人博客项目
   - https://github.com/mydevc/go-gin-mvc 基于框架gin+xorm搭建的MVC项目架子
   - https://github.com/gitobhub/zhihu.git 仿知乎网站

## 依赖

本项目已经整合了许多开发API所必要的组件：

    1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
    2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
    3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
    4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
    5. [Godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
    6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
    7. [Cron](https://github.com/robfig/cron): 定时任务管理模块
    8. [Log](https://github.com/sirupsen/logrus) : 日志模块
    9. [Excelize](https://github.com/360EntSecGroup-Skylar/excelize): 360公司写的处理excel的库,应该值得信赖 [中文文档](https://xuri.me/excelize/zh-hans/)

## 项目结构

```
go-crud
    |-conf 配置文件目录
    |-router 文件夹就是MVC框架的controller，负责协调各部件完成任务
    |-server 负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
    |-serializer 储存通用的json模型和struct，把model得到的数据库模型转换成api需要的json对象
    |-cache 负责redis缓存相关的代码
    |-helpders 公共方法目录
    |-util 一些通用的小工具
    |-model 文件夹负责存储数据库模型和基础性的数据库操作相关的代码
    |-views 模板文件目录(暂时缺失)
    |-static 静态资源目录(暂时缺失)
        |-css css文件目录
        |-images 图片目录
        |-js js文件目录
        |-libs js类库
    |-system 系统配置文件加载目录
    |-tests 测试目录
    |-vendor 项目依赖其他开源项目目录
    |-main.go 程序执行入口
```

## TODO

本项目已经预先实现了一些常用的代码方便参考和复用:
- [X] 系统日志
- [X] 用户管理
    1. 创建了用户模型
    2. 实现了```/v1/user/register```用户注册接口
    3. 实现了```/v1/user/login```用户登录接口
    4. 实现了```/v1/user/me```用户资料接口(需要登录后获取session)
    5. 实现了```/v1/user/logout```用户登出接口(需要登录后获取session)
    6. 自行实现了国际化i18n的一些基本功能
    7. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证
- [X] 图片管理(上传/增/删/改/查)
- [X] 分组管理(增/删/改/查)
- [X] 广告管理(增/删/改/查)
- [ ] 定时任务,广告访问统计与归档(每个广告每小时的访问情况,点击情况)

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)
    
```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRE="" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```
    
```shell
# window下不可以带双引号
MYSQL_DSN=db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local # Mysql连接地址
REDIS_ADDR=127.0.0.1:6379 # Redis端口和地址
REDIS_PW= # Redis连接密码
REDIS_DB= # Redis库从0到10
SESSION_SECRE= # Seesion密钥，必须设置而且不要泄露
GIN_MODE=debug
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。
   
```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```
## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)