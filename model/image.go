package model

import "fmt"

//  图片模型
type Image struct {
	Model
	GroupId int
	Name    string
	Url     string // 图片的存储URL
	Status  int    // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)
}

// 校验form或json入参的字段格式
type ImageForm struct {
	Id      int    `form:"id" json:"id"`
	GroupId int    `form:"group_id" json:"group_id"`
	Name    string `form:"name" json:"name" binding:"required,min=2,max=100"`
	Url     string `form:"url" json:"url" binding:"required,min=2,max=100"`
	Status  int    `form:"status" json:"status"` // 状态(0：已下架 1：已上架 2:暂存未发布 3:已删除弃用)
}

func GetImage(sql Image) (Image, error) {
	var image Image
	result := DB.Where(sql).First(&image)
	//result := DB.Find(&image, ID)
	// db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// Struct
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
	// Map
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	// 主键的Slice
	//db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);
	fmt.Println("err:", result.Error, "result:", result)
	return image, result.Error
}

func GetImage2(sql map[string]interface{}) (Image, error) {
	var image Image
	// Map
	result := DB.Where(sql).Find(&image)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	fmt.Println("err:", result.Error, "result:", result)
	return image, result.Error
}

func AddNewImage(sf ImageForm) error {
	fmt.Println("sf:", sf)
	//result := DB.Create(&Image{GroupId: Str2Int(sf.GroupId), Name: sf.Name, Url: sf.Url, Status: Str2Int(sf.Status)})	// 使用json或form, 不需要自己做类型转换
	result := DB.Create(&Image{GroupId: sf.GroupId, Name: sf.Name, Url: sf.Url, Status: sf.Status})
	fmt.Println("result:", result, result.Error)
	return result.Error

}

// 批量
func UpdateImage2(sql ImageForm) error {
	fmt.Println("sql:", sql)
	//result := DB.Model(&Image{}).Where("id = ?", sql.Id).Updates(map[string]interface{}{"name": sql.Name, "url": sql.Url, "group_id": sql.GroupId, "status": sql.Status}) // map式批量更新
	result := DB.Model(Image{}).Where("id = ?", sql.Id).Updates(Image{Name: sql.Name, Url: sql.Url, GroupId: sql.GroupId, Status: sql.Status}) // model式批量更新
	fmt.Println("result:", result, result.Error, result.RowsAffected)
	return result.Error
}

// 分步
func UpdateImage(sql ImageForm) error {
	var image Image
	err := DB.Where("id= ? ", sql.Id).Find(&image)
	fmt.Println("err:", err.Error)
	if err.Error != nil {
		panic(err.Error)
	}
	image.Name = sql.Name
	image.Url = sql.Url
	image.GroupId = sql.GroupId
	image.Status = sql.Status
	err = DB.Save(&image)
	fmt.Println("sql:", sql, image)
	return err.Error

}

func DelImage(id interface{}) error {
	fmt.Println("id:", id)
	result := DB.Where("id = ?", id).Delete(Image{})
	fmt.Println("result:", result, result.Error)
	return result.Error

}

/*
Select
指定要从数据库检索的字段，默认情况下，将选择所有字段;
db.Select("name, age").Find(&users)
//SELECT name, age FROM users;
db.Select([]string{"name", "age"}).Find(&users)
//SELECT name, age FROM users;
*/
