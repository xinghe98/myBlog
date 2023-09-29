package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Tags    []*Tag `json:"tags" form:"tags" gorm:"many2many:article_tags;" binding:"required"`
	Title   string `json:"title" form:"title" gorm:"column:title" binding:"required,max=50,min=1"`
	Status  int8   `json:"status" form:"status" gorm:"column:status;comment:'-1:草稿 1:发布'" binding:"required"`
	Content string `json:"content" form:"content" gorm:"column:content;type:text" binding:"required"`
}
type Tag struct {
	gorm.Model `json:"-"`
	Name       string     `json:"name" gorm:"column:name;unique" binding:"required"` // 标签名
	HasArt     []*Article `gorm:"many2many:article_tags;"`
}
