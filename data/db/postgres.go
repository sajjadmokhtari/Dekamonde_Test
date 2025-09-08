package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	// مقادیر را می‌توان از ENV هم گرفت
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print("❌ اتصال به دیتابیس PostgreSQL ناموفق بود: ", err)
		return err
	}

	log.Print("✅ اتصال موفق به PostgreSQL برقرار شد")
	DB = db
	return nil
}

func GetDb() *gorm.DB {
	return DB
}
