package database

import "gorm.io/gorm"

// 定义user表

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:100;not null;unique"`
	Password string `json:"password" gorm:"size:256;"`
	Email    string `json:"email" gorm:"size:256;"`
	Avatar   string `json:"avatar" gorm:"size:256;"`
}

type Users []User

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}