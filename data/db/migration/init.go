package migration

import (
	"log"

	"dekamonde/data/db"
	"dekamonde/data/models"
)

func RunMigrations() error {
	db := db.GetDb() // اتصال به gorm.DB

	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully")

	return  err
}
