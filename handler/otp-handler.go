package handler

import (
	"dekamonde/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhoneRequest struct {
	Phone string `json:"phone"`
}

type Response struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message,omitempty"`
}

// SendOtpHandler godoc
// @Summary ارسال کد OTP
// @Description این متد شماره موبایل کاربر را می‌گیرد و کد OTP تولید می‌کند
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body PhoneRequest true "شماره موبایل"
// @Success 200 {object} Response "کد OTP ارسال شد"
// @Failure 400 {object} Response "درخواست نامعتبر"
// @Failure 500 {object} Response "خطا در ارسال OTP"
// @Router /send-otp [post]
func SendOtpHandler(c *gin.Context) {
	var req PhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error decoding phone request:", err)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	fmt.Println("SendOtpHandler received phone:", req.Phone)

	// بررسی شماره موبایل ایرانی
	if err := services.IsIranianPhone(req.Phone); err != nil {
		fmt.Println("Invalid phone number:", req.Phone)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: err.Error()})
		return
	}

	if err := services.SendOTP(req.Phone); err != nil {
		fmt.Println("Error sending OTP:", err, "Phone:", req.Phone)
		c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ارسال OTP"})
		return
	}

	fmt.Println("OTP sent successfully to:", req.Phone)
	c.JSON(http.StatusOK, Response{Valid: true, Message: "کد OTP ارسال شد"})
}
