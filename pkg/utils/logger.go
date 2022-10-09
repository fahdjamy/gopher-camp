package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopher-camp/pkg/constants"
	"gopher-camp/pkg/types"
)

type CustomLogger struct{}

var logger = logrus.New()

func (l *CustomLogger) Warn(err error) {
	logger.Warnln(err)
}

func (l *CustomLogger) Trace(err error) {
	logger.Traceln(err)
}

func (l *CustomLogger) Fatal(err error) {
	logger.Fatalln(err)
}

func (l *CustomLogger) Error(err error) {
	logger.Errorln(err)
}

func (l *CustomLogger) Debug(err error) {
	logger.Debugln(err)
}

func (l *CustomLogger) Info(msg string) {
	logger.Info(msg)
}

func (l *CustomLogger) CustomError(err types.CustomError) {
	logMsg := fmt.Sprintf("source=%v, error=%v, date=%v",
		err.Source,
		err,
		err.DateTime,
	)
	logger.Error(logMsg)
}

func NewCustomLogger(level constants.LogLevel) *CustomLogger {
	//logger.SetLevel(logrus.Level(level))
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetFormatter(&logrus.TextFormatter{
		//DisableColors: true,
		FullTimestamp: true,
	})

	return &CustomLogger{}
}
