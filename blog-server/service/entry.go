package service

import "github.com/gin-gonic/gin"

type AdminInterface interface {
	Sigup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type Cruder interface {
	CreateOne(ctx *gin.Context)
	ReadAll(ctx *gin.Context)
	ReadAny(ctx *gin.Context)
	UpdateOne(ctx *gin.Context)
	DeleteOne(ctx *gin.Context)
}
