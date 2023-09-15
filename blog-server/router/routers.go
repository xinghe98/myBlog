package router

import (
	"myBlogServer/v1/controller"

	"github.com/gin-gonic/gin"
)

func RegistRouters(r *gin.Engine) *gin.Engine {
	r = ExampleRouter(r, controller.CO) // 注册测试时的路由
	r = AdminRouter(r, controller.CO)   // 注册生产实际的路由
	return r
}
