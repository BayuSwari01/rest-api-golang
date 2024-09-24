package models

type LoginInput struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pin         string `json:"pin" binding:"required"`
}