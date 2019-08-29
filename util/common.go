package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	OK  = 1000
	ERR = 1500
)

//上传文件到指定的路径
var FilePath = os.Getenv("UPLOADFILE") + string(os.PathSeparator)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
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
	panic(err)
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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
