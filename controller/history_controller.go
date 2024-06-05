package controller

import (
	"mikti-depublic/service"

	"github.com/labstack/echo/v4"
)

type HistoryController interface {
	GetHistory(c echo.Context) error
	GetHistoryByID(c echo.Context) error
	GetHistoryByStatus(c echo.Context) error
}

func NewHistoryController(historyService service.HistoryService) HistoryController {
	return &HistoryControllerImpl{historyService: historyService}
}
