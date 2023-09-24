package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Tags    []Tag  `json:"tags" form:"tags" gorm:"column:tags;comment:文章标签;foreignkey:Name;"`
	Status  uint8  `json:"status" form:"status" gorm:"type:uint;column:status;comment:文章状态;"`
	Content string `json:"content" form:"content" gorm:"type:text;column:content;comment:文章内容;"`
}
