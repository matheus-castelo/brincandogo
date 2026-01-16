package main

import "fmt"

type ApiErr struct {
	Name    string   `json:"name"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ApiErr) Error() string {
	return fmt.Sprintf("error: %s, code: %d", e.Name, e.Code)
}

func NewBadRequestError(message string, causes []Causes) *ApiErr {
	return &ApiErr{
		Name:    "bad_request",
		Code:    400,
		Message: message,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *ApiErr {
	return &ApiErr{
		Name:    "not_found",
		Code:    404,
		Message: message,
	}
}