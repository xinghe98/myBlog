package service

import (
	"myBlogServer/v1/httpresp"

	"github.com/gin-gonic/gin"
)

type TagCRUD struct{}

// NewArticle 实例化文章crud的函数
func NewTags() *TagCRUD {
	return &TagCRUD{}
}

// CreateOne 创建一片文章
func (a *TagCRUD) CreateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// ReadAll 查找所有文章
func (a *TagCRUD) ReadAll(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

func (a *TagCRUD) ReadAny(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

// UpdateOne 更新一篇文章
func (a *TagCRUD) UpdateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// DeleteOne 删除一篇文章
func (a *TagCRUD) DeleteOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}
