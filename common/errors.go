package common

import (
	"net/http"
)

type AppError struct {
	HttpCode int
	Message  string
}

var BadRequestError = AppError{
	HttpCode: http.StatusBadRequest,
	Message:  "Invalid Id, please provide a valid integer num",
}

var NotFoundError = AppError{
	HttpCode: http.StatusNotFound,
	Message:  "Id Not Found",
}

var CustomError = AppError{
	HttpCode: http.StatusInternalServerError,
	Message:  "Something Failed",
}

var OpenCsvError = AppError{
	HttpCode: http.StatusInternalServerError,
	Message:  "Cannot open csv",
}
