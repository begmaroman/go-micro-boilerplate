package health

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
)

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
