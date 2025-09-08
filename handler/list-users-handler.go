package handler

import (
	"dekamonde/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListUsersHandler godoc
// @Summary دریافت لیست کاربران
// @Description دریافت لیست کاربران با قابلیت جستجو و صفحه‌بندی
// @Tags Users
// @Accept json
// @Produce json
// @Param search query string false "جستجو بر اساس شماره تلفن یا سایر فیلدها"
// @Param page query int false "شماره صفحه"
// @Param limit query int false "تعداد نتایج در هر صفحه"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{} "خطا در دریافت لیست کاربران"
// @Router /users [get]
func ListUsersHandler(c *gin.Context) {
	search := c.Query("search")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	users, total, err := services.ListUsers(search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت لیست کاربران"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
