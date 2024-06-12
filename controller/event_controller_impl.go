package controller

import (
	"mikti-depublic/model"
	"mikti-depublic/model/web"
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventControllerImpl struct {
	service service.EventService
}

func NewEventController(service service.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		service: service,
	}
}

func (controller *EventControllerImpl) CreateEvent(c echo.Context) error {
	event := new(web.EventServiceReq)

	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(event); err != nil {
		return err
	}

	saveEvent, errCreateEvent := controller.service.CreateEvent(*event)

	if errCreateEvent != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errCreateEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Added", saveEvent))
}

func (controller *EventControllerImpl) GetEvent(c echo.Context) error {
	id := c.Param("id")

	getEvent, errGetEvent := controller.service.GetEvent(id)

	if errGetEvent != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errGetEvent.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Event Data", getEvent))
}

func (controller *EventControllerImpl) GetListEvent(c echo.Context) error {
	getEvents, errGetEvents := controller.service.GetListEvent()

	if errGetEvents != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errGetEvents.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Get List", getEvents))

}

func (controller *EventControllerImpl) UpdateEvent(c echo.Context) error {
	event := new(web.EventServiceReq)
	id := c.Param("id")

	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}
	eventUpdate, errUpdate := controller.service.UpdateEvent(*event, id)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errUpdate.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Update Data", eventUpdate))
}

func (controller *EventControllerImpl) DeleteEvent(c echo.Context) error {
	id := c.Param("id")

	deleteEvent, errDelete := controller.service.DeleteEvent(id)
	if errDelete != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errDelete.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "Success Delete", deleteEvent))

}
