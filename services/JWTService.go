package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JWTService struct {
	secretKey  string
	iss        string
	expireTime int64
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey:  viper.GetString("SECRET_KEY"),
		iss:        viper.GetString("ISSUER"),
		expireTime: viper.GetInt64("JWT_EXPIRE_TIME"),
	}
}

func (j *JWTService) GenerateToken(username string) (string, error) {
	expireToken := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.expireTime)))

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireToken,
			Issuer:    j.iss,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}
