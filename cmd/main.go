package main

import (
	"dekamonde/data/db"
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
