package utils

import (
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/types"
	"time"
)

type ArrayResponse[T any] struct {
	TotalCount int    `json:"totalCount"`
	ItemCount  int    `json:"itemCount"`
	Message    string `json:"message"`
	Data       []T    `json:"data"`
}

type ErrorData struct {
	Error  error  `json:"error"`
	Source string `json:"source"`
	Date   string `json:"datetime"`
}

type ErrorWithMessage struct {
	Message    string    `json:"message"`
	StackTrace ErrorData `json:"stackTrace"`
}

type DataResponse[T any] struct {
	Data T      `json:"data"`
	Date string `json:"datetime"`
}

func SuccessArray[T any](data []T, msg string, totalCnt int) ArrayResponse[T] {
	var response = make([]T, 0)
	if data != nil {
		response = data
	}
	return ArrayResponse[T]{
		Message:    msg,
		Data:       response,
		ItemCount:  len(data),
		TotalCount: totalCnt,
	}
}

func CreateFailure(err types.CustomError) ErrorData {
	return ErrorData{
		Error:  err.Err,
		Source: err.Source,
		Date:   DateTime(err.DateTime, constants.DateTimeResponseFormat),
	}
}

func CreateFailureWithMessage(err types.CustomError) ErrorWithMessage {
	errMsg := err.Message
	if errMsg == "" {
		errMsg = err.Err.Error()
	}
	return ErrorWithMessage{
		Message:    errMsg,
		StackTrace: CreateFailure(err),
	}
}

func SingleObject[T any](data T) DataResponse[T] {
	return DataResponse[T]{
		Date: DateTime(time.Now(), constants.DateTimeMinSecResponseFormat),
		Data: data,
	}
}
