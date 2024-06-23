package service

import (
	"errors"
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/model/web"
	"mikti-depublic/repository"
	"time"
)

type TransactionServiceImpl struct {
	transactionRepo repository.TransactionRepository
	eventRepo       repository.EventRepository
}

func NewTransactionServiceImpl(transactionRepo repository.TransactionRepository, eventRepo repository.EventRepository) TransactionService {
	return &TransactionServiceImpl{transactionRepo: transactionRepo, eventRepo: eventRepo}
}

func (service *TransactionServiceImpl) GetAllTransactions(userPayload helper.JwtCustomClaims) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	transactions, err := service.transactionRepo.GetAllTransactions(userPayload.ID)
	if err != nil {
		return []domain.Transaction{}, err
	}
	return transactions, nil
}

func (service *TransactionServiceImpl) ConfirmPayment(id string) (bool, error) {
	return service.transactionRepo.ConfirmPayment(id)
}

func (service *TransactionServiceImpl) CreateTransaction(request web.TransactionRequest, userPayload helper.JwtCustomClaims) (domain.Transaction, error) {
	_, err := service.eventRepo.CheckTicketAvailability(request.EventID, request.Quantity)
	if err != nil {
		return domain.Transaction{}, err
	}

	eventData, err := service.eventRepo.GetEvent(request.EventID)
	if err != nil {
		return domain.Transaction{}, err
	}
	transactionId := helper.GenerateId(5, "transaction")

	totalPrice := eventData.Price * float64(request.Quantity)
	transaction := domain.Transaction{
		ID:         transactionId,
		UserID:     userPayload.ID,
		EventID:    request.EventID,
		Date:       time.DateTime,
		Quantity:   request.Quantity,
		TotalPrice: totalPrice,
		Status:     "ordered",
		CreatedAt:  time.Now(),
	}
	newTransaction, err := service.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return domain.Transaction{}, err
	}

	isSuccessUpdateSoldTicket, err := service.eventRepo.UpdateSoldTicket(request.EventID, request.Quantity)
	if err != nil {
		return domain.Transaction{}, err
	}

	if !isSuccessUpdateSoldTicket {
		return domain.Transaction{}, errors.New("failed to update sold ticket")
	}

	return newTransaction, nil
}

func (service *TransactionServiceImpl) GetTransactionById(id string) (domain.Transaction, error) {
	return service.transactionRepo.GetTransactionById(id)
}
