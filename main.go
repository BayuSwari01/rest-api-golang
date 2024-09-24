package main

import (
	"log"
	"rest-api-golang/migrations"

	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	migrations.RunMigrations()
}