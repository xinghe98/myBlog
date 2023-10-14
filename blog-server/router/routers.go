package router

import (
	"github.com/gin-gonic/gin"
)

func RegistRouters(r *gin.Engine) *gin.Engine {
	r = AdminRouter(r)  // 注册后端的路由
	r = CommonRouter(r) // 注册公共路由
	return r
}
