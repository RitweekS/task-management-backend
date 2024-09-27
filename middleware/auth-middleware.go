package middleware

import (
	"net/http"
	"strings"
	"task-management/utils"

	"github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)


func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        // Extract the token from the "Authorization" header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
            c.Abort()
            return
        }
		claims := &utils.Claims{}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token,err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return utils.SecretKey, nil
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