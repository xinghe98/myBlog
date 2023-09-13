package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) *gin.Engine {
	adminco := new(controller.AdminCo)
	admin := r.Group("/admin/")

	{
		admin.POST("regist", adminco.Regist)
	}
	return r
}
