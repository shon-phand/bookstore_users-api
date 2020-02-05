package errors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

func StatusBadRequestError(message string) *RestErr {

	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
	}
}

func StatusNotFoundError(message string) *RestErr {

	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
