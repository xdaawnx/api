package handlers

import (
	"net/http"

	"api/services"
	util "api/utils"

	"github.com/labstack/echo/v4"
)

// PersonHandler defines the HTTP handler for person-related endpoints
type PersonHandler struct {
	Service *services.PersonService
}

// GetCountryByName handles the GetCountry API
// @Summary Get country by person name
// @Description Retrieve country based on person's name
// @Tags persons
// @Produce json
// @Param name path string true "Person Name"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /GetCountry/{name} [get]
func (h *PersonHandler) GetCountryByName(c echo.Context) error {
	name := c.Param("name")
	country, err := h.Service.GetCountry(name)
	if err != nil {
		if customErr, ok := err.(*util.CustomError); ok {
			return customErr
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error occurred")
	}
	return util.FormatResponse(c, http.StatusOK, map[string]string{"country": country}, "ok")
}
