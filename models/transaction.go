package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID 					string		`gorm:"type:char(36);primaryKey"`
	UserID				string		`gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	User				User		`gorm:"foreignKey:UserID;references:ID"`
	TransactionAction	string		`gorm:"column:transaction_action;type:enum('transfer', 'payment', 'top_up');not nul" json:"transaction_action"`
	TransactionType		string		`gorm:"column:transaction_type;type:enum('DEBIT', 'CREDIT'); not null" json:"transaction_type"`	
	Status				string		`gorm:"type:enum('SUCCESS', 'FAILED');not null"`
	Amount				float64		`gorm:"type:decimal(15,2);not null"`
	Remarks				string		`gorm:"type:varchar(255);default:''"`
	BalanceBefore		float64		`gorm:"column:balance_before;type:decimal(15,2);not null" json:"balance_before"`
	BalanceAfter		float64		`gorm:"column:balance_after;type:decimal(15,2);not null" json:"balance_after"`
	CreatedAt 			time.Time 	`gorm:"column:created_at"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
    t.ID = uuid.NewString()
    return
}