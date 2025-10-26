package jwt

import (
	"bluebell/global"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

const (
	TokenExpireDuration = 24 * time.Hour
)

var jwtSecret = []byte("MySecretKey")

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成带过期时间的 Token
func GenerateToken(userID int, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
			Issuer:    "bluebell",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	userIDStr := strconv.Itoa(userID)
	ctx := context.Background()
	userKey := fmt.Sprintf("bluebell:user:%s", userIDStr)
	err = global.Redis.Set(ctx, userKey, token, time.Duration(TokenExpireDuration)).Err()

	if err != nil {
		zap.S().Error("failed to set token to redis", err)
		return "", err
	}

	return token, nil
}

// ParseToken 解析并验证 Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 检查 token 是否有效，并获取载荷信息
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 二次确认：检查该 token 是否存在于 Redis 中
		if !IsTokenValidInRedis(claims.UserID, tokenString) {
			return nil, fmt.Errorf("token 已失效或被替换")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func IsTokenValidInRedis(userID int, tokenString string) bool {
	userIDStr := strconv.Itoa(userID)
	ctx := context.Background()
	userKey := fmt.Sprintf("bluebell:user:%s", userIDStr)
	redisToken, err := global.Redis.Get(ctx, userKey).Result()
	if err != nil {
		zap.S().Error("failed to get token from redis", err)
		return false
	}
	return redisToken == tokenString
}
