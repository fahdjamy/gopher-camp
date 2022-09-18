package types

import (
	"gopher-camp/pkg/models"
	"net/http"
)

type Domain interface {
	Me() string
	Validate() error
}

type Logger interface {
	LogInfo(msg string, src string, pkg string)
	LogError(err error, src string, pkg string)
}

type DTOMapper[T models.Model, R any] interface {
	GetDTO() R
	MapToDO(domain *T) *T
}

type RestRouters interface {
	Get(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
}

type Storage[T models.Model, R any, V any] interface {
	FindAll() []V
	Delete(id int) (bool, error)
	FindById(id int) (V, error)
	Create(model DTOMapper[T, R]) (V, error)
	Update(id int, model DTOMapper[T, R]) (V, error)
}
