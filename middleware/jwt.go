package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"rest-api-golang/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))


func GenerateAccessToken(firstName string, phoneNumber string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.TokenClaims{
		FirstName: firstName,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(firstName string, phoneNumber string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.TokenClaims{
		FirstName: firstName,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		claims, err := ValidateJWT(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("firstName", claims.FirstName)
		ctx.Set("phoneNumber", claims.PhoneNumber)
		ctx.Next()
	}
}