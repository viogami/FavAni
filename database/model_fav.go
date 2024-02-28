package database

import "gorm.io/gorm"

// 定义Fav表
type Fav struct {
	gorm.Model
	Username string `json:"username" gorm:"size:100;not null;"`
	Anime    string `json:"anime" gorm:"size:100;not null;"`
	AnimeID  string `json:"anime_id" gorm:"size:100;not null;"`
	Tag      string `json:"tag" gorm:"size:100;"`
}

// 用户列表
type Favs []Fav
