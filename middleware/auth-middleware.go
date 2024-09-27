package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"task-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
            c.Abort()
            return
        }
		claims := &utils.Claims{}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		SecretKeyString := os.Getenv("SECRET_KEY")
		if SecretKeyString == "" {
			fmt.Println("Secret key not found")
		}
		token,err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(SecretKeyString), nil
        })
		
		
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("userId", claims.UserID) // Store userId in the context for later use
        c.Next()
	}
}