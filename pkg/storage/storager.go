package storage

import (
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
)

type Storage[T models.Model, R any] interface {
	FindAll() []T
	Delete(id int) (bool, error)
	FindById(id int) (*T, error)
	Create(model types.DTOMapper[T, R]) (*T, error)
	Update(id int, model types.DTOMapper[T, R]) (*T, error)
}
