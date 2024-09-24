package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID 				string		`gorm:"type:char(36);primaryKey"`
	FirstName     	string    	`gorm:"column:first_name;type:varchar(100);not null"`
	LastName     	string    	`gorm:"column:last_name;type:varchar(100);not null"`
	PhoneNumber		string		`gorm:"column:phone_number;type:varchar(15);not null"`
	Pin				string		`gorm:"type:varchar(6);not null"`
	Address			string		`gorm:"type:varchar(255);not null"`
    Balance  		float64   	`gorm:"type:decimal(15,2);default:0.00"`
	CreatedAt 		time.Time 	`gorm:"column:created_at"`
    UpdatedAt		time.Time	`gorm:"column:updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.NewString()
    return
}