# 11ptdAChu+byhzq2dCc0&MLd
# grant all on *.* to 'toto1'@'%' identified by 'toto123';
set GO111MODULE=off
set GOPROXY=https://goproxy.io


NaN 代表 不是一个数 Not a number
Inf  代表 阶码溢出，前面的加减符号代表高地位溢出，说白了就是小数点位后面无限大，再别的地方使用不能很好的序列化
解决办法:
使用fmt.Sprintf("%0.2f",浮点数) 输出一个字符串的固定位的浮点数，在把这个字符串转为float 即可
