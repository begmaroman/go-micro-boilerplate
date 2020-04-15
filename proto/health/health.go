package health

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"

	"github.com/begmaroman/go-micro-boilerplate/utils/healthchecker"
)

// Wrap transforms a RPC healthcheck into a healthchecker.Check
func Wrap(check func(_ context.Context, _ *empty.Empty, res *HealthResponse) error) healthchecker.Check {
	return func() error {
		var res HealthResponse
		if err := check(context.TODO(), nil, &res); err != nil {
			return err
		}

		if res.GetStatus() != HealthResponse_SERVING {
			return fmt.Errorf("Unhealthy: %v", res.GetMessage())
		}

		return nil
	}
}

// SelectNodeByName returns a function similar to selector.Random that will select the node by the given name
func SelectNodeByName(serviceName, nodeID string) func([]*registry.Service) selector.Next {
	return func(services []*registry.Service) selector.Next {
		var node *registry.Node

		for _, service := range services {
			if service.Name != serviceName {
				continue
			}

			for _, n := range service.Nodes {
				if n.Id == nodeID {
					node = n
					break // We've found the node, no need to explore the rest
				}
			}
		}
		return func() (*registry.Node, error) {
			if node == nil {
				return nil, selector.ErrNoneAvailable
			}
			return node, nil
		}
	}
}
