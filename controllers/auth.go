package controllers

import (
	"log"
	"net/http"
	"rest-api-golang/config"
	"rest-api-golang/middleware"
	"rest-api-golang/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context){
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

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(user.Pin), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error encrypting Pin"})
		return
	}
	user.Pin = string(hashedPin)

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response := gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"user_id":     user.ID,
			"first_name":  user.FirstName,
			"last_name":   user.LastName,
			"phone_number": user.PhoneNumber,
			"address":     user.Address,
			"created_date": user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	ctx.JSON(http.StatusCreated, response)
}

func Login(ctx *gin.Context)  {
	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User
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

	if err := db.Where("phone_number = ?", input.PhoneNumber).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid phone number"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(input.Pin)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Phone Number and PIN doesn't match."})
		return
	}

	accessToken, err := middleware.GenerateAccessToken(user.FirstName, user.PhoneNumber)
	if err != nil {
		log.Printf("Failed to generate access token: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate access token"})
		return
	}

	refreshToken, err := middleware.GenerateRefreshToken(user.FirstName, user.PhoneNumber)
	if err != nil {
		log.Printf("Failed to generate refresh token: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"access_token": accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func Test(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, "asd")
}

func TestAuth(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, "masuk sini")
}