package repository

import (
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(transaction *domain.Transaction) error
	FindAll() ([]domain.Transaction, error)
	FindByID(id string) (*domain.Transaction, error)
	FindByStatus(status string) ([]domain.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) HistoryRepository {
	return NewHistoryRepositoryImpl(db)
}
