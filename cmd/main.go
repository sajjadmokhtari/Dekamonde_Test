package main

import (
	"dekamonde/data/db"
	_ "dekamonde/docs"
	"dekamonde/router"
	"log"
)

func main() {

	db.InitRedis()

	r := router.SetupRoutes()

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("❌ سرور نتوانست اجرا شود: %v", err)
	}

}
