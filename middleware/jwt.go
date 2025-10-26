package middleware

import (
	"bluebell/jwt"
	"bluebell/pkg/code"
	"bluebell/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			zap.S().Info("no authorization header")
			response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.JSON(c, http.StatusUnauthorized, code.TokenMalformed, nil)
			c.Abort()
			return
		}
		tokenString := parts[1]

		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.JSON(c, http.StatusUnauthorized, code.InvalidToken, nil)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
