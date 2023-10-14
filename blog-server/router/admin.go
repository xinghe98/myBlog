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
		article.POST("create", controller.ArticleCRUD.AddArticle)
		article.GET("findall", controller.ArticleCRUD.AllArticle)
		article.PUT(":aid", controller.ArticleCRUD.UpdateArticle)
		article.DELETE(":aid", controller.ArticleCRUD.DeleteArticle)
	}
	// 标签相关
	tags := r.Group("/tags/")
	tags.Use(middleware.JwtAuth())
	{
		tags.POST("create", controller.TagsCRUD.CreateTag)
		tags.GET("findall", controller.TagsCRUD.ReadTag)
		tags.GET("findwith", controller.TagsCRUD.MazyFind)
		tags.PUT(":id", controller.TagsCRUD.UpdateTag)
		tags.DELETE(":id", controller.TagsCRUD.DeleteTag)
	}
	// 图片的上传与删除，因为只有两个方法且不会再有类似的控制器，所以没有使用interface做依赖注入
	img := r.Group("/img/")
	img.Use(middleware.JwtAuth())
	{
		img.POST("upload", controller.UploadImg)
		img.DELETE(":key", controller.DeleteImg)
	}
	return r
}
