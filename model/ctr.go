package model

import "time"

//  点击模型
type Ctr struct {
	Model
	Show       int
	Click      int
	Crt        float32
	ShowDay    string `gorm:"size:1000"`
	ClickDay   string `gorm:"size:1000"`
	Node       string `gorm:"size:1000"`
	CreateDate time.Time
}
