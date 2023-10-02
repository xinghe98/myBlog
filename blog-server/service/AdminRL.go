package service

// 登录注册等操作

import (
	"fmt"
	"net/http"

	"myBlogServer/v1/dao"
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/models"
	"myBlogServer/v1/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminRL struct{}

// 注册
func (a *AdminRL) Sigup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		httpresp.ResOthers(ctx, 500, util.TransLate(err), "error")
		fmt.Println(err)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.PassWord), 14)
	user.PassWord = string(bytes)

	err = dao.DB.Create(&user).Error
	if err != nil {
		if util.CheckDup(err) {
			httpresp.ResOthers(ctx, http.StatusBadGateway, util.TransLate(err), "该用户名已被使用")
			return
		}

		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}
	httpresp.ResOK(ctx, "注册成功")
}

// 登录并发放token
func (a *AdminRL) Login(ctx *gin.Context) {
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
	ip := ctx.ClientIP()
	token, err := util.GenToken(user.UserName, user.Roles, user.ID, ip)
	if err != nil {
		fmt.Println(err)
		httpresp.ResOthers(ctx, http.StatusBadGateway, err, "服务器错误")
		return
	}
	httpresp.ResOK(ctx, gin.H{"JwtToen": token})
}

func NewAdminService() *AdminRL {
	return &AdminRL{}
}
