package handler

import (
	"dekamonde/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserHandler godoc
// @Summary دریافت جزئیات یک کاربر
// @Description با استفاده از شماره تلفن، جزئیات کاربر را دریافت می‌کند
// @Tags Users
// @Accept json
// @Produce json
// @Param phone path string true "شماره تلفن کاربر"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{} "کاربر یافت نشد"
// @Failure 500 {object} map[string]interface{} "خطا در دریافت کاربر"
// @Router /user/{phone} [get]
func GetUserHandler(c *gin.Context) {
	phone := c.Param("phone")
	user, err := services.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت کاربر"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
		return
	}
	c.JSON(http.StatusOK, user)
}
