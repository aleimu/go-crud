package util

import (
	"errors"
	"time"
)

type Page struct {
	//Kword    string    `form:"kword"  json:"kword"`
	Datefrom time.Time `form:"datefrom" time_format:"2006-01-02 15:04:05"`
	Dateto   time.Time `form:"dateto" time_format:"2006-01-02 15:04:05"`
	Size     int       `form:"pagesize" json:"size"`
	Index    int       `form:"pagefrom" json:"index"`
	Desc     string    `form:"desc" json:"desc"`
	Asc      string    `form:"asc" json:"asc"`
}

func (p *Page) Validate() (bool, error) {
	if p.Datefrom.IsZero() {
		return false, errors.New("请输入开始时间")
	}
	if p.Size > 100 {
		return false, errors.New("一次只能请求100条数据")
	}
	if p.Index < 0 {
		return false, errors.New("分页参数错误")
	}
	return true, nil
}

func (p *Page) GetSize() int {
	if p.Size == 0 {
		return 20
	}
	return p.Size
}

func (p *Page) GetIndex() int {
	return p.Index
}

func (p *Page) GetDesc() string {
	return p.Desc
}

func (p *Page) GetAsc() string {
	return p.Asc
}
