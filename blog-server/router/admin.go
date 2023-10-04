package router

import (
	"myBlogServer/v1/controller"
	"myBlogServer/v1/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) *gin.Engine {
	admin := r.Group("/admin/")

	{
		admin.POST("regist", controller.AdminResAndLogin.Sigup) // 注册用户
		admin.POST("login", controller.AdminResAndLogin.Login)  // 登录
	}
	// 文章相关
	article := r.Group("/article/")
	article.Use(middleware.JwtAuth())
	{
		article.GET("findall", controller.ArticleCRUD.AllArticle)
		article.POST("create", controller.ArticleCRUD.AddArticle)
		article.DELETE(":aid", controller.ArticleCRUD.DeleteArticle)
	}
	// 标签相关
	tags := r.Group("/tags/")
	tags.Use(middleware.JwtAuth())
	{
		tags.POST("create", controller.TagsCRUD.CreateTag)
		tags.GET("findall", controller.TagsCRUD.ReadTag)
		tags.PUT(":id", controller.TagsCRUD.UpdateTag)
		tags.DELETE(":id", controller.TagsCRUD.DeleteTag)
	}
	return r
}
