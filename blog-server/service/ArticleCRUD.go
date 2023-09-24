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

// ReadAll 查找所有文章
func (a *Article) ReadAll(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

// UpdateOne 更新一篇文章
func (a *Article) UpdateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}
