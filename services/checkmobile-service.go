package services

import (
	"errors"
	"regexp"
)

// IsIranianPhone بررسی می‌کند شماره موبایل ایرانی باشد
func IsIranianPhone(phone string) error {
	// شماره موبایل ایرانی معمولاً با 09 شروع می‌شود و 11 رقم دارد
	re := regexp.MustCompile(`^09\d{9}$`)
	if !re.MatchString(phone) {
		return errors.New("شماره موبایل معتبر نیست")
	}
	return nil
}
