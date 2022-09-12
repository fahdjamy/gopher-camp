package utils

type ArrayResponse[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

func SuccessArrayResponse[T any](data []T, msg string) ArrayResponse[T] {
	var response = make([]T, 0)
	if data != nil {
		response = data
	}
	return ArrayResponse[T]{
		Message: msg,
		Data:    response,
	}
}
