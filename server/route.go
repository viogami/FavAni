package server

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/viogami/FavAni/database"
	"github.com/viogami/FavAni/middleware"
	pb "github.com/viogami/FavAni/pb/gcn"
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
	r.GET("/routes", func(c *gin.Context) {
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
			c.JSON(401, gin.H{"error": "Username or Password is Invalid,please retry"})
			return
		}
		// 创建token
		token, err := s.jwtService.CreateToken(&database.User{Username: loginRequest.Username})
		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}
		c.JSON(200, gin.H{"message": "Login successful", "user": user.Username, "token": token})
	})

	// auth/logout POST路由，用于注销用户
	auth.POST("/logout", func(c *gin.Context) {
		// 调用 Repository_user 的注销方法
		err := userRepository.Logout()
		user := c.MustGet("user").(*database.User)
		if err != nil {
			c.JSON(401, gin.H{"error": "logout failed " + err.Error()})
			return
		}
		if user == nil {
			c.JSON(401, gin.H{"error": "jwt auth failed , user not found"})
			return
		}
		c.JSON(200, gin.H{"message": user.Username + " Logout successfully"})
	})

	// auth/jwt POST路由，用于验证token
	auth.POST("/jwt", func(c *gin.Context) {
		user := c.MustGet("user").(*database.User) // 无需做panic检查，因为中间件已经检查过了
		c.JSON(200, gin.H{"message": "jwt is valid", "user": user.Username})
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

	// get路由，用于获取全部用户
	r.GET("/user_list", func(c *gin.Context) {
		// 调用 UserRepository 的List方法
		users, err := userRepository.List()
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Get all user successfully", "data": users})
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
	r.GET("/getfav/:username", func(c *gin.Context) {
		// 获取URL参数
		username := c.Param("username")
		// 调用 FavRepository 的GetFav方法
		favs, err := repository.Fav().GetFav(username)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Get fav successfully", "data": favs})
	})

	// post路由，添加用户收藏
	r.POST("/addfav", func(c *gin.Context) {
		// 直接获取请求体中的JSON字符串
		favJSON, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = rdb.LPush(context.Background(), "fav_queue", favJSON).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to push fav info to queue." + err.Error()})
			return
		}

		// 启动后台处理程序
		err = repository.Fav().ProcessFavQueue(rdb)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to process fav queue." + err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Add fav request received", "data": string(favJSON)})
	})

	// post路由，删除用户收藏
	r.POST("/delfav", func(c *gin.Context) {
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

	// get路由，grpc请求
	r.GET("/gcn", func(c *gin.Context) {
		// 读取节点和边的JSON文件
		nodesData, _ := os.ReadFile("asserts/GraphData/nodes.json")
		edgesData, _ := os.ReadFile("asserts/GraphData/edges.json")

		// 解析JSON数据
		var nodes []*pb.Node
		if err := json.Unmarshal(nodesData, &nodes); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		var edges []*pb.Edge
		if err := json.Unmarshal(edgesData, &edges); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		// 创建一个实例的GraphData请求
		G_example := &pb.GraphData{
			Nodes: nodes,
			Edges: edges,
		}
		params := &pb.ModelParams{
			InputDims:  1,
			HiddenDims: 8,
			OutputDims: 1,
		}
		r_gcn := &pb.GCNRequest{
			Graph:  G_example,
			Params: params,
		}
		// 调用 grpc 的请求方法
		res, err := GCN_request(s.gRPCConn, r_gcn)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// 返回成功响应
		c.JSON(200, gin.H{"message": "Get gcn_result successfully , this is a grpc server(python)", "data": res})
	})

}
