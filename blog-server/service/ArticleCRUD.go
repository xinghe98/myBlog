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

type Article struct{}

// NewArticle 实例化文章crud的函数
func NewArticle() *Article {
	return &Article{}
}

// CreateOne 创建一片文章
func (a *Article) CreateOne(ctx *gin.Context) {
	var article models.Article
	var tags []*models.Tag
	err := ctx.ShouldBindJSON(&article) // 绑定前端的json数据
	if err != nil {
		fmt.Println(err)
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, util.TransLate(err), "数据不合法")
		return
	}
	taglist := article.Tags // 获取前端传来的tag数组
	tag := make([]string, len(taglist))
	for i, v := range taglist {
		tag[i] = v.Name
	}
	dao.DB.Find(&tags, "name in ?", tag) // 查找数据库中是否有这些tag
	article.Tags = tags                  // 将查找到的tag赋值给文章的tag
	err = dao.DB.Create(&article).Error  // 创建文章
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}

	httpresp.ResOK(ctx, nil)
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
