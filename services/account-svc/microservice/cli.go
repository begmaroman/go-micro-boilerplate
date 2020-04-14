package microservice

import (
	"github.com/micro/cli"
)

// flags contains the list of configuration parameters
var flags = []cli.Flag{
	cli.BoolFlag{
		Name:   "docker_compose",
		EnvVar: "DOCKER_COMPOSE",
		Usage:  "Set to true if we are running in docker-compose",
	},
}

// buildOptions builds the buildOptions based on the cli context
func buildOptions(c *cli.Context) *Options {
	return &Options{
		IsTest: c.Bool("docker_compose"),
	}
}
