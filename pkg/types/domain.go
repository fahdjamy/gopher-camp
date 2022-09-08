package types

import (
	"gopher-camp/pkg/config/database"
)

type Domain[T any] interface {
	Create(db database.Database) *T
	FindAll(db database.Database) []T
	FindById(id int64, db database.Database) *T
	DeleteById(id int64, db database.Database) T
}
