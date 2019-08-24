package api

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

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
