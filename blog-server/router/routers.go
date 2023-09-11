package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func RegistRouters(r *gin.Engine) *gin.Engine {
	admin := r.Group("/admin/")
	{
		admin.GET("ping", controller.Pong)
	}
	return r
}
