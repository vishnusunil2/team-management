package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	ApiLogger *logrus.Logger
)

func Init(logLevel string) {
	level, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		logrus.Fatalf("Invalid log level: %v ", err)
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
func LogMessage(level logrus.Level, format string, data map[string]interface{}, args ...interface{}) {
	entry := logrus.WithFields(logrus.Fields(data))
	if len(args) > 0 {
		entry.Logf(level, format, args...)
	} else {
		entry.Log(level, format)
	}
}
func Debugf(format string, data map[string]interface{}, args ...interface{}) {
	LogMessage(logrus.DebugLevel, format, data, args...)
}

func Infof(format string, args ...interface{}) {
	LogMessage(logrus.InfoLevel, format, nil, args...)
}

func Errorf(format string, args ...interface{}) {
	LogMessage(logrus.ErrorLevel, format, nil, args...)
}

func Fatalf(format string, args ...interface{}) {
	LogMessage(logrus.FatalLevel, format, nil, args...)
}
