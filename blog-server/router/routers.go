package router

import (
	"github.com/gin-gonic/gin"
)

func RegistRouters(r *gin.Engine) *gin.Engine {
	r = ExampleRouter(r)
	r = AdminRouter(r)
	return r
}
