package repos

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type repository struct {
	user UserRepository
	db   *gorm.DB
	rdb  *redis.Client
}

func NewRepository(db *gorm.DB, rdb *redis.Client) Repository {
	r := &repository{
		db:   db,
		rdb:  rdb,
		user: NewUserRepository(db, rdb),
	}
	return r
}

func (r *repository) User() UserRepository {
	return r.user
}
