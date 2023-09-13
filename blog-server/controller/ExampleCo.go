package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleCo struct{}

func (e ExampleCo) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}
