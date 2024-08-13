package transaction

import "errors"

type Service interface {
	GetAll() ([]Transaction, error)
	GetByID(id uint) (Transaction, error)
	Create(transaction *Transaction)
	Delete(id uint)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(t *Transaction) {
	s.repository.Create(t)
}

func (s *service) Delete(id uint) {
	s.repository.Delete(id)
}

func (s *service) GetAll() ([]Transaction, error) {
	transactions := s.repository.GetAll()
	
	if len(transactions) == 0 {
		return nil, errors.New("no transactions found")
	}

	return transactions, nil
}

func (s *service) GetByID(id uint) (Transaction, error) {
	return s.repository.GetByID(id)
}