package main

import (
	// We need these imports to register NATS broker, registry, and transport.
	// This type is defined through evars in docker-compose.yaml file.
	_ "github.com/micro/plugins/v5/broker/nats"
	_ "github.com/micro/plugins/v5/registry/nats"
	_ "github.com/micro/plugins/v5/transport/nats"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v5/selector"

	"github.com/begmaroman/go-micro-boilerplate/pkg/rpc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/microservice"
)

// Version may be changed during build via --ldflags parameter
var Version = "latest"

func init() {
	// Set the default selector strategy to the custom RoundRobin.
	selector.DefaultSelector = selector.NewSelector(selector.SetStrategy(rpc.Random))
}

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
