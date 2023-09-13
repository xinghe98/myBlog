package service

import (
	"encoding/json"
	"fmt"
	"myBlogServer/v1/dao"
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/models"
	"myBlogServer/v1/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Sigup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		httpresp.ResOthers(ctx, 500, util.TransLate(err), "error")
		fmt.Println(err)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.PassWord), 14)
	user.PassWord = string(bytes)

	// TODO:下面可以进行gorm错误的封装
	type GormErr struct {
		Number  int    `json:"Number"`
		Message string `json:"Message"`
	}

	if err := dao.DB.Create(&user).Error; err != nil {
		byteErr, _ := json.Marshal(err)
		var newError GormErr
		json.Unmarshal((byteErr), &newError)
		switch newError.Number {
		case 1062:
			fmt.Println("Duplicate Key !")
			httpresp.ResOthers(ctx, http.StatusBadGateway, util.TransLate(gorm.ErrDuplicatedKey), "该用户名已被使用")
		}
		return
	}
	httpresp.ResOK(ctx, "注册成功")
}
