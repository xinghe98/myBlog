package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type AdminCo struct {
	AdminInterface service.AdminInterface
}

// 注册的控制器
func (a AdminCo) Sigup(ctx *gin.Context) {
	a.AdminInterface.Sigup(ctx)
}

// 登录的控制器
func (a AdminCo) Login(ctx *gin.Context) {
	a.AdminInterface.Login(ctx)
}

func NewAdmin(Admin service.AdminInterface) AdminCo {
	return AdminCo{Admin}
}
