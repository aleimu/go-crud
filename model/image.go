package model

import "time"

//  图片模型
type Image struct {
	ID        uint `gorm:"primary_key"`
	GroupId   int
	Name      string
	Url       string
	status    int // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)
	CreatedAt time.Time
	UpdatedAt time.Time
}
