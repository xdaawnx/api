package util

import "net/http"

type CustomError struct {
	Code       int
	Message    string
	Validation interface{}
}

// Error implements the error interface
func (e *CustomError) Error() string {
	return e.Message
}

var BadRequest = &CustomError{Code: http.StatusBadRequest, Message: "bad request"}
var NotFound = &CustomError{Code: http.StatusNotFound, Message: "not found"}
var InternalServerError = &CustomError{Code: http.StatusInternalServerError, Message: "internal server error"}
var Unauthorized = &CustomError{Code: http.StatusUnauthorized, Message: "unauthorized"}
