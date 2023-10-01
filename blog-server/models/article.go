package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Tags    []*Tags `form:"tags" gorm:"many2many:article_tags;" binding:"required"`
	Title   string  `json:"title," form:"title" gorm:"column:title" binding:"required,max=50,min=1"`
	Status  int8    `json:"status" form:"status" gorm:"column:status;comment:'-1:草稿 1:发布'" binding:"required"`
	Content string  `json:"content" form:"content" gorm:"column:content;type:text" binding:"required"`
}
type Tags struct {
	ID uint `gorm:"primarykey"`
	// select会将其他字段赋值为零值并返回，所以模型结构体需要加上json-tag,避免前端获得无意义的字段，下同
	CreatedAt time.Time      `json:"omitempty"`
	UpdatedAt time.Time      `json:"omitempty"`
	DeletedAt gorm.DeletedAt `json:"omitempty" gorm:"index"`
	Name      string         `json:"name" gorm:"column:name;unique" binding:"required"` // 标签名
	HasArt    []*Article     `gorm:"many2many:article_tags;"`
}
