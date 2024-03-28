package repos

import (
	"context"
	"encoding/json"
	"errors"
	"log"

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

	for i := range favs {
		var animeData database.AnimeData
		if err := f.db.Where("bangumi_id = ?", favs[i].BangumiID).First(&animeData).Error; err != nil {
			return nil, err
		}
		favs[i].AnimeData = animeData
	}
	return favs, nil
}

// 添加收藏
func (f *favRepository) AddFav(newfav database.Fav) error {
	var existingFav database.Fav
	result := f.db.Where("username = ? AND bangumi_id = ?", newfav.Username, newfav.AnimeData.BangumiID).First(&existingFav)
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
	result := f.db.Where("username = ? AND bangumi_id = ?", username, fav.AnimeData.BangumiID).Delete(&fav)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 运行收藏队列
func (f *favRepository) ProcessFavQueue(rdb *redis.Client) error {
	for {
		// 从Redis队列中获取收藏信息，阻塞等待
		result, err := rdb.BLPop(context.Background(), 0, "fav_queue").Result()
		if err != nil {
			log.Println("Failed to pop fav info from queue:", err)
			continue
		}
		// 提取收藏信息
		favJSON := result[1]

		var favInfo database.Fav
		err = json.Unmarshal([]byte(favJSON), &favInfo)
		if err != nil {
			return err
		}

		// 将收藏信息插入数据库
		err = f.AddFav(favInfo)
		if err != nil {
			return err
		}
		return nil
	}
}
