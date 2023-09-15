package router

import (
	"myBlogServer/v1/controller"
	"myBlogServer/v1/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine, c *controller.BlogCo) *gin.Engine {
	admin := r.Group("/admin/")

	{
		admin.POST("regist", c.Admin.Regist) // 注册用户
		admin.POST("login", c.Admin.Login)   // 登录
	}
	article := r.Group("/article/")
	article.Use(middleware.JwtAuth())
	{
		article.GET("findall", c.Admin.AllArticle)
	}
	return r
}
