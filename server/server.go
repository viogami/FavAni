package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/viogami/FavAni/auth"
	"github.com/viogami/FavAni/config"
	"github.com/viogami/FavAni/database"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

func New(conf *config.Config) (*Server, error) {
	if conf != nil {
		var err error
		db, err = database.Newdb(&conf.DB)
		if err != nil {
			return nil, errors.Wrap(err, "db init failed")
		}

		rdb, err = database.NewRedisClient(&conf.Redis)
		if err != nil {
			return nil, errors.Wrap(err, "redis client failed")
		}
	} else {
		log.Println("Error:config file is nil")
	}

	// 自动迁移模型（初始创建表）
	if conf.DB.Migrate {
		db.AutoMigrate(&database.User{}, &database.Fav{}, &database.AnimeData{}, &database.Tags{})
	}

	// jwt
	jwtService := auth.NewJWTService(conf.Server.JWTSecret)

	//设置Gin的模式
	gin.SetMode(conf.Server.Env)
	// 创建一个 Gin 引擎
	r := gin.Default()

	return &Server{
		engine:     r,
		config:     conf,
		db:         db,
		rdb:        rdb,
		jwtService: jwtService,
	}, nil
}

type Server struct {
	engine     *gin.Engine
	config     *config.Config
	db         *gorm.DB
	rdb        *redis.Client
	jwtService *auth.JWTService
}

func (s *Server) Run() error {
	s.initRouter()

	//读取服务器地址
	addr := fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port)

	// 启动Gin服务器
	err := s.engine.Run(addr)
	if err != nil {
		fmt.Println("Failed to start Gin server")
	}
	return err
}

func (s *Server) getRoutes() []string {
	paths := []string{}
	for _, r := range s.engine.Routes() {
		if r.Path != "" {
			paths = append(paths, r.Path)
		}
	}
	return paths
}
