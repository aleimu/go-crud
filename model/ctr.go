package model

import (
	"fmt"
	"time"
)

//  点击模型
type Ctr struct {
	Model
	StyleId    int
	Show       int
	Click      int
	Crt        float64
	ShowDay    string `gorm:"size:1000"`
	ClickDay   string `gorm:"size:1000"`
	Node       string `gorm:"size:1000"`
	CreateDate time.Time
}

func AddCtr(c Ctr) error {
	fmt.Println("c:", c)
	//result := DB.Create(&Ctr{Show: c.Show, Click: c.Click, ShowDay: c.ShowDay, ClickDay: c.ClickDay,
	//	CreateDate: c.CreateDate, Crt: float64(c.Click) / float64(c.Show)})
	result := DB.Create(&Ctr{Show: c.Show, Click: c.Click, ShowDay: c.ShowDay, ClickDay: c.ClickDay,
		CreateDate: c.CreateDate, Crt: 0, StyleId: c.StyleId})
	fmt.Println("result:", result, result.Error)
	return result.Error
}

func GetCtr(filter interface{}) (Ctr, error) {
	var ctr Ctr
	result := DB.Where(filter).First(&ctr)
	fmt.Println("err:", result.Error, "result:", result)
	return ctr, result.Error
}

func GetCtrs(filter interface{}) *[]Ctr {
	var ctr []Ctr
	result := DB.Where(filter).Find(&ctr)
	fmt.Println("err:", result.Error, "result:", result)
	if result.Error != nil {
		panic(result.Error)
	}
	return &ctr
}

func UpdateCtr(c Ctr) error {
	result := DB.Model(&Ctr{}).Where("style_id = ?", c.StyleId).
		Updates(&Ctr{Show: c.Show, Click: c.Click, ShowDay: c.ShowDay, ClickDay: c.ClickDay,
			CreateDate: c.CreateDate, Crt: float64(c.Click) / float64(c.Show)}) // model式批量更新
	fmt.Println("result:", result, result.Error, result.RowsAffected)
	if result.RowsAffected != 1 || result.Error != nil {
		return result.Error
	}
	return nil
}
