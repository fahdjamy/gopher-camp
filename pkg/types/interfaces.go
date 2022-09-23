package types

import (
	"gopher-camp/pkg/models"
	"net/http"
)

type Domain[T Model] interface {
	ToString() string
	Validate() error
	Me() *T
}

type Logger interface {
	LogError(err CustomError)
	LogInfo(msg string, src string, pkg string)
}

type DTOMapper[T Model, R any] interface {
	GetDTO() R
	MapToDO(domain *T) *T
}

type RestRouters interface {
	Get(path string, handler http.HandlerFunc)
	Put(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
	Delete(path string, handler http.HandlerFunc)
}

// DOServiceProvider stands for Domain Object Service Provider
type DOServiceProvider[T Model] interface {
	FindAll() []T
	Delete(id uint) (bool, error)
	FindById(id uint) (*T, error)
	Create(model *T) (*T, error)
	Update(id uint, model *T) (*T, error)
}

type StorageProvider[T Model] interface {
	DeleteOne(table string, id int) error
	DeleteMany(table string, id int) error
	Update(table string, domain Domain[T]) *T
	FindOne(table string, criteria string) *T
	InsertOne(table string, domain Domain[T]) *T
	FindMany(table string, criteria string) []*T
	InsertMany(table string, domains ...Domain[T]) []*T
}

type Model interface {
	models.Project | models.Company | models.Founder
}
