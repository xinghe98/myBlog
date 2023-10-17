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

// 添加文章的控制器
func (a ArticleCo) AddArticle(ctx *gin.Context) {
	a.Article.CreateOne(ctx)
}

// 删除文章的控制器
func (a ArticleCo) DeleteArticle(ctx *gin.Context) {
	a.Article.DeleteOne(ctx)
}

// 查看所有文章的控制器
func (a ArticleCo) AllArticle(ctx *gin.Context) {
	a.Article.ReadAll(ctx)
}

// 查看单个文章的控制器
func (a ArticleCo) OneArticle(ctx *gin.Context) {
	a.Article.ReadOne(ctx)
}

// 更新单个文章的控制器
func (a ArticleCo) UpdateArticle(ctx *gin.Context) {
	a.Article.UpdateOne(ctx)
}
