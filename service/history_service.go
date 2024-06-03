package service

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/repository"
)

type HistoryService interface {
	CreateTransaction(transaction *domain.Transaction) error
	GetHistories() ([]domain.Transaction, error)
	GetHistoryByID(id string) (*domain.Transaction, error)
	GetHistoryByStatus(status string) ([]domain.Transaction, error)
}

func NewHistoryService(repository repository.HistoryRepository) HistoryService {
	return NewHistoryServiceImpl(repository)
}
