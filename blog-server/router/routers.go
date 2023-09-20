package router

import (
	"github.com/gin-gonic/gin"
)

func RegistRouters(r *gin.Engine) *gin.Engine {
	r = AdminRouter(r) // 注册生产实际的路由
	return r
}
