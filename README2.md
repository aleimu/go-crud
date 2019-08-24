
curl "http://127.0.0.1:3000/advert/statistic?radio=0&start_date=2019-08-24+00:00:00&end_date=2019-08-24+23:59:59&sort=&id=89&format=hour&page_index=1&page_size=9999&token=e06fcaa4-c616-11e9-86d4-00163e050akk"



curl "http://127.0.0.1:3000/advert/statistic?radio=0&start_date=&end_date=&sort=&id=89&format=day&page_index=1&page_size=9999&token=e06fcaa4-c616-11e9-86d4-00163e050akk"



curl "http://127.0.0.1:3000/advert/show?code=e5902ce6-c4bc-11e9-86d4-00163e050a03"



curl "http://127.0.0.1:3000/advert/statistic?radio=1&start_date=2019-08-23+00:00:00&end_date=2019-08-23+23:59:59&sort=&id=76&format=hour&page_index=1&page_size=9999&token=e06fcaa4-c616-11e9-86d4-00163e050akk"



curl "http://127.0.0.1:3000/advert/statistic?radio=0&start_date=2019-07-02+00:00:00&end_date=2019-08-24+23:59:59&sort=&id=76&format=month&page_index=1&page_size=9999&token=e06fcaa4-c616-11e9-86d4-00163e050akk"



set MYSQL_DSN="toto:toto123@tcp(118.190.87.8:3306)/camel_test?charset=utf8"
set REDIS_ADDR="118.190.87.8:6379"
set REDIS_PW=""
set REDIS_DB="15"
set SESSION_SECRE=""

set MYSQL_DSN="root:123@tcp(127.0.0.1:3306)/test?charset=utf8"


set MYSQL_DSN=root:123@tcp(127.0.0.1:3306)/test?charset=utf8


set MYSQL_DSN="<toto>:<toto123>@118.190.87.8:3306/camel_test?charset=utf8&parseTime=True&loc=Local"


set MYSQL_DSN="toto:toto123@tcp(118.190.87.8:3306)/camel_test?charset=utf8"

set MYSQL_DSN="server:ptdAChu+byhzq2dCc0&MLd@tcp(115.28.27.231:3306)/camel_test?charset=utf8"

'mysql://server:ptdAChu+byhzq2dCc0&MLd@115.28.27.231:3306/camel_test?charset=utf8'


db, err := sql.Open("mysql", "user:password@unix(/tmp/mysql.sock)/test")
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/test")   //指定IP和端口
db, err := sql.Open("mysql", "user:password@/test")  //默认方式

mysqld.exe --install MySql --defaults-file="D:\soft\mysql-5.7\my.ini"



curl 
-F "pic=@/mnt/shared/Image/jpg/Screensho1t.jpg; filename='Screensho1t.jpg'" 
http://127.0.0.1:8080/picture




