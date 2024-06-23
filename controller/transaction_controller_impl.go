package controller

import (
	"mikti-depublic/helper"
	"mikti-depublic/model"
	"mikti-depublic/model/web"
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionControllerImpl struct {
	service service.TransactionService
}

func NewTransactionControllerImpl(transactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{service: transactionService}
}

func (controller *TransactionControllerImpl) CreateTransaction(c echo.Context) error {
	transaction := new(web.TransactionRequest)

	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(transaction); err != nil {
		return err
	}

	userPayload := c.Get("claims").(helper.JwtCustomClaims)
	createTransaction, err := controller.service.CreateTransaction(*transaction, userPayload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Added", createTransaction))
}

func (controller *TransactionControllerImpl) GetTransactionById(c echo.Context) error {
	id := c.Param("id")

	getTransaction, err := controller.service.GetTransactionById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Transaction Data", getTransaction))
}

func (controller *TransactionControllerImpl) GetAllTransactions(c echo.Context) error {
	userPayload := c.Get("claims").(helper.JwtCustomClaims)

	getTransactions, err := controller.service.GetAllTransactions(userPayload)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Get List", getTransactions))
}

func (controller *TransactionControllerImpl) ConfirmPayment(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, "ID is required", nil))
	}

	confirmPayment, err := controller.service.ConfirmPayment(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Confirm Payment", confirmPayment))
}
