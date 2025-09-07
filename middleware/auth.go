package middlewares

import (
	"dekamonde/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware بررسی JWT و قرار دادن اطلاعات کاربر در کانتکست
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// دریافت توکن از کوکی
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن یافت نشد"})
			c.Abort()
			return
		}

		// اعتبارسنجی توکن
		claims, err := services.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن نامعتبر یا منقضی شده است"})
			c.Abort()
			return
		}

		// ذخیره اطلاعات کاربر در کانتکست Gin
		c.Set("userPhone", claims.Phone)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}
