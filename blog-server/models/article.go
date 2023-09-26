package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Tags    []*Tag `json:"tags" form:"tags" gorm:"many2many:article_tags;" `
	Status  uint   `json:"status" form:"status" gorm:"column:status"`
	Content string `json:"content" form:"content" gorm:"column:content;type:text" binding:"required"`
}
type Tag struct {
	gorm.Model
	Name   string     `json:"name" gorm:"column:name;unique" binding:"required"` // 标签名
	Intro  string     `json:"intro" gorm:"column:desc" binding:"required"`       // 描述
	Color  string     `json:"color" gorm:"column:color"`                         // 颜色
	HasArt []*Article `gorm:"many2many:article_tags;"`
}
