package transaction

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() []Transaction
	GetByID(id string) (Transaction, error)
	Create(transaction *Transaction)
	Delete(id string) error
}

type repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &repository{database: database}
}

func (r *repository) Create(t *Transaction) {
	r.database.Create(t)
}


func (r *repository) Delete(id string) error {
	t := &Transaction{}
	fmt.Println(t)
	r.database.Find(t, "id = ?", id)

	if t.ID == uuid.Nil {
		return errors.New("transaction not found")
	}

	err := r.database.Delete(t, "id = ?", id).Error

    if err != nil {
        return errors.New("transaction could not be deleted")
    }

	return nil
}

func (r *repository) GetAll() []Transaction {
	var transactions []Transaction
	r.database.Find(&transactions)
	return transactions
}

func (r *repository) GetByID(id string) (Transaction, error) {
	t := Transaction{}
	result := r.database.First(&t, "id = ?", id)
	return t, result.Error
}