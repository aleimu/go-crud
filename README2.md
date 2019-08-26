## 常用的几种连接数据库的方式
db, err := sql.Open("mysql", "user:password@unix(/tmp/mysql.sock)/test")
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/test")   //指定IP和端口
db, err := sql.Open("mysql", "user:password@/test")  //默认方式

## curl上传图片的方式
curl -X POST http://localhost:3000/v1/image/upload -F "file=@a.jpg" -H "Content-Type: multipart/form-data"





