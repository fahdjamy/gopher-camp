package types

type Logger interface {
	LogInfo(msg string, src string, pkg string)
	LogError(err error, src string, pkg string)
}
