package middleware

import (
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 验证前端传来的jwttoken
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstr := ctx.GetHeader("Authorization")
		if tokenstr == "" || !strings.HasPrefix(tokenstr, "Bearer") {
			httpresp.ResOthers(ctx, http.StatusUnauthorized, nil, "请先登录")
			ctx.Abort()
			return
		}
		tokenstr = tokenstr[7:]
		token, claims, err := util.ParseToken(tokenstr)
		if err != nil || !token.Valid {
			httpresp.ResOthers(ctx, http.StatusUnauthorized, nil, "不合法的token")
			ctx.Abort()
			return
		}
		roles := claims.Roles
		if roles == "" {
			httpresp.ResOthers(ctx, http.StatusLocked, nil, "无权限")
			ctx.Abort()
			return
		}
		ctx.Set("who", roles)
	}
}
