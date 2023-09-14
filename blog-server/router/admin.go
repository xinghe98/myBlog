package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine, c *controller.BlogCo) *gin.Engine {
	admin := r.Group("/admin/")

	{
		admin.POST("regist", c.Admin.Regist)
		admin.POST("login", c.Admin.Login)
	}
	return r
}
