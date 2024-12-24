package rpc

import (
	"math/rand/v2"
	"strings"

	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/selector"
)

// Random is a randomized strategy algorithm for node selection.
// Fixes the issue described here: https://github.com/micro/go-micro/issues/2741
func Random(services []*registry.Service) selector.Next {
	nodes := make([]*registry.Node, 0, len(services))

	for _, service := range services {
		for _, node := range service.Nodes {
			for _, sn := range AllServiceNames {
				if strings.HasPrefix(node.Id, sn) {
					nodes = append(nodes, node)
				}
			}
		}
	}

	return func() (*registry.Node, error) {
		if len(nodes) == 0 {
			return nil, selector.ErrNoneAvailable
		}

		i := rand.Int() % len(nodes)
		return nodes[i], nil
	}
}
