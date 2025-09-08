package services

import (
	"errors"
	"log"
	"strings"
	"time"

	"dekamonde/cache"
	"dekamonde/data/db"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Phone     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}

// VerifyOTP بررسی OTP و ایجاد کاربر جدید در صورت عدم وجود
func VerifyOTP(phone, otp string) (*User, error) {
	storedOtp, err := cache.GetOTP(phone)
	if err != nil {
		log.Print("❌ OTP not found or expired:", phone, err)
		return nil, errors.New("کد OTP منقضی شده یا موجود نیست")
	}

	storedOtp = strings.TrimSpace(storedOtp)
	otp = strings.TrimSpace(otp)

	log.Println("🔹 Stored OTP:", storedOtp)
	log.Println("🔹 Provided OTP:", otp)

	if storedOtp != otp {
		log.Print("❌ OTP mismatch:", phone, "provided:", otp, "stored:", storedOtp)
		return nil, errors.New("کد OTP نادرست است")
	}

	cache.DeleteOTP(phone)
	log.Print("✅ OTP verified successfully:", phone)

	// بررسی کاربر در دیتابیس
	var user User
	err = db.GetDb().Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// کاربر وجود ندارد → ایجاد کاربر جدید
			user = User{
				Phone: phone,
			}
			if err := db.GetDb().Create(&user).Error; err != nil {
				log.Print("❌ Failed to create user:", err)
				return nil, errors.New("خطا در ایجاد کاربر")
			}
			log.Print("✅ New user created:", phone)
		} else {
			log.Print("❌ DB error:", err)
			return nil, errors.New("خطا در دسترسی به دیتابیس")
		}
	} else {
		log.Print("ℹ️ User already exists:", phone)
	}

	return &user, nil
}
