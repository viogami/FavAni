package repos

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/viogami/FavAni/database"
	"gorm.io/gorm"
)

type favRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewFavRepository(db *gorm.DB, rdb *redis.Client) FavRepository {
	return &favRepository{
		db:  db,
		rdb: rdb,
	}
}

// 获取用户收藏
func (f *favRepository) GetFav(username string) (database.Favs, error) {
	var favs database.Favs
	result := f.db.Where("username = ?", username).Find(&favs)
	if result.Error != nil {
		return nil, result.Error
	}
	return favs, nil
}

// 添加收藏
func (f *favRepository) AddFav(newfav database.Fav) error {
	var existingFav database.Fav
	result := f.db.Where("username = ? AND anime_id = ?", newfav.Username, newfav.AnimeData.AnimeID).First(&existingFav)
	if result.Error == nil {
		// 收藏已存在，返回错误
		return errors.New("the adding fav already exists")
	}
	result = f.db.Create(&newfav)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 删除收藏
func (f *favRepository) DeleteFav(username string, fav database.Fav) error {
	result := f.db.Where("username = ? AND anime_id = ?", username, fav.AnimeData.AnimeID).Delete(&fav)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
