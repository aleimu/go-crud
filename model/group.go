package model

import "time"

// 分组模型
type Group struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null;unique"`
	status    int // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)
	CreatedAt time.Time
	UpdatedAt time.Time
}
