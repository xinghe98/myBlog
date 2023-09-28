package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Tags    []*Tag `json:"tags" form:"tags" gorm:"many2many:article_tags;" binding:"required"`
	Title   string `json:"title" form:"title" gorm:"column:title" binding:"required"`
	Status  uint   `json:"status" form:"status" gorm:"column:status"`
	Content string `json:"content" form:"content" gorm:"column:content;type:text" binding:"required"`
}
type Tag struct {
	gorm.Model `json:"-"`
	Name       string     `json:"name" gorm:"column:name;unique" binding:"required"` // 标签名
	HasArt     []*Article `gorm:"many2many:article_tags;"`
}
