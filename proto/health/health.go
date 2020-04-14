package health

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"golang.org/x/net/context"

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

// Pinger is the interface implemented by the client of RPC services that can ping themselves
type Pinger interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*empty.Empty, error)
}

// SelfPingClient wraps the functionality required to call our node
type SelfPingClient struct {
	pinger         Pinger
	selfCallOption client.CallOption
}

// NewSelfPingClient create a new SelfPingClient
func NewSelfPingClient(service micro.Service, pinger Pinger) *SelfPingClient {
	serviceName := service.Server().Options().Name
	nodeID := service.Server().Options().Id

	return &SelfPingClient{
		pinger: pinger,
		selfCallOption: client.WithSelectOption(
			selector.WithStrategy(SelectNodeByName(serviceName, serviceName+"-"+nodeID)),
		),
	}
}

// Ping uses the SelfPingClient to ping the current node
func (c *SelfPingClient) Ping(ctx context.Context) error {
	_, err := c.pinger.Ping(ctx, &empty.Empty{}, c.selfCallOption)
	return err
}
