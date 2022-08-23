package env

import "os"

func GetEnv(name string) string {
	return os.Getenv(name)
}

func SetUpEnvironment() {
	os.Setenv("host", "localhost")
	os.Setenv("port", "5234")
}
