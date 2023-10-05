package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type tagsco struct {
	tagService service.Cruder
}

func NewTagCo(tagService service.Cruder) tagsco {
	return tagsco{tagService}
}

// createTag 创建标签
func (t tagsco) CreateTag(ctx *gin.Context) {
	t.tagService.CreateOne(ctx)
}

// readTag 查看标签
func (t tagsco) ReadTag(ctx *gin.Context) {
	t.tagService.ReadAll(ctx)
}

// UpdateTag 更新标签
func (t tagsco) UpdateTag(ctx *gin.Context) {
	t.tagService.UpdateOne(ctx)
}

// DeleteTag 删除标签
func (t tagsco) DeleteTag(ctx *gin.Context) {
	t.tagService.DeleteOne(ctx)
}

// MazyFind 复杂查询：根据标签名查找文章
func (t tagsco) MazyFind(ctx *gin.Context) {
	t.tagService.ReadWithAnother(ctx)
}
