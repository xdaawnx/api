package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response struct to format API responses
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Data is omitted if empty
}
type ResponseValidation struct {
	Message    string      `json:"message"`
	Validation interface{} `json:"validation"`
}
type ResponseWithoutData struct {
	Message string `json:"message"`
}

// FormatResponse formats the response using a struct
func FormatResponse(c echo.Context, code int, data interface{}, message string) error {
	response := Response{
		Message: message,
		Data:    data,
	}

	// If code is 204 (No Content), exclude data
	if code == http.StatusNoContent {
		response.Data = nil
	}

	return c.JSON(code, response)
}
