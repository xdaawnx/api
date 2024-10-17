package middleware

import (
	util "api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if he, ok := err.(*util.CustomError); ok {
		switch he.Code {
		case http.StatusBadRequest:
			if he.Validation != nil {
				c.JSON(he.Code, util.ResponseValidation{
					Message:    he.Message,
					Validation: he.Validation,
				})
				return
			}
			c.JSON(he.Code, util.ResponseWithoutData{
				Message: he.Message,
			})
			return
		}
		c.JSON(he.Code, util.ResponseWithoutData{
			Message: he.Message,
		})
		return
	}
	// Handle echo HTTP errors
	if ex, ok := err.(*echo.HTTPError); ok {
		if unmarshalErr, ok := ex.Internal.(*json.UnmarshalTypeError); ok {
			// Create a formatted error response
			regExp := regexp.MustCompile("[0-9]")

			inputVal := regExp.ReplaceAllString(unmarshalErr.Value, "")
			inputType := regExp.ReplaceAllString(unmarshalErr.Type.String(), "")
			resp := util.ResponseValidation{
				Message: "bad request",
				Validation: map[string]interface{}{
					unmarshalErr.Field: fmt.Sprintf("%s expected %s got %s",
						unmarshalErr.Field, inputType, inputVal),
				},
			}
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		c.JSON(ex.Code, util.ResponseWithoutData{
			Message: fmt.Sprintf("%s", ex.Message),
		})
		return
	}
	// Handle unexpected errors (generic case)
	log.Printf("Unexpected error: %v", err)
	c.JSON(http.StatusInternalServerError, util.ResponseWithoutData{
		Message: "Internal Server Error",
	})
}
