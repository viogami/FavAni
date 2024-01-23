package database

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LogoutRequest 登出请求结构体
type LogoutRequest struct {
	Token string `json:"token" binding:"required"`
}
