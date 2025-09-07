package services

import (
	"dekamonde/cache"
	"errors"
	"log"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// تولید OTP شش رقمی
func GenerateOTP() string {
	otp := ""
	for i := 0; i < 6; i++ {
		otp += string('0' + rune(r.Intn(10)))
	}
	return otp
}

// SendOTP با محدودسازی درخواست‌ها
func SendOTP(phone string) error {
	// بررسی محدودیت زمانی بین ارسال‌ها
	if !cache.CanSendOTP(phone) {
		log.Println("OTP request blocked: please try later | phone:", phone)
		return errors.New("لطفاً کمی صبر کنید و دوباره تلاش کنید")
	}

	// بررسی تعداد درخواست‌ها در ۱۰ دقیقه اخیر
	if cache.OTPRequestCount(phone) >= 3 {
		log.Println("OTP request limit reached | phone:", phone)
		return errors.New("تعداد درخواست‌های مجاز در ۱۰ دقیقه گذشته تمام شده است")
	}

	otp := GenerateOTP()
	log.Println("Generated OTP | phone:", phone, " | otp:", otp)

	// ذخیره OTP در Redis
	if err := cache.SetOTP(phone, otp); err != nil {
		log.Println("Failed to save OTP in cache | phone:", phone, " | error:", err)
		return err
	}

	// ثبت زمان ارسال و افزایش شمارش درخواست
	cache.MarkOTPSent(phone)
	cache.IncrementOTPRequest(phone)

	log.Println("OTP sent successfully | phone:", phone, " | otp:", otp)
	return nil
}
