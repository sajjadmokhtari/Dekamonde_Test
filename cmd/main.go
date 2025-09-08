package main

import (
	"log"

	"dekamonde/data/db"
	"dekamonde/data/db/migration"
	_ "dekamonde/docs"
	"dekamonde/router"
)

func main() {

	// اتصال به Redis
	if err := db.InitRedis(); err != nil {
		log.Fatalf("❌ اتصال به Redis برقرار نشد: %v", err)
	}

	// اتصال به PostgreSQL
	if err := db.InitDb(); err != nil {
		log.Fatalf("❌ اتصال به PostgreSQL برقرار نشد: %v", err)
	}

	// اجرای AutoMigrate برای ایجاد جداول
	if err := migration.RunMigrations(); err != nil {
		log.Fatalf("❌ Migration شکست خورد: %v", err)
	}

	// راه‌اندازی Router
	r := router.SetupRoutes()

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("❌ سرور نتوانست اجرا شود: %v", err)
	}
}
