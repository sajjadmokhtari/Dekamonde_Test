package cache

import (
	"dekamonde/data/db"
	"log"
	"strconv"
	"time"
	"github.com/redis/go-redis/v9"

	
)

func SetOTP(phone, otp string) error {
	key := "otp:" + phone
	return db.RedisClient.Set(db.Ctx, key, otp, 2*time.Minute).Err()
}

func GetOTP(phone string) (string, error) {
	key := "otp:" + phone
	val, err := db.RedisClient.Get(db.Ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, err
}

func CanSendOTP(phone string) bool {
	key := "otp:last:" + phone
	lastTimeStr, err := db.RedisClient.Get(db.Ctx, key).Result()
	if err == redis.Nil {
		return true
	}
	if err != nil {
		log.Println("Redis error in CanSendOTP:", err)
		return true
	}
	lastTime, err := strconv.ParseInt(lastTimeStr, 10, 64)
	if err != nil {
		log.Println("Failed to parse last OTP timestamp:", err)
		return true
	}
	return time.Now().Unix()-lastTime >= 60
}

func MarkOTPSent(phone string) {
	key := "otp:last:" + phone
	if err := db.RedisClient.Set(db.Ctx, key, time.Now().Unix(), 10*time.Minute).Err(); err != nil {
		log.Println("Failed to mark OTP sent:", err)
	}
}

func OTPRequestCount(phone string) int {
	key := "otp:count:" + phone
	countStr, err := db.RedisClient.Get(db.Ctx, key).Result()
	if err == redis.Nil {
		return 0
	}
	if err != nil {
		return 0
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0
	}
	return count
}

func IncrementOTPRequest(phone string) {
	key := "otp:count:" + phone
	if err := db.RedisClient.Incr(db.Ctx, key).Err(); err != nil {
		log.Println("Failed to increment OTP request count:", err)
	}
	if err := db.RedisClient.Expire(db.Ctx, key, 10*time.Minute).Err(); err != nil {
		log.Println("Failed to set expire for OTP request count:", err)
	}
}


func DeleteOTP(phone string) error {
	key := "otp:" + phone
	err := db.RedisClient.Del(db.Ctx, key).Err()
	if err != nil {
		log.Println("❌ Failed to delete OTP:", phone, err)
	}
	return err
}
