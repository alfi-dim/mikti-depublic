package service

import (
	"mikti-depublic/model/entity"
	"mikti-depublic/model/web"
)

type EventService interface {
	CreateEvent(request web.EventServiceReq) (map[string]interface{}, error)
	GetEvent(eventId string) (entity.EventEntity, error)
	GetListEvent() ([]entity.EventEntity, error)
	UpdateEvent(request web.EventServiceReq, pathId string) (map[string]interface{}, error)
	DeleteEvent(Id string) (entity.EventEntity, error)
}
