package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func ExampleRouter(r *gin.Engine, c *controller.BlogCo) *gin.Engine {
	admin := r.Group("/test/")
	{
		admin.GET("ping", c.Example.Pong)
	}
	return r
}
