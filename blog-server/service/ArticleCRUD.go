package service

import (
	"fmt"
	"net/http"

	"myBlogServer/v1/dao"
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/models"
	"myBlogServer/v1/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	if len(tags) != len(tag) || len(tags) == 0 {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, util.TransLate(err), "请选择正确的标签")
		return
	}
	article.Tags = tags                 // 将查找到的tag赋值给文章的tag
	err = dao.DB.Create(&article).Error // 创建文章
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}

	httpresp.ResOK(ctx, nil)
}

// ReadAll 查找所有文章(分页)
// INFO: 这里的分页是通过前端传来的参数来实现的，如果前端不传参数，那么就是查找所有文章
func (a *Article) ReadAll(ctx *gin.Context) {
	pagination := util.GeneratePaginationFromRequest(ctx)
	var article models.Article
	var articles []*models.Article
	// 没有查询参数时查询所有文章
	if pagination.Limit == 0 && pagination.Page == 0 && len(pagination.Sort) == 0 {
		err := dao.DB.Preload("Tags", func(DB *gorm.DB) *gorm.DB {
			return DB.Debug().Omit("HasArt")
		}).Find(&articles).Error
		if err != nil {
			httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
			return
		}
		httpresp.ResOK(ctx, articles)
	} else {
		offset := (pagination.Page - 1) * pagination.Limit
		err := dao.DB.Preload("Tags", func(DB *gorm.DB) *gorm.DB {
			return DB.Debug().Omit("HasArt")
		}).Where(article).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&articles).Error
		if err != nil {
			httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
			return
		}
		httpresp.ResOK(ctx, articles)
	}
}

// UpdateOne 更新一篇文章
// FIX: 还没写完
func (a *Article) UpdateOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// DeleteOne 删除一篇文章
// FIX: 还没写完
func (a *Article) DeleteOne(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"code": 200})
}

// 下面是一些其他复杂的查询方法

// ReadWithAnother 通过其他条件查找文章
// FIX: 还没写完
func (a *Article) ReadWithAnother(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}
