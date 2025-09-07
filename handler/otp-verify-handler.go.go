package handler

import (
	"dekamonde/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyOTPRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

func VerifyOtpHandler(c *gin.Context) {
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("❌ Invalid request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "درخواست نامعتبر است"})
		return
	}

	log.Println("VerifyOtpHandler received phone:", req.Phone, "otp:", req.OTP)

	if err := services.VerifyOTP(req.Phone, req.OTP); err != nil {
		log.Println("❌ OTP verification failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := services.GenerateJWT(req.Phone,"user")
	if err != nil {
		log.Println("❌ Failed to generate JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
		return
	}
	log.Println("✅ JWT for", req.Phone, ":", token)

	log.Println("✅ OTP verified and JWT generated for phone:", req.Phone)
	c.JSON(http.StatusOK, gin.H{
		"message": "کد OTP با موفقیت تایید شد",
		"token":   token,
	})
}
