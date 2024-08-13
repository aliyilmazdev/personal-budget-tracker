package transaction

import "gorm.io/gorm"

type Repository interface {
	GetAll() []Transaction
	GetByID(id uint) (Transaction, error)
	Create(transaction *Transaction)
	Delete(id uint)
}

type repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &repository{database: database}
}

func (r *repository) Create(t *Transaction) {
	r.database.Create(&t)
}


func (r *repository) Delete(id uint) {
	r.database.Delete(&Transaction{}, id)
}

func (r *repository) GetAll() []Transaction {
	var transactions []Transaction
	r.database.Find(&transactions)
	return transactions
}

func (r *repository) GetByID(id uint) (Transaction, error) {
	var t Transaction
	result := r.database.First(&t, "id = ?", id)
	return t, result.Error
}