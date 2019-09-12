# go orm 的选择
设计成公有属性的目的是需要通过反射给结构体中的变量再赋值，如果是私有属性是报reflect的panic

    https://my.oschina.net/u/168737/blog/1531834

    https://github.com/volatiletech/sqlboiler   // 很特别,实现思路和其他的不同
    
    https://github.com/go-xorm/xorm
    https://github.com/go-xorm/cmd
    https://github.com/xormplus/xorm
    https://www.kancloud.cn/xormplus/xorm/167077   // 中文文档


    https://github.com/jinzhu/gorm
    

# json
    https://github.com/tidwall/gjson


```cgo
// 定义的struct可以作为table,form,json的response
type Config struct {
	Name   string `xorm:"not null pk varchar(20)" form:"name" json:"name"`
	Value  string `xorm:"varchar(1024)" form:"value" json:"value"`
	Label  string `xorm:"varchar(40)" form:"label" json:"label"`
	Format string `xorm:"varchar(10)" form:"format" json:"format"`
}


```


```cgo
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

// 怎么选择部分字段多列返回
func GetSystems4() []Temp {
	var t []Temp
	var s []Style
	//rows := DB.Model(&Style{}).Select("system", "image_name", "created_at").Find(&s) // SELECT system FROM `styles`  'image_name' ->错误的方式,Select内的string会被当做table
	//rows := DB.Model(&Style{}).Select([]string{"system", "image_name", "created_at"}).Find(&s) //  SELECT system, image_name, created_at FROM `styles` ->返回的是[]Style,但是字段太多,不合适
	//rows := DB.Model(&Style{}).Select([]string{"system", "image_name", "created_at"}).Find(&t) //   SELECT system, image_name, created_at FROM `temps` -> 错误的方式,会找错table,Model(&Style{})的设定会被Find(&t)替代掉
	// rows := DB.Table("styles").Select([]string{"system", "image_name", "created_at"}).Find(&t) //   SELECT system, image_name, created_at FROM `styles` ->可以按需返回,但是Table("styles")不好看
	//rows := DB.Select([]string{"system", "image_name", "created_at"}).Find(&t).Model(&Style{}) //   SELECT system, image_name, created_at FROM `temps` -> 错误的方式,看来Find优先级比Model高
	//rows := DB.Raw("select system,image_name,created_at from styles").Scan(&t) // select system,image_name,created_at from styles -> 使用sql语句查询,返回 [map],挺好的,但是不orm
	//rows := DB.Model(&Style{}).Select([]string{"system", "image_name", "created_at"}).Scan(&t) //   SELECT system, image_name, created_at FROM `temps` ->返回[map] 挺好的,也orm风格
	rows := DB.Model(Style{}).Select([]string{"system", "image_name", "created_at"}).Scan(&t) //   SELECT system, image_name, created_at FROM `temps` ->返回[map] 挺好的,也orm风格,同上,并不一定需要&Style{},不清楚会有什么影响
	fmt.Println("eerrr1:", rows, s)
	return t
	/*	期望返回如下的json给前端
		{"code":1000,"data":[{"System":3,"ImageName":"qqqq","CreatedAt":"2019-08-29T02:31:38Z"},{"System":3,"ImageName":"qqqq","CreatedAt":"2019-08-29T02:31:39Z"}}]}
	*/
}

```