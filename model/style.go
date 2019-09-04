package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 广告模型
type Style struct {
	Model
	GroupId   int    `json:"group_id;index"`
	ImageId   int    `json:"image_id"`
	ImageUrl  string `gorm:"type:varchar(256);"`
	ImageName string `gorm:"size:255"`
	Url       string
	OperId    int
	OperName  string
	Status    int    `gorm:"default:1"` // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)  -> 同一个位置只能上架一个产品和图片展示方式有关
	Close     int    `gorm:"default:1"` // 是否可点击关闭(0：可关闭；1：不可关闭 )
	Mode      int    `gorm:"default:1"` // 图片展示方式: 1:轮播,2:横幅
	Frequency string `gorm:"default:2"` // 图片轮播的频率0.1-5s
	Position  int    `gorm:"default:1"` // 图片摆放位置: 1:首页banner,2:首页底部
	System    int    `gorm:"default:1"` // 1: driver_advert, 2:dispatch_advert, 3:order_advert, 4: camel_advert
	Note      string                    // 备注-历次上下架的时间记录
	UpTime    time.Time                 // 上架时间
	DownTime  time.Time                 // 下架时间
}

type StyleForm struct {
	Id        int    `form:"id" json:"id"`
	GroupId   int    `form:"group_id" json:"group_id" binding:"required"`
	ImageId   int    `form:"image_id" json:"image_id" binding:"required"`
	ImageUrl  string `form:"image_url" json:"image_url" binding:"required"`
	ImageName string `form:"image_name" json:"image_name" binding:"required"`
	Url       string `form:"url" json:"url" binding:"required"`
	OperId    int    `form:"oper_id" json:"oper_id" binding:"required"`
	OperName  string `form:"oper_name" json:"oper_name"`
	Status    int    `form:"default:1"` // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)  -> 同一个位置只能上架一个产品和图片展示方式有关
	//Close     int    `form:"default:1"` // 是否可点击关闭(0：可关闭；1：不可关闭 )
	//Mode      int    `form:"default:1"` // 图片展示方式: 1:轮播,2:横幅
	//Frequency string `form:"default:2"` // 图片轮播的频率0.1-5s
	//Position  int    `form:"default:1"` // 图片摆放位置: 1:首页banner,2:首页底部
	//System    int    `form:"default:1"` // 1: driver_advert, 2:dispatch_advert, 3:order_advert, 4: camel_advert
	//Note      string                    // 备注-历次上下架的时间记录
}

func GetStyleById(ID interface{}) (Style, error) {
	var style Style
	result := DB.Find(&style, ID)
	fmt.Println("err:", result.Error, "result:", result)
	return style, result.Error
}

func GetStyleList(filter interface{}, skip int, limit int, sortKey string) ([]Style, error) {
	var style []Style
	var result *gorm.DB
	if skip == 0 && limit == 0 {
		result = DB.Where(filter).Find(&style)
	} else {
		result = DB.Where(filter).Offset(skip).Limit(limit).Order(sortKey).Find(&style)
	}
	fmt.Println("err:", result.Error, "result:", result)
	return style, result.Error
}

func GetStyleTotal(filter interface{}) int {
	var count int
	err := DB.Model(&Style{}).Where(filter).Count(&count).Error
	//err := DB.Table("styles").Where(filter).Count(&count).Error
	if err != nil {
		panic(err.Error)
	}
	return count
}

func AddNewStyle(sf StyleForm) error {
	result := DB.Create(&Style{GroupId: sf.GroupId, ImageId: sf.ImageId, ImageUrl: sf.ImageUrl, ImageName: sf.ImageName,
		Url: sf.Url, OperId: sf.OperId, OperName: sf.OperName, UpTime: time.Now(), DownTime: time.Now()})
	return result.Error

}
