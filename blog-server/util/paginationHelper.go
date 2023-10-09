package util

import (
	"fmt"

	"myBlogServer/v1/models"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) (pagination models.Pageination) {
	if err := c.ShouldBind(&pagination); err != nil {
		fmt.Printf("参数绑定错误:%s\n", err)
	}

	// 校验参数
	if pagination.Limit == 0 && pagination.Page == 0 && len(pagination.Sort) == 0 && pagination.Status == 0 {
		return
	}

	if pagination.Limit == 0 && pagination.Page == 0 && len(pagination.Sort) == 0 && pagination.Status != 0 {
		return
	}
	if pagination.Limit < 0 {
		pagination.Limit = 2
	}
	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if len(pagination.Sort) == 0 {
		pagination.Sort = "created_at desc"
	}
	if pagination.Status < -1 || pagination.Status > 1 {
		pagination.Status = 1
	}

	return
}
