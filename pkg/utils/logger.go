package utils

import (
	"fmt"
	"gopher-camp/pkg/types"
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

func (l CustomLogger) LogError(err types.CustomError) {
	logMsg := fmt.Sprintf("source=%v, error=%v, date=%v",
		err.Source,
		err,
		err.DateTime,
	)
	log.Println(logMsg)
}

func NewCustomLogger() CustomLogger {
	return CustomLogger{}
}
