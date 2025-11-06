package middleware

import (
	"net/http"
	"strings"
	"test/repository"
	"test/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth(jwtService *service.JWTService, userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не предоставлен"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена. Используйте: Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен: " + err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

		user, err := userRepo.FindByID(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
