package services

import (
	"dekamonde/data/db"
	"dekamonde/data/models"
	"fmt"

	"gorm.io/gorm"
)

func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	if err := db.GetDb().Where("phone = ?", phone).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func ListUsers(search string, page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := db.GetDb().Model(&models.User{})
	if search != "" {
		query = query.Where("phone ILIKE ?", fmt.Sprintf("%%%s%%", search))
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
