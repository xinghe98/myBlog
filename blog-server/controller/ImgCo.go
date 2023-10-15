package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type imgCo struct {
	imgCDService service.ImgCDer
}

func NewImgCo(imgCDService service.ImgCDer) *imgCo {
	return &imgCo{imgCDService}
}

func (i *imgCo) UploadImg(ctx *gin.Context) {
	i.imgCDService.UploadImg(ctx)
}

func (i *imgCo) GetAllHeadline(ctx *gin.Context) {
	i.imgCDService.ReadAllHeadImgArt(ctx)
}

func (i *imgCo) DeleteImg(ctx *gin.Context) {
	i.imgCDService.DeleteImg(ctx)
}
