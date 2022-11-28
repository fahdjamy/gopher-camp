package controllers

import (
	"encoding/json"
	"net/http"
	"profiler/pkg/types"
	"profiler/pkg/utils"
)

func writeJSONSuccessResponse(w http.ResponseWriter, _ *http.Request, response []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(response)
	if err != nil {
		panic(err)
	}
}

func writeFailureResponse(w http.ResponseWriter, err error, status int) {
	res, _ := json.Marshal(utils.CreateFailureWithMessage(err.(types.CustomError)))
	_, _ = w.Write(res)
	w.WriteHeader(status)
}
