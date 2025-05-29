package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"thunder_hoster/config"
	"time"
)

const JWT_COOKIE_NAME = "jwt_token"
const GROUP_USER = "user"
const GROUP_ADMIN = "admin"

func InitKeys() {
	if config.Cfg.Secret.SecretKey == "" {
		log.Fatalln("You should fill the secret key")
	}
}

// GenerateJWT 生成 JWT 令牌
func GenerateJWT(userGroup string) (string, error) {
	claims := jwt.MapClaims{
		"group": userGroup,
		"exp":   time.Now().Add(time.Duration(config.Cfg.Service.ValidMin) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.Secret.SecretKey))
}

// VerifyJWT 核对 JWT 令牌
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Cfg.Secret.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
