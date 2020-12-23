package microservice

import (
	"github.com/micro/cli/v2"
)

var opts Options

// flags contains the list of configuration parameters.
var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:        "docker_compose",
		EnvVars:     []string{"DOCKER_COMPOSE"},
		Usage:       "Set to true if we are running in docker-compose",
		Destination: &opts.IsTest,
	},
}
