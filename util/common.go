package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 返回随机字符串
func RandStringRunes(n int) string {
	var a = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = a[rand.Intn(len(a))]
	}
	return string(b)
}

// 返回随机int字符串数组
func RandStringInt(n int) []string {
	var a = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	rand.Seed(time.Now().UnixNano())
	b := make([]string, n)
	for i := range b {
		b[i] = a[rand.Intn(len(a))]
	}
	return b
}

// 格式化时间
func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

func Str2Int(str string) int {
	tmp, err := strconv.Atoi(str)
	if err == nil {
		return tmp
	}
	panic("Str2Int err: " + err.Error())
}

func Int2Str(int int) string {
	tmp := strconv.Itoa(int)
	return tmp
}

func Str2Int64(str string) int64 {
	tmp, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return tmp
	}
	panic(err)
}

// 字符数组拼接成字符串
func ListStr2Str(strs []string) string {
	return strings.Join(strs, ",")
}

//保留float两位下小数
func Float2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func StrSum(str string) int {
	tmp := strings.Split(str, ",")
	sum := 0
	for _, v := range tmp {
		sum = sum + Str2Int(v)
	}
	return sum
}

func ListStrSum(str []string) int {
	sum := 0
	for _, v := range str {
		sum = sum + Str2Int(v)
	}
	return sum
}

func Today() time.Time {
	// 获取今天零点
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)

	fmt.Println(t1.Unix() + 1)
	fmt.Println(t2.AddDate(0, 0, 1).Unix())
	return t1
}

func Yesterday() time.Time {
	// 获取昨天零点
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	//t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	t2 := t1.AddDate(0, 0, -1)
	return t2
}
func GetRandomString(l int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

// 计算字符串的md5值
func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}

func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

func GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	return smtp.SendMail(host, auth, user, send_to, msg)
}

type errorString struct {
	s string
}

type errorInfo struct {
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

func (e *errorString) Error() string {
	return e.s
}
func Str2Map(jsonData string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Println(err.Error())
		panic("str2Map err: " + err.Error())
	}
	return result
}

func Str2Slice(jsonData string) (result []interface{}) {
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Println(err.Error())
		panic("Str2Slice err: " + err.Error())
	}
	return result
}

func Map2Str(mapData interface{}) (result string) {
	resultByte, err := json.Marshal(mapData)
	result = string(resultByte)
	if err != nil {
		fmt.Println(err.Error())
		panic("map2Str err: " + err.Error())
	}
	return result
}

func Byte2Map(jsonData []byte) (result map[string]interface{}) {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println(err.Error())
		panic("Byte2Map err: " + err.Error())
	}
	return result
}
func Map2Byte(mapData interface{}) (result []byte) {
	resultByte, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err.Error())
		panic("Map2Byte err: " + err.Error())
	}
	return resultByte
}

func byte2str(a []byte) []string {
	b := []string{}
	for _, v := range a {
		b = append(b, string(v))
	}
	fmt.Println(a, "-----------", b)
	return b
}
