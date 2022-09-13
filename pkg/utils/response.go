package utils

import (
	"gopher-camp/pkg/constants"
	"time"
)

type ArrayResponse[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type ErrorData struct {
	Error string `json:"error"`
	Date  string `json:"datetime"`
}

type ErrorWithMessage struct {
	Error   ErrorData `json:"error"`
	Message string    `json:"message"`
}

type DataResponse[T any] struct {
	Data T      `json:"data"`
	Date string `json:"datetime"`
}

func SuccessArray[T any](data []T, msg string) ArrayResponse[T] {
	var response = make([]T, 0)
	if data != nil {
		response = data
	}
	return ArrayResponse[T]{
		Message: msg,
		Data:    response,
	}
}

func CreateFailure(err error) ErrorData {
	return ErrorData{
		Date:  DateTime(time.Now(), constants.DateTimeResponseFormat),
		Error: err.Error(),
	}
}

func CreateFailureWithMessage(err error, message string) ErrorWithMessage {
	return ErrorWithMessage{
		Message: message,
		Error:   CreateFailure(err),
	}
}

func SingleObject[T any](data T) DataResponse[T] {
	return DataResponse[T]{
		Date: DateTime(time.Now(), constants.DateTimeMinSecResponseFormat),
		Data: data,
	}
}
