package storage

import "gopher-camp/pkg/types"

type Storage[T any] interface {
	FindAll() []types.Domain[T]
	Delete(id int) (bool, error)
	FindById(id int) (*types.Domain[T], error)
	Create(model types.Domain[T]) (*types.Domain[T], error)
	Update(id int, model types.Domain[T]) (*types.Domain[T], error)
}
