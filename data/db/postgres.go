package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	// مقادیر را می‌توان از ENV هم گرفت
	dsn := "host=localhost user=postgres password=admin dbname=MarketPlace_db port=5432 sslmode=disable"
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
