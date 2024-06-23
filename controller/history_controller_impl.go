package controller

import (
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HistoryControllerImpl struct {
	historyService service.HistoryService
}

func NewHistoryControllerImpl(historyService service.HistoryService) *HistoryControllerImpl {
	return &HistoryControllerImpl{historyService: historyService}
}

func (hc *HistoryControllerImpl) GetHistory(c echo.Context) error {
	histories, err := hc.historyService.GetHistories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, histories)
}

func (hc *HistoryControllerImpl) GetHistoryByID(c echo.Context) error {
	id := c.Param("id")
	histories, err := hc.historyService.GetHistoryByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, histories)
}

func (hc *HistoryControllerImpl) GetHistoryByStatus(c echo.Context) error {
	status := c.Param("status")
	histories, err := hc.historyService.GetHistoryByStatus(status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, histories)
}
