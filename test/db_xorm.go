package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

// Database 在中间件中初始化mysql链接
func Database2() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	engine.SetMaxIdleConns(5)
	//打开
	engine.SetMaxOpenConns(10)
}

type User2 struct {
	Id    int64
	Name1 string `xorm:"varchar(10) notnull"`
	Name2 string `xorm:"varchar(20) notnull"`
	Name3 string `xorm:"varchar(30) notnull"`
}

func Test1() {
	Database2()
	//err := engine.Sync(new(User2))
	//fmt.Println("------->", err)
	user := new(User2)
	user.Name1 = "myname"
	user.Name2 = "myname2"
	user.Name3 = "myname3"
	affected, err := engine.Insert(user)
	fmt.Println("------->", affected, err)
	fmt.Println(user.Id)
	fmt.Println("-------------------------------")
	//user = new(User2)
	var user2 User2
	has, err := engine.Get(&user2)
	fmt.Println("------->", has, err, user2)
	fmt.Println("-------------------------------")
	var name string
	var user3 User2
	//user = new(User2)
	has, err = engine.Where("id = ?", 1).Cols("name1").Get(&user3)
	fmt.Println("------->", has, err, name, user3)


	users := make(map[int64]User2)
	err = engine.Find(&users)
	fmt.Println("------->", err, users)
	fmt.Println("-------------------------------")
	everyone := make([]User2, 0)
	err = engine.Find(&everyone)
	fmt.Println("------->", err, everyone)
	fmt.Println("-------------------------------")
}
