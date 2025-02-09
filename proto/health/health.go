package health

import (
	"go-micro.dev/v5"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/selector"
)

// SelectNodeByName returns a function similar to selector.Random that will select the node by the given name
func SelectNodeByName(svc micro.Service) func([]*registry.Service) selector.Next {
	serviceName := svc.Server().Options().Name
	nodeID := svc.Server().Options().Id

	return func(services []*registry.Service) selector.Next {
		var node *registry.Node

		for _, service := range services {
			if service.Name != serviceName {
				continue
			}

			for _, n := range service.Nodes {
				if n.Id == serviceName+"-"+nodeID {
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
