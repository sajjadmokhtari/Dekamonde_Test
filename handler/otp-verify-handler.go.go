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

	log.Println("✅ OTP verified successfully for phone:", req.Phone)
	c.JSON(http.StatusOK, gin.H{"message": "کد OTP با موفقیت تایید شد"})
}
