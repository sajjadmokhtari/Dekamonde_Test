package services

import (
	"errors"
	"log"
	"strings"

	"dekamonde/cache"
)

func VerifyOTP(phone, otp string) error {
	storedOtp, err := cache.GetOTP(phone)
	if err != nil {
		log.Print("❌ OTP not found or expired:", phone, err)
		return errors.New("کد OTP منقضی شده یا موجود نیست")
	}

	storedOtp = strings.TrimSpace(storedOtp)
	otp = strings.TrimSpace(otp)

	log.Println("🔹 Stored OTP:", storedOtp)
	log.Println("🔹 Provided OTP:", otp)

	if storedOtp != otp {
		log.Print("❌ OTP mismatch:", phone, "provided:", otp, "stored:", storedOtp)
		return errors.New("کد OTP نادرست است")
	}

	cache.DeleteOTP(phone)
	log.Print("✅ OTP verified successfully:", phone)
	return nil
}
