package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims شامل شماره تلفن و نقش کاربر
type CustomClaims struct {
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("supersecretkey") // بهتر است در env قرار گیرد

// GenerateJWT تولید JWT با اعتبار 24 ساعت
func GenerateJWT(phone, role string) (string, error) {
	claims := CustomClaims{
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT اعتبارسنجی JWT و بازگرداندن Claims
func ValidateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// بررسی الگوریتم
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("توکن نامعتبر است")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("توکن نامعتبر یا منقضی شده است")
	}

	return claims, nil
}
