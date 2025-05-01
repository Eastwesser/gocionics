package middleware

import (
	"github.com/gin-gonic/gin"
	"gocionics/internal/usecases/auth"
)

func AuthMiddleware(authUC *auth.AuthUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		// Проверка токена
		user, err := authUC.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
