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

// VerifyOtpHandler godoc
// @Summary تایید کد OTP
// @Description این متد شماره موبایل و کد OTP را گرفته و در صورت صحت، JWT تولید می‌کند
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body VerifyOTPRequest true "شماره موبایل و OTP"
// @Success 200 {object} map[string]interface{} "JWT تولید شد"
// @Failure 400 {object} map[string]interface{} "درخواست نامعتبر"
// @Failure 401 {object} map[string]interface{} "OTP نادرست یا منقضی شده"
// @Failure 500 {object} map[string]interface{} "خطا در تولید توکن"
// @Router /verify-otp [post]
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
	token, err := services.GenerateJWT(req.Phone, "user")
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
