package routes

import (
	"api/handlers"
	"api/models"
	"api/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// NewPersonHandler registers person routes
func NewPersonHandler(e *echo.Echo, db *gorm.DB) {
	personRepo := models.NewPersonQueries(db)
	personService := services.NewPersonService(personRepo)

	handler := &handlers.PersonHandler{Service: personService}
	e.GET("/GetCountry/:name", handler.GetCountryByName)
}
