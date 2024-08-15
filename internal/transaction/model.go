package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID 				uuid.UUID 	`gorm:"type:uuid"`
	Amount 			float64 	`json:"amount"`
	Description 	string 		`json:"description"`
	TransactionType string 		`json:"transaction_type"`
	CreatedAt       time.Time 	`json:"created_at"`
}