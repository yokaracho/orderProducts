package logger

import "github.com/sirupsen/logrus"

func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	return logger
}
