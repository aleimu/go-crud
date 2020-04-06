package test

import "fmt"

/*
Go语言中byte和rune实质上就是uint8和int32类型。
Go的字符称为rune，等价于C中的char，可直接与整数转换
rune实际是整型，必需先将其转换为string才能打印出来，否则打印出来的是一个整数
*/

func test() {
	var c rune = 'a'
	fmt.Println(c)
	fmt.Println(string(c))
	fmt.Println(string(97))
	fmt.Println("'a' convert to", int(c))

	var i int = 98
	c1 := rune(i)
	fmt.Println(c1)
	fmt.Println("98 convert to", string(c1))

	//string to rune
	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}

	fmt.Println(string([]uint8{49, 50, 51, 52}))
	fmt.Println(string([]byte{49, 50, 51, 52}))
	fmt.Println(string(312))
}

func test_float() {
	var a = 123;
	var b = 345;
	fmt.Println("--------------------------------------")
	fmt.Println("%0.2f:", (float32(a) / float32(b)))
	fmt.Println("--------------------------------------")

}

/* 查询的姿势1
func ListMaxCommentPost() (posts []*Post, err error) {
	var (
		rows *sql.Rows
	)
	rows, err = DB.Raw("select p.*,c.total comment_total from posts p inner join (select post_id,count(*) total from comments group by post_id) c on p.id = c.post_id order by c.total desc limit 5").Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		DB.ScanRows(rows, &post)
		posts = append(posts, &post)
	}
	return
}
*/
