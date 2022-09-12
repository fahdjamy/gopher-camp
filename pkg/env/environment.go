package env

import (
	"os"
)

func GetEnv(name string) string {
	return os.Getenv(name)
}

func GetEnvOrDefault(port, fallback string) string {
	hostPort := os.Getenv(port)
	if len(hostPort) == 0 {
		return fallback
	}
	return hostPort
}
