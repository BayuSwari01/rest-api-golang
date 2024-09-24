package migrations

import (
	"fmt"
	"log"
	"os"
	"rest-api-golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RunMigrations() {
	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    err = db.AutoMigrate(&models.User{}, &models.Transaction{})
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    log.Println("Migration completed successfully")
}
