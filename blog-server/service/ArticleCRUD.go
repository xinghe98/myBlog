package service

import (
	"myBlogServer/v1/httpresp"

	"github.com/gin-gonic/gin"
)

type Article struct{}

// NewArticle 实例化文章crud的函数
func NewArticle() *Article {
	return &Article{}
}

// CreateOne 创建一片文章
func (a *Article) CreateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// ReadAll 查找所有文章
func (a *Article) ReadAll(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

func (a *Article) ReadAny(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

// UpdateOne 更新一篇文章
func (a *Article) UpdateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// DeleteOne 删除一篇文章
func (a *Article) DeleteOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}
