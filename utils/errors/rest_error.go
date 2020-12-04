package errors

import "net/http"

// RestError error message struct
type RestError struct {
  Message string `json:"message"`
  Status int `json:"status"`
  Error string `json:"error"`
}

// NewBadRequestError creates a new bad request error
func NewBadRequestError(message string) *RestError{
  return &RestError{
    Message: message,
    Status: http.StatusBadRequest,
    Error: "bad_request",
  }
}

// NewNotFoundError creates a new bad request error
func NewNotFoundError(message string) *RestError{
  return &RestError{
    Message: message,
    Status: http.StatusNotFound,
    Error: "not_found",
  }
}