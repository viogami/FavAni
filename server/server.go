package server

import (
	"DigBGM/config"
	"DigBGM/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

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

	// 自动迁移模型
	if conf.DB.Migrate {
		db.AutoMigrate(&database.User{})
	}

	//设置Gin的模式
	gin.SetMode(conf.Server.Env)
	// 创建一个 Gin 引擎
	r := gin.Default()

	return &Server{
		engine: r,
		config: conf,
		db:     db,
		rdb:    rdb,
	}, nil
}

type Server struct {
	engine *gin.Engine
	config *config.Config
	db     *gorm.DB
	rdb    *redis.Client
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

func (s *Server) Close() {

}

func (s *Server) initRouter() {
	r := s.engine

	// 设置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 设置路由，用于演示MySQL数据库操作
	r.GET("/createUser", func(c *gin.Context) {
		// 创建用户
		user := database.User{Name: "vio", Age: 20}
		db.Create(&user)

		c.JSON(200, gin.H{
			"message": "User created successfully",
		})
	})

	// 设置路由，用于演示Redis操作
	r.GET("/setRedis", func(c *gin.Context) {
		// 设置Redis键值对
		err := rdb.Set(c, "example_key", "example_value", 0).Err()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to set Redis key",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Redis key set successfully",
		})
	})

	r.GET("/getRedis", func(c *gin.Context) {
		// 获取Redis键值对
		val, err := rdb.Get(c, "example_key").Result()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to get Redis key",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Redis key value: " + val,
		})
	})
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
