package utils

import (
	"fmt"
	"log"
	"time"
)

type CustomLogger struct{}

var dateNow = time.Now()

func (l CustomLogger) LogInfo(msg string, src string, pkg string) {
	logMsg := fmt.Sprintf("source=%v, message=%v, pakage=%v, date=%v",
		src,
		msg,
		pkg,
		dateNow,
	)
	log.Println(logMsg)
}

func (l CustomLogger) LogError(err error, src string, pkg string) {
	logMsg := fmt.Sprintf("source=%v, error=%v, file=%v, date=%v",
		src,
		err,
		pkg,
		dateNow,
	)
	log.Println(logMsg)
}

func NewLogger() CustomLogger {
	return CustomLogger{}
}
