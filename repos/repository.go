package repos

import (
	"github.com/viogami/FavAni/database"
	"gorm.io/gorm"
)

type repository struct {
	user UserRepository
	fav  FavRepository
	db   *gorm.DB
	rdb  *database.RedisDB
}

func NewRepository(db *gorm.DB, rdb *database.RedisDB) Repository {
	r := &repository{
		db:   db,
		rdb:  rdb,
		user: NewUserRepository(db, rdb),
		fav:  NewFavRepository(db, rdb),
	}
	return r
}

func (r *repository) User() UserRepository {
	return r.user
}

func (r *repository) Fav() FavRepository {
	return r.fav
}
