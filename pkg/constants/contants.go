package constants

import (
	"fmt"
	"gopher-camp/pkg/env"
)

var DbName = "DATABASE_NAME"
var DbHost = "DATABASE_HOST"
var DbPort = "DATABASE_PORT"
var DbUser = "DATABASE_USER"
var DbPassword = "DATABASE_PASS"
var URIPort = "URL_PORT"

func DatabaseURI() string {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		env.GetEnv(DbUser),
		env.GetEnv(DbPassword),
		env.GetEnv(DbHost),
		env.GetEnv(DbPort),
		env.GetEnv(DbName))
	return url
}

func HostPort() string {
	return env.GetEnvOrDefault(URIPort, ":8008")
}

var DateResponseFormat = "Jan 2, 2006"
var DateTimeResponseFormat = "Jan 2, 2006 15:04"
var DateTimeMinSecResponseFormat = "Jan 2, 2006 15:04:05"
