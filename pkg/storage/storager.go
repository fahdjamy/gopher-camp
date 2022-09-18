package storage

import (
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
)

type Storage[T models.Model, R any, V any] interface {
	FindAll() []V
	Delete(id int) (bool, error)
	FindById(id int) (V, error)
	Create(model types.DTOMapper[T, R]) (V, error)
	Update(id int, model types.DTOMapper[T, R]) (V, error)
}
