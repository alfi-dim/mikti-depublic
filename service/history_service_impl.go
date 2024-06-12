package service

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/repository"
)

type historyServiceImpl struct {
	repository repository.HistoryRepository
}

func NewHistoryServiceImpl(repository repository.HistoryRepository) HistoryService {
	return &historyServiceImpl{repository: repository}
}

func (s *historyServiceImpl) CreateTransaction(transaction *domain.Transaction) error {
	return s.repository.Create(transaction)
}

func (s *historyServiceImpl) GetHistories() ([]domain.Transaction, error) {
	return s.repository.FindAll()
}

func (s *historyServiceImpl) GetHistoryByID(id string) (*domain.Transaction, error) {
	return s.repository.FindByID(id)
}

func (s *historyServiceImpl) GetHistoryByStatus(status string) ([]domain.Transaction, error) {
	return s.repository.FindByStatus(status)
}
