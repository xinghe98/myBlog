package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type AdminCo struct{}

// 注册的控制器
func (a AdminCo) Regist(ctx *gin.Context) {
	service.LOC.AdminRL.Sigup(ctx)
}

// 登录的控制器
func (a AdminCo) Login(ctx *gin.Context) {
	service.LOC.AdminRL.Login(ctx)
}

// 查看所有文章的控制器
func (a AdminCo) AllArticle(ctx *gin.Context) {
	service.LOC.Article.FindAll(ctx)
}
