package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type ArticleCo struct {
	Article service.Cruder
}

// 实例化article的控制器
func NewArticleCO(Article service.Cruder) ArticleCo {
	return ArticleCo{Article}
}

// 查看所有文章的控制器
func (a ArticleCo) AllArticle(ctx *gin.Context) {
	a.Article.ReadAll(ctx)
}

// 更新单个文章的控制器
func (a ArticleCo) UpdateArticle(ctx *gin.Context) {
	a.Article.UpdateOne(ctx)
}
