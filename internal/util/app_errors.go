package util

import (
	"net/http"
	"strings"
)

type AppErrors struct {
	Errors []ItemError `json:"errors"`
}

type ItemError struct {
	Code    string `json:"code,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e AppErrors) Error() string {

	message := []string{}
	for _, v := range e.Errors {
		message = append(message, v.Message)
	}

	return strings.Join(message, ",")
}

func (e *AppErrors) Add(err ItemError) {

	e.Errors = append(e.Errors, err)
}

func NewNotFoundError(message string) ItemError {

	return ItemError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() ItemError {

	return ItemError{
		Status:  http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewValidationError(message string) ItemError {

	return ItemError{
		Status:  http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewBadRequestError() ItemError {

	return ItemError{
		Status:  http.StatusBadRequest,
		Message: "bad request",
	}
}

func NewCustomBadRequestError(message string) ItemError {

	return ItemError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func NewCustomError(message string) ItemError {

	return ItemError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func NewUnauthorizedError(message string) ItemError {

	return ItemError{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}
