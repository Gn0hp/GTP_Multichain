package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	logrusAdapter "logur.dev/adapter/logrus"
	"logur.dev/logur"
	"os"
)

func NewLogger(config Config) logur.LoggerFacade {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	switch config.Format {
	case "logfmt":
		logger.SetFormatter(&logrus.TextFormatter{
			DisableColors:             config.NoColor,
			EnvironmentOverrideColors: true,
			FullTimestamp:             config.FullTimestamp,
		})
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: true,
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{})
	}
	if level, err := logrus.ParseLevel(config.Level); err == nil {
		logger.SetLevel(level)
	}

	return logrusAdapter.New(logger)
}

// SetStandartLogger sets the standard logger to a custom logger instance.
func SetStandartLogger(logger logur.Logger) {
	log.SetOutput(logur.NewLevelWriter(logger, logur.Info))
}
