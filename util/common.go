package util

import (
	"math/rand"
	"os"
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

func GetRandomString(l int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

//func GetRandomString() string {
//	b := make([]byte, 10)
//	_, err := rand.Read(b)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(time.Now().UnixNano())
//	return string(b)
//}
