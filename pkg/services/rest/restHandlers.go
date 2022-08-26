package rest

import "net/http"

type Routers interface {
	Get(path string, handler http.Handler)
	Put(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Delete(path string, handler http.Handler)
}
