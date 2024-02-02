package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/viogami/FavAni/database"
)

const (
	Issuer = "favani"
)

type CustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type JWTService struct {
	signKey        []byte
	issuer         string
	expireDuration time.Duration
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		signKey:        []byte(secret),
		issuer:         Issuer,
		expireDuration: 5 * 24 * time.Hour,
	}
}

func (s *JWTService) CreateToken(user *database.User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("empty user")
	}
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			Name: user.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(now.Add(s.expireDuration)),
				NotBefore: jwt.NewNumericDate(now.Add(-1000 * time.Second)),
				Issuer:    s.issuer,
			},
		},
	)

	return token.SignedString(s.signKey)
}

func (s *JWTService) ParseToken(tokenString string) (*database.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return s.signKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invaild token")
	}

	user := &database.User{
		Username: claims.Name,
	}

	return user, nil
}
