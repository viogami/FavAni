package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/viogami/FavAni/config"
)

var (
	ErrRedisDisable = errors.New("redis disable")
)

type RedisDB struct {
	enable bool
	*redis.Client
}

func NewRedisClient(conf *config.RedisConfig) (*RedisDB, error) {
	if !conf.Enable {
		return &RedisDB{}, nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisDB{
		enable: true,
		Client: rdb,
	}, nil
}

func (rdb *RedisDB) Endable() bool {
	return rdb.enable
}

func (rdb *RedisDB) HGet(key, field string, obj interface{}) error {
	if !rdb.enable {
		return ErrRedisDisable
	}

	return rdb.Client.HGet(context.Background(), key, field).Scan(obj)
}

func (rdb *RedisDB) HSet(key, field string, val interface{}) error {
	if !rdb.enable {
		return nil
	}

	return rdb.Client.HSet(context.Background(), key, field, val).Err()
}

func (rdb *RedisDB) HDel(key string, fields ...string) error {
	if !rdb.enable {
		return nil
	}

	return rdb.Client.HDel(context.Background(), key, fields...).Err()
}

// 消息队列
func (rdb *RedisDB) LPush(key string, values ...interface{}) error {
	if !rdb.enable {
		return ErrRedisDisable
	}

	return rdb.Client.LPush(context.Background(), key, values...).Err()
}

func (rdb *RedisDB) BLPop(key string) ([]string, error) {
	if !rdb.enable {
		return nil, ErrRedisDisable
	}
	result, err := rdb.Client.BLPop(context.Background(), 0, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
