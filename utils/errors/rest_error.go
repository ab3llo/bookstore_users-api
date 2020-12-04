package errors

import "net/http"

// RestError error message struct
type RestError struct {
  Message string `json:"message"`
  Status int `json:"status"`
  Error string `json:"error"`
}

// NewBadRequestErr creates a new bad request error
func NewBadRequestErr(message string) *RestError{
  return &RestError{
    Message: message,
    Status: http.StatusBadRequest,
    Error: "bad_request",
  }
}