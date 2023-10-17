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
	var tags []*models.Tags
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
	article.Tags = tags // 将查找到的tag赋值给文章的tag
	// 判断是否存在headerlineimg
	if article.HeadImg != "" {
		article.Image = "https://blog-1308532731.cos.ap-guangzhou.myqcloud.com/" + article.HeadImg
	}
	err = dao.DB.Create(&article).Error // 创建文章
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}

	httpresp.ResOK(ctx, nil)
}

// ReadAll 查找所有文章(分页)
// INFO: 这里的分页是通过前端传来的参数来实现的，如果前端不传参数，那么就是查找所有文章
// NOTE: 该service实现了4种查询方式，1.查询所有已发布文章,2.查询所有草稿箱文章,3.以上两种的分页查询,4.status值为0时查询所有文章(包括已发布和草稿箱)
func (a *Article) ReadAll(ctx *gin.Context) {
	pagination := util.GeneratePaginationFromRequest(ctx)
	var article models.Article
	var articles []models.Article
	var total int64

	// 查询所有已发布文章总数
	dao.DB.Model(&article).Where("status = ?", 1).Count(&total)

	// 没有查询参数时查询所有文章
	if pagination.Limit == 0 && pagination.Page == 0 && len(pagination.Sort) == 0 && pagination.Status == 0 {
		err := dao.DB.Preload("Tags", func(DB *gorm.DB) *gorm.DB {
			// select会将其他字段赋值为零值并返回，所以模型结构体需要加上json-tag
			return DB.Debug().Select("ID", "Name")
		}).Find(&articles).Error
		if err != nil {
			httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
			return
		}
		httpresp.ResOK(ctx, gin.H{"articles": articles, "total": total, "current_page_size": len(articles), "page": pagination.Page, "limit": pagination.Limit})
	} else if pagination.Limit == 0 && pagination.Page == 0 && len(pagination.Sort) == 0 && pagination.Status != 0 {
		// 查询所有文章(根据status)
		err := dao.DB.Model(&article).Preload("Tags", func(DB *gorm.DB) *gorm.DB {
			return DB.Debug().Select("ID", "Name")
		}).Where("status = ?", pagination.Status).Find(&articles).Error
		if err != nil {
			httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
			return
		}
		httpresp.ResOK(ctx, gin.H{"articles": articles, "total": total, "current_page_size": len(articles), "page": pagination.Page, "limit": pagination.Limit})
	} else {
		// 分页查询
		offset := (pagination.Page - 1) * pagination.Limit
		err := dao.DB.Model(&article).Preload("Tags", func(DB *gorm.DB) *gorm.DB {
			return DB.Debug().Select("ID", "Name")
		}).Where("status = ?", pagination.Status).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&articles).Error
		if err != nil {
			httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
			return
		}
		httpresp.ResOK(ctx, gin.H{"articles": articles, "total": total, "current_page_size": len(articles), "page": pagination.Page, "limit": pagination.Limit})
	}
}

// ReadOne 查找一篇文章
func (a *Article) ReadOne(ctx *gin.Context) {
	articleid, ok := ctx.Params.Get("aid")
	if !ok {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	var article models.Article
	dao.DB.Where("id=?", articleid).Preload("Tags").First(&article)
	if article.Title == "" {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, nil, "没有这篇文章")
		return
	}
	httpresp.ResOK(ctx, gin.H{"articles": article})
}

// UpdateOne 更新一篇文章
func (a *Article) UpdateOne(ctx *gin.Context) {
	articleid, ok := ctx.Params.Get("aid")
	if !ok {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	var oldarticle models.Article
	var newarticle models.Article
	var tags []*models.Tags
	dao.DB.Where("id=?", articleid).First(&oldarticle)
	if oldarticle.Title == "" {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, nil, "没有这篇文章")
		return
	}
	err := ctx.ShouldBindJSON(&newarticle)
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, util.TransLate(err), "数据不合法")
		fmt.Println(err)
		return
	}
	taglist := newarticle.Tags // 获取前端传来的tag数组
	tag := make([]string, len(taglist))
	for i, v := range taglist {
		tag[i] = v.Name
	}
	dao.DB.Find(&tags, "name in ?", tag) // 查找数据库中是否有这些tag
	if len(tags) != len(tag) || len(tags) == 0 {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, util.TransLate(err), "请选择正确的标签")
		return
	}

	// 1. 更新标签需要先删除原因映射关系
	err = dao.DB.Preload("Tags").Take(&oldarticle, articleid).Error
	dao.DB.Model(&oldarticle).Association("Tags").Delete(oldarticle.Tags)

	// 2. 再重建映射关系
	dao.DB.Model(&oldarticle).Association("Tags").Replace(tags)

	// 3. 最后更新文章的其他内容
	if newarticle.HeadImg != "" {
		newarticle.Image = "https://blog-1308532731.cos.ap-guangzhou.myqcloud.com/" + newarticle.HeadImg
	}
	err = dao.DB.Model(&oldarticle).Updates(&newarticle).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusNotExtended, nil, "未预期错误")
		return
	}
	httpresp.ResOK(ctx, gin.H{"msg": "更新成功"})
}

// DeleteOne 删除一篇文章
func (a *Article) DeleteOne(ctx *gin.Context) {
	articleID, ok := ctx.Params.Get("aid")
	if !ok {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	var article models.Article
	dao.DB.Where("id=?", articleID).Preload("Tags").First(&article)
	if article.Title == "" {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, nil, "没有这篇文章")
		return
	}
	// 删除文章的同时删除文章与标签的映射关系
	err := dao.DB.Model(&article).Association("Tags").Delete(article.Tags)
	err = dao.DB.Delete(&models.Article{}, articleID).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	httpresp.ResOK(ctx, gin.H{"msg": "删除成功"})
}

// 下面是一些其他复杂的查询方法

// ReadWithAnother 通过其他条件查找文章
// FIX: 还没写完
func (a *Article) ReadWithAnother(ctx *gin.Context) {
	httpresp.ResOK(ctx, gin.H{"dfaf": 200})
}
