package model

import "time"

// 广告模型
type Style struct {
	ID        uint `gorm:"primary_key"`
	GroupId   int
	ImageId   int
	ImageUrl  string
	ImageName string
	Url       string
	OperId    int
	OperName  string
	status    int       // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)  -> 同一个位置只能上架一个产品和图片展示方式有关
	close     int       // 是否可点击关闭(0：可关闭；1：不可关闭 )
	mode      int       // 图片展示方式: 1:轮播,2:横幅
	frequency string    // 图片轮播的频率0.1-5s
	position  int       // 图片摆放位置: 1:首页banner,2:首页底部
	system    int       // 1: driver_advert, 2:dispatch_advert, 3:order_advert, 4: camel_advert
	note      string    // 备注-历次上下架的时间记录
	UpTime    time.Time // 上架时间
	DownTime  time.Time // 下架时间
	CreatedAt time.Time
	UpdatedAt time.Time
}
