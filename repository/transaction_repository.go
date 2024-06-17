package repository

import "mikti-depublic/model/domain"

type TransactionRepository interface {
	CreateTransaction(transaction domain.Transaction) (domain.Transaction, error)
	GetTransactionById(id string) (domain.Transaction, error)
	GetAllTransactions(ownerId string) ([]domain.Transaction, error)
	ConfirmPayment(id string) (bool, error)
}
