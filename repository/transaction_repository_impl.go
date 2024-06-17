package repository

import (
	"errors"
	"gorm.io/gorm"
	"mikti-depublic/model/domain"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (repo *TransactionRepositoryImpl) CreateTransaction(newTransaction domain.Transaction) (domain.Transaction, error) {
	err := repo.db.Create(&newTransaction).Error

	if err != nil {
		return domain.Transaction{}, errors.New("Failed when create new transaction")
	}

	return newTransaction, nil
}

func (repo *TransactionRepositoryImpl) GetTransactionById(id string) (domain.Transaction, error) {
	var transactionData domain.Transaction

	err := repo.db.First(&transactionData, "id = ?", id).Error

	if err != nil {
		return domain.Transaction{}, errors.New("Transaction not found")
	}

	return transactionData, nil
}

func (repo *TransactionRepositoryImpl) GetAllTransactions(ownerId string) ([]domain.Transaction, error) {
	var transactionsData []domain.Transaction

	err := repo.db.Find(&transactionsData, "users_id = ?", ownerId).Error

	if err != nil {
		return []domain.Transaction{}, errors.New("Transactions not found")
	}

	return transactionsData, nil
}

func (repo *TransactionRepositoryImpl) ConfirmPayment(id string) (bool, error) {
	var transactionData domain.Transaction

	err := repo.db.First(&transactionData, "id = ?", id).Error

	if err != nil {
		return false, errors.New("Transaction not found")
	}

	transactionData.Status = "paid"

	err = repo.db.Save(&transactionData).Error

	if err != nil {
		return false, errors.New("Failed when update transaction")
	}

	return true, nil
}
