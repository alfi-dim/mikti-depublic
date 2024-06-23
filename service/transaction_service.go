package service

import (
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/model/web"
)

type TransactionService interface {
	CreateTransaction(request web.TransactionRequest, userPayload helper.JwtCustomClaims) (domain.Transaction, error)
	GetTransactionById(id string) (domain.Transaction, error)
	GetAllTransactions(userPayload helper.JwtCustomClaims) ([]domain.Transaction, error)
	ConfirmPayment(id string) (bool, error)
}
