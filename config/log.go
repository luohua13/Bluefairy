package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetLogger() {
	f, err := os.OpenFile("/var/log/bulefairy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)
}
