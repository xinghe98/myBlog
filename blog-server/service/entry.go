package service

import "github.com/gin-gonic/gin"

type AdminInterface interface {
	Sigup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type Cruder interface {
	CreateOne(ctx *gin.Context)
	ReadAll(ctx *gin.Context)
	UpdateOne(ctx *gin.Context)
	DeleteOne(ctx *gin.Context)
	MazyFinder
}

type MazyFinder interface {
	ReadWithAnother(ctx *gin.Context)
}

type ImgCDer interface {
	UploadImg(ctx *gin.Context)
	DeleteImg(ctx *gin.Context)
	ReadAllHeadImgArt(ctx *gin.Context)
}
