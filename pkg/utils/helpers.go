package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, data interface{}) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(body, data)
	}

	if err != nil {
		log.Println("utils: Parsing body = ", err)
	}
}
