package config

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func ConfigureLogging() {
	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
}
