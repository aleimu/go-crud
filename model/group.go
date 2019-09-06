package model

import (
	"fmt"
	"time"
)

// 分组模型
type Group struct {
	Model
	Name   string `gorm:"not null;unique"`
	Status int    `gorm:"default:1"` // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)
}

func FindGroupById(ID interface{}) (Group, error) {
	var group Group
	result := DB.Find(&group, ID)
	fmt.Println("err:", result.Error, "result:", result)
	return group, result.Error
}

func FindGroups() []Group {
	var group []Group
	result := DB.Find(&group)
	fmt.Println("err:", result.Error, "result:", result)
	return group
}

func (group *Group) SetGroupInfo(name string, status int) error {
	group.Name = name
	group.Status = status
	group.UpdatedAt = time.Now()
	return nil
}

/*
err: sql: Scan error on column index 2, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
使用gorm框架，数据库使用的mysql  直接上解决办法 最后加上这个即可
db,err:=gorm.Open("mysql","用户名:密码@tcp(localhost:3306)/数据库名?charset=utf8&parseTime=true")
*/
