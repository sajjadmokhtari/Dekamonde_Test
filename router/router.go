package router

import (
	"dekamonde/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// استاتیک برای فرانت
	r.Static("/", "./frontend")

	// OTP endpoints
	r.POST("/send-otp", handler.SendOtpHandler)
	r.POST("/verify-otp", handler.VerifyOtpHandler)
	return r
}
