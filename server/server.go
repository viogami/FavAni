package server

import (
	"FavAni/auth"
	"FavAni/config"
	"FavAni/database"
	"FavAni/middleware"
	"FavAni/repos"
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

func (s *Server) Close() {

}

// 创建路由，路由初始化
func (s *Server) initRouter() {
	r := s.engine

	// 初始化数据库
	repository := repos.NewRepository(db, rdb)
	userRepository := repos.NewUserRepository(db, rdb)

	// 路由分组
	auth := r.Group("/auth").Use(
		middleware.JwtAcMiddleware(s.jwtService, repository.User()),
	)
	// 设置根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, this is favani`s backend!",
		})
	})
	// 获取全部路由
	r.GET("/route", func(c *gin.Context) {
		path := s.getRoutes()
		c.JSON(200, gin.H{
			"routes": path,
		})
	})

	// login POST路由，用于用户登陆
	r.POST("/login", func(c *gin.Context) {
		var loginRequest database.LoginRequest
		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		// 调用 Repository_user 的登录方法
		user, err := userRepository.Login(loginRequest.Username, loginRequest.Password)
		if err != nil {
			log.Println(err)
			c.JSON(401, gin.H{"error": "Username or Password is Invalid,please retry"})
			return
		}
		// 创建token
		token, err := s.jwtService.CreateToken(&database.User{Username: loginRequest.Username})
		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}
		c.JSON(200, gin.H{"message": "Login successful", "user": user, "token": token})
	})

	// auth/logout POST路由，用于注销用户
	auth.POST("/logout", func(c *gin.Context) {
		// 调用 Repository_user 的注销方法
		err := userRepository.Logout()
		if err != nil {
			log.Println(err)
			c.JSON(401, gin.H{"error": "logout failed"})
			return
		}
		c.JSON(200, gin.H{"message": "Logout successful"})
	})

	// POST路由，用于用户注册
	r.POST("/register", func(c *gin.Context) {
		var newUser database.User
		// 解析请求体中的JSON数据
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 调用 UserRepository 的注册方法
		err := userRepository.Register(newUser)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "registered successfully", "data": newUser.Username})
	})

	// post路由,用于删除用户
	r.POST("/deluser", func(c *gin.Context) {
		var delUser database.User
		// 解析请求体中的JSON数据
		if err := c.ShouldBindJSON(&delUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 调用 UserRepository 的Delete方法
		err := userRepository.Delete(delUser)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Delete successfully:" + delUser.Username})
	})

	// get路由，用于演示Redis操作
	// r.GET("/setRedis", api.RedisSet)

	// r.GET("/getRedis", api.RedisHandel)
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
