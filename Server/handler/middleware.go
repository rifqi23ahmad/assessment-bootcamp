package handler

import (
	"assessment-bootcamp/Server/auth"
	"assessment-bootcamp/Server/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(userService user.UserService, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorize user",
			})
			return
		}

		token, err := authService.ValidateToken(authHeader)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorize user",
			})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unauthorize user",
			})
			return
		}

		userID := int(claim["user_id"].(float64))

		c.Set("currentUser", userID)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
