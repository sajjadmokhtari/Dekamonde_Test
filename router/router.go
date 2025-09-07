package router

import (
	"dekamonde/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Swagger - قبل از مسیر استاتیک
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// OTP endpoints
	r.POST("/send-otp", handler.SendOtpHandler)
	r.POST("/verify-otp", handler.VerifyOtpHandler)

	// استاتیک برای فرانت‌اند
	r.Static("/front", "./frontend")

	return r
}
