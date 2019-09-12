package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Json(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "msg": msg})
}

func JsonOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": OK, "data": data, "msg": "ok"})
}

func JsonFail(ctx *gin.Context, data interface{}, err interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": ERR, "data": data, "msg": err})
}

func JsonList(ctx *gin.Context, data interface{}, total int64) {
	ctx.JSON(http.StatusOK, gin.H{"code": OK, "rows": data, "msg": "", "total": total})
}
