package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Roles    string `json:"-" gorm:"column:roles;default:user;"`
	UserName string `json:"username" gorm:"column:username;unique" binding:"required"`
	PassWord string `json:"password" gorm:"column:password;not null" binding:"required"`
}
