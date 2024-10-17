package handlers

import (
	"net/http"

	request "api/handlers/Request"
	"api/services"
	util "api/utils"

	"github.com/labstack/echo/v4"
)

// TimeHandler defines the HTTP handler for time-related endpoints
type TimeHandler struct {
	Service *services.TimeService
}

// GetCurrentTime handles the GetCurrentTime API
// @Summary Get current time by timezone
// @Description Retrieve current time based on the provided timezone
// @Tags time
// @Produce json
// @Security BearerAuth
// @Param timezone query string true "Timezone"
// @Success 200 {object} util.Response{data=dto.TimeResponse}
// @Success 201 {object} util.ResponseWithoutData{}
// @Router /GetCurrentTime [get]
func (h *TimeHandler) GetCurrentTime(c echo.Context) error {
	req := new(request.TimeReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	timeResponse, err := h.Service.GetCurrentTime(req.Timezone)
	if err != nil {
		if customErr, ok := err.(*util.CustomError); ok {
			return customErr
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error occurred")
	}
	return util.FormatResponse(c, http.StatusOK, timeResponse, "ok")
}

// Retrieves details of the logged-in user
// @Summary Get user details
// @Description Retrieves details of the logged-in user
// @Tags users
// @Security BearerAuth
// @Produce  json
// @Success 200 {object} util.Response{data=dto.TimeResponse}
// @Router /users/test [get]
func (h *TimeHandler) Test(c echo.Context) {
	return
}
