package main

import (
	"log"
	"rest-api-golang/migrations"
	"rest-api-golang/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	migrations.RunMigrations()

	r := gin.Default()

	routes.AuthRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}