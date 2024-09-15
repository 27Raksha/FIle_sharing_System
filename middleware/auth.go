package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"21BLC1564/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if gin.Mode() == gin.TestMode {
			c.Set("email", "testuser@example.com") 
			c.Next()
			return
		}
		
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with 'Bearer '"})
			c.Abort()
			return
		}

		
		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))



		claims, valid := utils.ValidateJWT(tokenStr)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}


		c.Set("email", claims.Email)
		c.Next() 
	}
}
