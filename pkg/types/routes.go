package types

import "net/http"

type RestRouters interface {
	Get(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
}
