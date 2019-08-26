package server

import "os"

const (
	OK  = 1000
	ERR = 1500
)
//上传文件到指定的路径
var FilePath = os.Getenv("UPLOADFILE") + string(os.PathSeparator)