package service

import (
	"net/http"

	"myBlogServer/v1/dao"
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/models"

	"github.com/gin-gonic/gin"
)

type TagCRUD struct{}

// NewTags 实例化标签crud的函数
func NewTags() *TagCRUD {
	return &TagCRUD{}
}

// CreateOne 创建一个标签
func (a *TagCRUD) CreateOne(ctx *gin.Context) {
	var tag models.Tag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, nil, "数据不合法")
		return
	}
	err = dao.DB.Create(&tag).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
	}

	httpresp.ResOK(ctx, nil)
}

// ReadAll 查找所有标签
func (a *TagCRUD) ReadAll(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

func (a *TagCRUD) ReadAny(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}

// UpdateOne 更新一个标签
func (a *TagCRUD) UpdateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// DeleteOne 删除一个标签
func (a *TagCRUD) DeleteOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}
