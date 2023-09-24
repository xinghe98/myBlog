package service

import "github.com/gin-gonic/gin"

type AdminInterface interface {
	Sigup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type ArticleInterface interface {
	ReadAll(ctx *gin.Context)
	UpdateOne(ctx *gin.Context)
}
