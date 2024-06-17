package controller

import "github.com/labstack/echo/v4"

type TransactionController interface {
	CreateTransaction(c echo.Context) error
	GetTransactionById(c echo.Context) error
	GetAllTransactions(c echo.Context) error
	ConfirmPayment(c echo.Context) error
}
