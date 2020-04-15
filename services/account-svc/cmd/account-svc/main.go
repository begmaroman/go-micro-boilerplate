package main

import (
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
	"github.com/sirupsen/logrus"

	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/microservice"
)

// Version is set during build via --ldflags parameter
var Version = "latest"

func main() {
	logger := logrus.New()

	// Initialize service
	microService, err := microservice.Init(&microservice.ClientOptions{
		Version: Version,
		Log:     logger,
	})
	if err != nil {
		logger.WithError(err).Fatal("failed to initialize micro-service")
	}

	// Run service
	if err := microService.Run(); err != nil {
		logger.WithError(err).Fatal("failed to run micro-service")
	}
}
