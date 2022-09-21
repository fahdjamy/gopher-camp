package types

import (
	"gopher-camp/pkg/models"
	"net/http"
)

type Domain interface {
	ToString() string
	Validate() error
}

type Logger interface {
	LogError(err CustomError)
	LogInfo(msg string, src string, pkg string)
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

// DOServiceProvider stands for Domain Object Service Provider
type DOServiceProvider[T models.Model] interface {
	FindAll() []T
	Delete(id uint) (bool, error)
	FindById(id uint) (*T, error)
	Create(model *T) (*T, error)
	Update(id uint, model *T) (*T, error)
}

type StorageProvider[T models.Model] interface {
	DeleteOne(table string, domainID int) error
	DeleteMany(table string, domainID int) error
	Update(table string, domain Domain) *T
	FindOne(table string, criteria string) *T
	InsertOne(table string, domain Domain) *T
	FindMany(table string, criteria string) []*T
	InsertMany(table string, domains ...Domain) []*T
}
