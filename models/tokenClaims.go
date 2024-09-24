package models

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	FirstName string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}