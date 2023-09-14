package controller

import (
	"myBlogServer/v1/service"

	"github.com/gin-gonic/gin"
)

type AdminCo struct{}

func (a AdminCo) Regist(ctx *gin.Context) {
	service.Sigup(ctx)
}

func (a AdminCo) Login(ctx *gin.Context) {
	service.Login(ctx)
}
