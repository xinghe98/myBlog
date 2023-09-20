package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type ArticleCo struct {
	Article service.ArticleInterface
}

// 实例化article的控制器
func NewArticleCO(Article service.ArticleInterface) ArticleCo {
	return ArticleCo{Article}
}

// 查看所有文章的控制器
func (a ArticleCo) AllArticle(ctx *gin.Context) {
	a.Article.ReadAll(ctx)
}
