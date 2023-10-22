// Func: 公共路由
package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.Engine) *gin.Engine {
	// 所有headline
	r.GET("/headline", controller.ImgCD.GetAllHeadline)
	article := r.Group("/article/")
	{
		// 查询所有文章
		article.GET("findall", controller.ArticleCRUD.AllArticle)
		article.GET(":aid", controller.ArticleCRUD.OneArticle)
		article.GET("archive", controller.ArticleCRUD.ArchiveArticle)
	}
	tags := r.Group("/tags/")
	{

		// 查询所有标签
		tags.GET("findall", controller.TagsCRUD.ReadTag)
		tags.GET("findwith", controller.TagsCRUD.MazyFind)
		tags.GET(":name", controller.TagsCRUD.ReadOneTag)
	}
	return r
}
