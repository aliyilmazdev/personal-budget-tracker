package transaction

import (
	"errors"
)

type Service interface {
	GetAll() ([]Transaction, error)
	GetByID(id string) (Transaction, error)
	Create(transaction *Transaction)
	Delete(id string) error
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

func (s *service) Delete(id string) error{
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAll() ([]Transaction, error) {
	transactions := s.repository.GetAll()
	
	if len(transactions) == 0 {
		return nil, errors.New("no transactions found")
	}

	return transactions, nil
}

func (s *service) GetByID(id string) (Transaction, error) {
	t, err := s.repository.GetByID(id)
	if err != nil {
		return Transaction{}, err
	}
	return t, nil
	
}