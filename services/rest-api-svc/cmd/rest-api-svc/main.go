package main

import (
	// We need these imports to register NATS broker, registry, and transport.
	// This type is defined through evars in docker-compose.yaml file.
	_ "github.com/micro/go-plugins/broker/nats/v2"
	_ "github.com/micro/go-plugins/registry/nats/v2"
	_ "github.com/micro/go-plugins/transport/nats/v2"
	"github.com/sirupsen/logrus"

	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/microservice"
)

// Version may be changed during build via --ldflags parameter
var Version = "latest"

func main() {
	logger := logrus.New()

	// Initialize service.
	microService, err := microservice.Init(&microservice.ClientOptions{
		Version: Version,
		Log:     logger,
	})
	if err != nil {
		logger.WithError(err).Fatal("failed to initialize web-service")
	}

	// Run service.
	if err := microService.Run(); err != nil {
		logger.WithError(err).Fatal("failed to run web-service")
	}
}
