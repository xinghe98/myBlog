package service

import (
	"myBlogServer/v1/httpresp"

	"github.com/gin-gonic/gin"
)

type Article struct{}

// 查找所有文章
func (a *Article) FindAll(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}
