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

// 登录并发放token
func Login(ctx *gin.Context) {
	var postuser models.User
	ctx.BindJSON(&postuser)
	var user models.User
	dao.DB.Where("username=?", postuser.UserName).First(&user)
	if user.UserName == "" {
		httpresp.ResOthers(ctx, http.StatusUnprocessableEntity, nil, "用户不存在")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(postuser.PassWord))
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadRequest, nil, "密码错误")
		return
	}
	// 给个token
	fmt.Println(user.Roles)
	token, err := util.GenToken(user.UserName, user.Roles, user.ID)
	if err != nil {
		fmt.Println(err)
		httpresp.ResOthers(ctx, http.StatusBadGateway, err, "服务器错误")
		return
	}
	httpresp.ResOK(ctx, gin.H{"JwtToen": token})
}
