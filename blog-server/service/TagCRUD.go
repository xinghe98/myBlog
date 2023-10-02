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
	var tag models.Tags
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, util.TransLate(err), "数据不合法")
		fmt.Println(err)
		return
	}
	err = dao.DB.Create(&tag).Error
	if err != nil {
		if util.CheckDup(err) {
			httpresp.ResOthers(ctx, http.StatusBadGateway, util.TransLate(err), "该标签已存在")
			return
		}
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}

	httpresp.ResOK(ctx, nil)
}

// ReadAll 查找所有标签
func (a *TagCRUD) ReadAll(ctx *gin.Context) {
	var tag []models.Tags
	dao.DB.Find(&tag)
	httpresp.ResOK(ctx, tag)
}

// UpdateOne 更新一个标签
func (a *TagCRUD) UpdateOne(ctx *gin.Context) {
	tagid, ok := ctx.Params.Get("id")
	if !ok {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	var tag models.Tags
	dao.DB.Where("id=?", tagid).First(&tag)
	if tag.Name == "" {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, nil, "没有这标签")
		return
	}
	newTagName := ctx.PostForm("name")
	err := dao.DB.Model(&models.Tags{}).Where("id=?", tagid).Update("name", newTagName).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusNotExtended, nil, "未预期错误")
		return
	}
	httpresp.ResOK(ctx, gin.H{"msg": "更新成功"})
}

// DeleteOne 删除一个标签
func (a *TagCRUD) DeleteOne(ctx *gin.Context) {
	tagid, ok := ctx.Params.Get("id")
	if !ok {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	var tag models.Tags
	dao.DB.Where("id=?", tagid).First(&tag)
	if tag.Name == "" {
		httpresp.ResOthers(ctx, http.StatusMethodNotAllowed, nil, "没有这标签")
		return
	}
	err := dao.DB.Delete(&models.Tags{}, tagid).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "请求无效")
		return
	}
	httpresp.ResOK(ctx, gin.H{"msg": "删除成功"})
}

// 其他一些复杂的查询
// ReadWithAnother 查询所有标签，并且查询每个标签下的文章
// TODO: 这里要删掉，因为这个查询是为了测试gorm的preload功能，不然会影响性能
func (a *TagCRUD) ReadWithAnother(ctx *gin.Context) {
	var tag []models.Tags
	dao.DB.Preload("HasArt").Find(&tag)
	httpresp.ResOK(ctx, tag)
}
