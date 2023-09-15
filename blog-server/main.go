package main

import (
	"myBlogServer/v1/dao"
	"myBlogServer/v1/middleware"
	"myBlogServer/v1/router"
	"myBlogServer/v1/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	util.InitConfig()
	// 连接数据库
	dao.ConnectMysql()
	// 初始化数据模型
	dao.InitModel()
	// 初始化全局错误翻译（针对校验库）
	util.InitTrans()
	// 初始化gin
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 注册中间件
	r.Use(middleware.Cors())
	// 注册路由
	r = router.RegistRouters(r)
	// 启用
	r.Run(":3001")
}
