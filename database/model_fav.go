package database

import "gorm.io/gorm"

// 定义Fav表
type Fav struct {
	gorm.Model
	Username  string    `json:"username" gorm:"size:100;not null;"`
	BangumiID int       `json:"bangumi_id" gorm:"size:100;not null;"`
	AnimeData AnimeData `json:"data_anime" gorm:"foreignKey:ID"`
}

type AnimeData struct {
	gorm.Model
	BangumiID   int    `json:"bangumi_id" gorm:"size:100;not null;"`
	AnimeName   string `json:"anime" gorm:"size:100;not null;"`
	AnimeNameCN string `json:"anime_cn" gorm:"size:100;not null;"`
	Tags        []Tags `json:"tags" gorm:"many2many:animedata_tags;"`
}

type Tags struct {
	gorm.Model
	Name      string      `json:"name" gorm:"size:100;not null;"`
	Count     int         `json:"count" gorm:"size:100;not null;"`
	AnimeData []AnimeData `json:"anime_data" gorm:"many2many:animedata_tags;"`
}

// Fav表名
func (Fav) TableName() string {
	return "favs"
}

// AnimeData表名
func (AnimeData) TableName() string {
	return "anime_data"
}

// 收藏列表
type Favs []Fav
