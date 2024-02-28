package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/viogami/FavAni/database"
	"github.com/viogami/FavAni/middleware"
	"github.com/viogami/FavAni/repos"
)

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

	// get路由，查询用户收藏
	r.GET("getfav", func(c *gin.Context) {
		var user database.User
		// 解析请求体中的JSON数据
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 调用 FavRepository 的GetFav方法
		favs, err := repository.Fav().GetFav(user.Username)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Get fav successfully", "data": favs})
	})

	// post路由，添加用户收藏
	r.POST("addfav", func(c *gin.Context) {
		var fav_info database.Fav
		// 解析请求体中的JSON数据
		if err := c.ShouldBindJSON(&fav_info); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 调用 FavRepository 的AddFav方法
		err := repository.Fav().AddFav(fav_info.Username, fav_info)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Add fav successfully", "data": fav_info})
	})

	// post路由，删除用户收藏
	r.POST("delfav", func(c *gin.Context) {
		var fav_info database.Fav
		// 解析请求体中的JSON数据
		if err := c.ShouldBindJSON(&fav_info); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 调用 FavRepository 的DeleteFav方法
		err := repository.Fav().DeleteFav(fav_info.Username, fav_info)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Delete fav successfully", "data": fav_info})
	})
}
