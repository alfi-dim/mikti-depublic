package repository

import (
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type historyRepositoryImpl struct {
	db *gorm.DB
}

func NewHistoryRepositoryImpl(db *gorm.DB) HistoryRepository {
	return &historyRepositoryImpl{db: db}
}

func (r *historyRepositoryImpl) Create(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *historyRepositoryImpl) FindAll() ([]domain.Transaction, error) {
	var histories []domain.Transaction
	if err := r.db.Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *historyRepositoryImpl) FindByID(id string) (*domain.Transaction, error) {
	var histories domain.Transaction
	if err := r.db.First(&histories, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &histories, nil
}

func (r *historyRepositoryImpl) FindByStatus(status string) ([]domain.Transaction, error) {
	var histories []domain.Transaction
	if err := r.db.Where("status = ?", status).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}
