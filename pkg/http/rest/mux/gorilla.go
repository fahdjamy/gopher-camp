package mux

import (
	"net/http"
	"time"
)

type mux struct {
	Addr         string
	ReadTimeout  time.Time
	WriteTimeout time.Time
}

func (m mux) Get(path string, handler http.Handler) {
	//TODO implement me
	panic("implement me")
}

func (m mux) Put(path string, handler http.Handler) {
	//TODO implement me
	panic("implement me")
}

func (m mux) Post(path string, handler http.Handler) {
	//TODO implement me
	panic("implement me")
}

func (m mux) Delete(path string, handler http.Handler) {
	//TODO implement me
	panic("implement me")
}

func NewMux() *mux {
	return &mux{}
}
