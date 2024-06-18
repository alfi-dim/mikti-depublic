package service

import (
	"mikti-depublic/model/entity"
	"mikti-depublic/model/web"

	"github.com/labstack/echo/v4"
)

type EventService interface {
	CreateEvent(request web.EventServiceReq, c echo.Context) (map[string]interface{}, error)
	GetEvent(eventId string) (entity.EventEntity, error)
	GetListEvent() ([]entity.EventEntity, error)
	UpdateEvent(request web.EventServiceReq, pathId string) (map[string]interface{}, error)
	DeleteEvent(Id string) (entity.EventEntity, error)
}
