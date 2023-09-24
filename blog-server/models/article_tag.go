package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `json:"name" gorm:"column:name;unique" binding:"required"` // 标签名
	Intro string `json:"intro" gorm:"column:desc" binding:"required"`       // 描述
	Color string `json:"color" gorm:"column:color"`                         // 颜色
}
