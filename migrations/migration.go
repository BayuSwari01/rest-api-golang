package migrations

import (
	"log"
	"rest-api-golang/config"
	"rest-api-golang/models"
)

func RunMigrations() {
	db, _ := config.Connect()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	
	err := db.AutoMigrate(&models.User{}, &models.Transaction{})

    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    log.Println("Migration completed successfully")
}
