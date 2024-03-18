package database

import "gorm.io/gorm"

// 定义Fav表
type Fav struct {
	gorm.Model
	Username  string    `json:"username" gorm:"size:100;not null;"`
	AnimeID   int       `json:"anime_id" gorm:"not null;"` // 外键
	AnimeData AnimeData `json:"data_anime" gorm:"ForeignKey:AnimeID;"`
}

type AnimeData struct {
	AnimeID     int    `json:"anime_id" gorm:"size:100;not null"`
	AnimeName   string `json:"anime" gorm:"size:100;not null;"`
	AnimeNameCN string `json:"anime_cn" gorm:"size:100;not null;"`
	// Tags        []Tags `json:"tags" gorm:"size:100;not null;"`
}

type Tags struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:100;not null;"`
	Count int    `json:"count" gorm:"size:100;not null;"`
}

// Fav表的外键
func (Fav) TableName() string {
	return "favs"
}

// Fav表的外键
func (AnimeData) TableName() string {
	return "anime_data"
}

// 收藏列表
type Favs []Fav
