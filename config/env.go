package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

var (
	GlobalConfig Config
)

type Config struct {
	Kubernetes struct {
		Timeout              int      `envconfig:"default=180"`
		DiscoveryBlackGroups []string `envconfig:"optional"`
		DisableCache         bool     `envconfig:"default=false"`
	}

	ClientGo struct {
		Burst int     `envconfig:"default=100"`
		QPS   float32 `envconfig:"default=100.0"`
	}

	Furion struct {
		Endpoint   string `envconfig:"default=http://mock-server:80"`
		APIVersion string `envconfig:"default=v1"`
		Timeout    int    `envconfig:"default=3"`
	}

	Saas struct {
		Endpoint string `envconfig:"default=http://10.200.108.54:8080"`
		Timeout  int    `envconfig:"default=3"`
	}
}

func Initconfig() {
	err := envconfig.Init(&GlobalConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
}
