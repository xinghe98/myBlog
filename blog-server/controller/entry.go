package controller

import "myBlogServer/v1/service"

var (
	AdminResAndLogin = NewAdmin(service.NewAdminService())
	ArticleCRUD      = NewArticleCO(service.NewArticle())
	TagsCRUD         = NewTagCo(service.NewTags())
)
