package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func ExampleRouter(r *gin.Engine) *gin.Engine {
	coEx := new(controller.ExampleCo)
	admin := r.Group("/test/")
	{
		admin.GET("ping", coEx.Pong)
	}
	return r
}
