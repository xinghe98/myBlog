package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func RegistRouters(r *gin.Engine) *gin.Engine {
	r = ExampleRouter(r, controller.CO)
	r = AdminRouter(r, controller.CO)
	return r
}
