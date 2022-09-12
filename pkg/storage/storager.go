package storage

import (
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types/dto"
)

type Storage[T models.Model] interface {
	FindAll() []T
	Delete(id int) (bool, error)
	FindById(id int) (*T, error)
	Create(model dto.ToDomainMapper[T]) (*T, error)
	Update(id int, model dto.ToDomainMapper[T]) (*T, error)
}
