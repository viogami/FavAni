package middleware

import (
	"FavAni/auth"
	"FavAni/repos"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(jwtService *auth.JWTService, userRepo repos.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := getTokenFromAuthorizationHeader(c)
		if token == "" {
			token, _ = getTokenFromCookie(c)
		}

		user, _ := jwtService.ParseToken(token)
		if user != nil {
			user, err := userRepo.GetUserByName(user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "GetUserByName failed!"})
				return
			}
			c.Set("user", user)
		}

		c.Next()
	}
}

func getTokenFromCookie(c *gin.Context) (string, error) {
	return c.Cookie("token")
}

func getTokenFromAuthorizationHeader(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", nil
	}

	token := strings.Fields(auth)
	if len(token) != 2 || strings.ToLower(token[0]) != "bearer" || token[1] == "" {
		return "", fmt.Errorf("authorization header invaild")
	}

	return token[1], nil
}
