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
// @Description این متد شماره موبایل و کد OTP را گرفته و در صورت صحت، JWT تولید می‌کند و در کوکی ذخیره می‌کند
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body VerifyOTPRequest true "شماره موبایل و OTP"
// @Success 200 {object} map[string]interface{} "JWT تولید شد و در کوکی ذخیره شد"
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

	// تغییر اصلی: دریافت User از VerifyOTP
	user, err := services.VerifyOTP(req.Phone, req.OTP)
	if err != nil {
		log.Println("❌ OTP verification failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token برای کاربر موجود یا جدید
	token, err := services.GenerateJWT(user.Phone, "user")
	if err != nil {
		log.Println("❌ Failed to generate JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	// ست کردن توکن در کوکی امن و HttpOnly
	c.SetCookie(
		"token",  // نام کوکی
		token,    // مقدار توکن JWT
		3600*24,  // انقضا 24 ساعت به ثانیه
		"/",      // مسیر کوکی
		"",       // دامنه، خالی برای localhost
		false,    // Secure (برای HTTPS باید true باشد)
		true,     // HttpOnly
	)

	log.Println("✅ JWT generated and stored in cookie for", user.Phone)
	c.JSON(http.StatusOK, gin.H{
		"message": "کد OTP با موفقیت تایید شد و توکن در کوکی ذخیره شد",
		"token":   token,
	})
}
