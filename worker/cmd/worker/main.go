package main

import (
	log "github.com/sirupsen/logrus"
	"worker/internal/api/routing"
	"worker/internal/common/config"
)

func main() {
	config.ConfigureLogging()

	handler := routing.SetUpRouting()
	if err := routing.Run(handler); err != nil {
		log.Panicf("Failed to start the server: %v", err)
	}
}
