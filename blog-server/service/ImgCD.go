package service

import (
	"log"
	"net/http"

	"myBlogServer/v1/dao"
	"myBlogServer/v1/httpresp"
	"myBlogServer/v1/models"
	"myBlogServer/v1/util"

	"github.com/gin-gonic/gin"
)

type ImgCD struct{}

func NewImgCD() *ImgCD {
	return &ImgCD{}
}

// UploadImg 上传图片
// @Summary 上传图片
// @Description 上传图片
// @Tags 上传图片
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {string} json "{"code":200,"data":{file_cos_uuid},"msg":"ok"}"
// @Router /api/v1/uploadImg [post]
func (i ImgCD) UploadImg(ctx *gin.Context) {
	// 上传图片
	file, err := ctx.FormFile("file")
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "上传失败，获取文件失败！")
		return
	}
	f, err := file.Open()
	if err != nil {
		log.Println(err)
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "上传失败，打开文件失败！")
	}
	defer f.Close()
	// 上传文件到腾讯云
	filename := util.UploadImg(f)
	httpresp.ResOK(ctx, gin.H{"file_cos_uuid": filename})
}

// DeleteImg 删除图片
// @Summary 删除图片
// @Description 删除图片
func (i ImgCD) DeleteImg(ctx *gin.Context) {
	// 删除图片
	filename := ctx.Param("key")
	ok := util.DeleteImg(filename)
	if !ok {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "删除失败！")
		return
	}

	httpresp.ResOK(ctx, gin.H{"msg": "删除成功！", "success": true})
}

// 查询所有有headlineimg的文章
func (i ImgCD) ReadAllHeadImgArt(ctx *gin.Context) {
	// 查询所有有headlineimg的文章
	var articles []models.Article
	err := dao.DB.Where("head_img != ?", "").Find(&articles).Error
	if err != nil {
		httpresp.ResOthers(ctx, http.StatusBadGateway, nil, "服务器错误")
		return
	}
	httpresp.ResOK(ctx, gin.H{"articles": articles})
}
