package database

import (
	"DigBGM/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Newdb(conf *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
