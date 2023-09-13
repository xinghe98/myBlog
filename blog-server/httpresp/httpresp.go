package httpresp

// Package HttpResp 封装响应方法，避免过多大括号

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Res 其他响应
func ResOthers(ctx *gin.Context, status int, data interface{}, msg string) {
	ctx.JSON(status, gin.H{
		"code": status,
		"data": data,
		"msg":  msg,
	})
}

// ResOK 请求成功的响应。
func ResOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}
