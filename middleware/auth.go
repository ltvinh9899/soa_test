package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Lấy token từ header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// 2. Tách bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		// 3. Kiểm tra token hợp lệ
		if err != nil || !token.Valid {
            log.Println("Invalid token:", token.Valid)
            log.Println("Token parsing error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// 4. Lưu thông tin từ token vào context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", uint(claims["sub"].(float64))) // Ép kiểu từ JWT number sang uint
			c.Set("userRole", claims["role"])
		}

		c.Next()
	}
}

func AdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Admin access required",
			})
			return
		}
		c.Next()
	}
}