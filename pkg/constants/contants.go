package constants

import (
	"gopher-camp/pkg/env"
)

var DBName = "DATABASE_NAME"
var DBHost = "DATABASE_HOST"
var DBSSLMode = "DATABASE_SSL_MODE"
var DBPort = "DATABASE_PORT"
var DBUser = "DATABASE_USER"
var DBPassword = "DATABASE_PASS"
var URIPort = "URL_PORT"

func HostPort() string {
	return env.GetEnvOrDefault(URIPort, ":8008")
}

var DateResponseFormat = "Jan 2, 2006"
var DateTimeResponseFormat = "Jan 2, 2006 15:04"
var DateTimeMinSecResponseFormat = "Jan 2, 2006 15:04:05"

type LogLevel uint

const (
	ErrorLevel LogLevel = iota
	InfoLevel
	TraceLevel
	WarningLevel
)
