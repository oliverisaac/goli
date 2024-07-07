package goli

import (
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
)

// InitLogrus initializes logrus using env vars
func InitLogrus(defaultLevel logrus.Level) {
	ll := DefaultEnv("LOG_LEVEL", defaultLevel.String())
	logLevel, err := logrus.ParseLevel(ll)
	if err != nil {
		logrus.Errorf("Failed to parse LOG_LEVEL=%s", ll)
		logLevel = defaultLevel
	}
	logrus.SetLevel(logLevel)

	logFormat := DefaultEnv("LOG_FORMAT", "text")
	if logFormat == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
