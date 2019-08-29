package model

//执行数据迁移

func Merge() {
	// 自动迁移模式
	DB.AutoMigrate(&User{}, &Ctr{}, &Group{}, &Image{}, &Style{})
}
