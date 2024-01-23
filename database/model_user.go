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

// 用户列表
type Users []User
