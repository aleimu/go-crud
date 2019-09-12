package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 广告模型
type Style struct {
	Model
	GroupId   int    `json:"group_id;index" form:"group_id" json:"group_id" binding:"required"`
	ImageId   int    `json:"image_id" form:"image_id" json:"image_id" binding:"required"`
	ImageUrl  string `gorm:"type:varchar(256);" form:"image_url" json:"image_url" binding:"required"`
	ImageName string `gorm:"size:255" form:"image_name" json:"image_name" binding:"required"`
	Url       string `form:"url" json:"url" binding:"required"`
	OperId    int    `form:"oper_id" json:"oper_id" binding:"required"`
	OperName  string `form:"oper_id" json:"oper_id" binding:"required"`
	Status    int    `gorm:"default:1"` // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)  -> 同一个位置只能上架一个产品和图片展示方式有关
	Close     int    `gorm:"default:1"` // 是否可点击关闭(0：可关闭；1：不可关闭 )
	Mode      int    `gorm:"default:1"` // 图片展示方式: 1:轮播,2:横幅
	Frequency string `gorm:"default:2"` // 图片轮播的频率0.1-5s
	Position  int    `gorm:"default:1"` // 图片摆放位置: 1:首页banner,2:首页底部
	System    int    `gorm:"default:1"` // 1: 系统1, 2:系统2, 3:系统3, 4: 系统4
	Note      string                    // 备注-历次上下架的时间记录
	UpTime    time.Time                 // 上架时间
	DownTime  time.Time                 // 下架时间
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

func AddNewStyle(sf Style) error {
	result := DB.Create(&Style{GroupId: sf.GroupId, ImageId: sf.ImageId, ImageUrl: sf.ImageUrl, ImageName: sf.ImageName,
		Url: sf.Url, OperId: sf.OperId, OperName: sf.OperName, UpTime: time.Now(), DownTime: time.Now()})
	return result.Error

}

// 这样用循环,感觉复杂了
func GetSystems2() []string {
	var tmp []byte
	var s []string
	rows, err := DB.Model(&Style{}).Select("distinct(system)").Rows()
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&tmp)
		s = append(s, string(tmp))
	}
	fmt.Println("tmp:", tmp, rows, err, s)
	return s
}

type Temp1 struct {
	System int
}

// 单列最合适的用法
func GetSystems3() []int {
	var s []string
	var i []int
	var t []Temp1
	// Pluck，查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	//rows := DB.Model(&Style{}).Pluck("distinct(system)", &s) // list {"code":1000,"data":["3","1","2"]}
	rows := DB.Model(&Style{}).Pluck("distinct(system)", &i) // {"code":1000,"data":[3,1,2]}  看来Pluck会转换类型
	//rows := DB.Model(Style{}).Select("distinct(system)").Scan(&t) // 会扫描到结构体中转成map {"code":1000,"data":[{"System":3},{"System":1},{"System":2}]}
	fmt.Println("rows:", rows, s, t, i)
	return i
}

// 设定返回的格式
type Temp struct {
	System    int
	ImageName string
	CreatedAt time.Time
}

// 测试下怎么选择部分字段返回
func GetSystems4() []Temp {
	var t []Temp
	var s []Style
	//rows := DB.Model(&Style{}).Select("system", "image_name", "created_at").Find(&s) // SELECT system FROM `styles`  'image_name' ->错误的方式,Select内的string会被当做table
	//rows := DB.Model(&Style{}).Select([]string{"system", "image_name", "created_at"}).Find(&s) //  SELECT system, image_name, created_at FROM `styles` ->返回的是[]Style,但是字段太多,不合适
	//rows := DB.Model(&Style{}).Select([]string{"system", "image_name", "created_at"}).Find(&t) //   SELECT system, image_name, created_at FROM `temps` -> 错误的方式,会找错table,Model(&Style{})的设定会被Find(&t)替代掉
	// rows := DB.Table("styles").Select([]string{"system", "image_name", "created_at"}).Find(&t) //   SELECT system, image_name, created_at FROM `styles` ->可以按需返回,但是Table("styles")不好看
	//rows := DB.Select([]string{"system", "image_name", "created_at"}).Find(&t).Model(&Style{}) //   SELECT system, image_name, created_at FROM `temps` -> 错误的方式,看来Find优先级比Model高
	//rows := DB.Raw("select system,image_name,created_at from styles").Scan(&t) // select system,image_name,created_at from styles -> 使用sql语句查询,返回 [map],挺好的,但是不orm
	rows := DB.Model(&Style{}).Select([]string{"system", "image_name", "created_at"}).Scan(&t) //   SELECT system, image_name, created_at FROM `temps` ->返回[map] 挺好的,也orm风格
	//rows := DB.Model(Style{}).Select([]string{"system", "image_name", "created_at"}).Scan(&t) //   SELECT system, image_name, created_at FROM `temps` ->返回[map] 挺好的,也orm风格,同上,并不一定需要&Style{},不清楚会有什么影响
	fmt.Println("eerrr1:", rows, s)
	return t
	/*	期望返回如下的json给前端
		{"code":1000,"data":[{"System":3,"ImageName":"qqqq","CreatedAt":"2019-08-29T02:31:38Z"},{"System":3,"ImageName":"qqqq","CreatedAt":"2019-08-29T02:31:39Z"}}]}
	*/
}

func GetSystems() []string {
	var tmp []byte
	var s []string
	//rows, err := DB.Model(&Style{}).Select("distinct(system)").Rows()
	rows, err := DB.Model(&Style{}).Select("system,image_name,created_at").Rows()
	fmt.Println("eerrr1:", rows, err)
	if err != nil {
		panic(err.Error)
	}
	defer rows.Close()
	ss, err := List(rows, tmp)

	fmt.Println("tmp:", tmp, rows, err, s, ss)
	return s
}

// 通用的解析rows返回的函数
func List(rows *sql.Rows, unit interface{}) (units []interface{}, err error) {
	for rows.Next() {
		err = DB.ScanRows(rows, &unit)
		if err != nil {
			fmt.Println("eerrr2:", rows, err.Error())
			panic(err.Error)
		}
		err = rows.Scan(&unit)
		if err != nil {
			fmt.Println("eerrr3:", rows, err.Error())
			panic(err.Error)
		}
		units = append(units, &unit)
	}
	return
}
