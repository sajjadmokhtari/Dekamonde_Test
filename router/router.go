package router

import (
	"dekamonde/handler"
	middlewares "dekamonde/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// OTP endpoints - بدون JWT
	r.POST("/send-otp", handler.SendOtpHandler)
	r.POST("/verify-otp", handler.VerifyOtpHandler)

	// گروهی از مسیرهای محافظت‌شده با JWT
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/user/:phone", handler.GetUserHandler)
		auth.GET("/users", handler.ListUsersHandler)
	}

	// استاتیک برای فرانت‌اند
	r.Static("/front", "/app/frontend")

	return r
}
