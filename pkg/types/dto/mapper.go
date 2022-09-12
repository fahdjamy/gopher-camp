package dto

import "gopher-camp/pkg/models"

type ToDomainMapper[T models.Model] interface {
	MapToDO(domain *T) *T
}
