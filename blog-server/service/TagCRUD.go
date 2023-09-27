package service

import (
	"fmt"
	"net/http"

	"myBlogServer/v1/dao"
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/models"
	"myBlogServer/v1/util"

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
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, util.TransLate(err), "数据不合法")
		fmt.Println(err)
		return
	}
	err = dao.DB.Create(&tag).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}

	httpresp.ResOK(ctx, nil)
}

// ReadAll 查找所有标签
func (a *TagCRUD) ReadAll(ctx *gin.Context) {
	var tag []models.Tag
	dao.DB.Find(&tag)
	httpresp.ResOK(ctx, tag)
}

// UpdateOne 更新一个标签
func (a *TagCRUD) UpdateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// DeleteOne 删除一个标签
func (a *TagCRUD) DeleteOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// 其他一些复杂的查询
// ReadWithAnother 查询所有标签，并且查询每个标签下的文章
func (a *TagCRUD) ReadWithAnother(ctx *gin.Context) {
	var tag []models.Tag
	dao.DB.Preload("HasArt").Find(&tag)
	httpresp.ResOK(ctx, tag)
}
