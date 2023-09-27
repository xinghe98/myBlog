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
	t.tagService.ReadWithAnother(ctx)
}
