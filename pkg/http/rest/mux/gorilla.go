package mux

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	Router *mux.Router
}

var (
	getMethod    = "GET"
	putMethod    = "PUT"
	postMethod   = "POST"
	deleteMethod = "DELETE"
)

func (m Server) Get(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, handler).Methods(getMethod)
}

func (m Server) Put(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, handler).Methods(putMethod)
}

func (m Server) Post(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, handler).Methods(postMethod)
}

func (m Server) Delete(path string, handler http.HandlerFunc) {
	m.Router.HandleFunc(path, handler).Methods(deleteMethod)
}

func NewMuxServer(address string, readTimeOut time.Duration, writeTimeOut time.Duration) (Server, *http.Server) {
	router := mux.NewRouter().StrictSlash(true)
	muxSvr := Server{
		Router: router,
	}
	httpSrv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: writeTimeOut,
		ReadTimeout:  readTimeOut,
	}
	return muxSvr, httpSrv
}
