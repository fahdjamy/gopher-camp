package utils

import "time"

type ArrayResponse[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type Failure struct {
	Error string    `json:"error"`
	Date  time.Time `date:"datetime"`
}

type DataResponse[T any] struct {
	Data T         `json:"data"`
	Date time.Time `date:"datetime"`
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

func FailureResponse(err error) Failure {
	return Failure{
		Date:  time.Now(),
		Error: err.Error(),
	}
}

func SingleObject[T any](data T) DataResponse[T] {
	return DataResponse[T]{
		Date: time.Now(),
		Data: data,
	}
}
