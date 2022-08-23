package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, data interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(body, data)
	}

	if err != nil {
		log.Println("utils: Parsing body = ", err)
	}
}
