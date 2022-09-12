package types

import "gopher-camp/pkg/models"

type DTOMapper[T models.Model, R any] interface {
	GetDTO() R
	MapToDO(domain *T) *T
}
