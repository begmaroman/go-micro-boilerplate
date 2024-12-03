package main

import (
	// We need these imports to register NATS broker, registry, and transport.
	// This type is defined through evars in docker-compose.yaml file.
	_ "github.com/micro/plugins/v5/broker/nats"
	_ "github.com/micro/plugins/v5/registry/nats"
	_ "github.com/micro/plugins/v5/transport/nats"
)
