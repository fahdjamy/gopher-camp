package constants

import (
	"fmt"
	"gopher-camp/pkg/services/env"
)

var DbName = "DATABASE_NAME"
var DbHost = "DATABASE_HOST"
var DbPort = "DATABASE_PORT"
var DbUser = "DATABASE_USER"
var DbPassword = "DATABASE_PASS"

func DatabaseURI() string {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		env.GetEnv(DbUser),
		env.GetEnv(DbPassword),
		env.GetEnv(DbHost),
		env.GetEnv(DbPort),
		env.GetEnv(DbName))
	return url
}
