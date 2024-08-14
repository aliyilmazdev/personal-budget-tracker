package transaction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID 				uuid.UUID 	`gorm:"type:uuid"`
	Amount 			float64 	`json:"amount"`
	Description 	string 		`json:"description"`
	TransactionType string 		`json:"transaction_type"`
}