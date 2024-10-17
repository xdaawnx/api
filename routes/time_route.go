package routes

import (
	"api/handlers"
	"api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// NewTimeHandler registers time routes
func NewTimeHandler(e *echo.Echo, service *gorm.DB) {
	timeService := services.NewTimeService()

	handler := &handlers.TimeHandler{Service: timeService}
	e.GET("/GetCurrentTime", handler.GetCurrentTime)
}
