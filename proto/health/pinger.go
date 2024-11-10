package health

import (
	"context"

	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Pinger is the interface implemented by the client of RPC services that can ping themselves
type Pinger interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error)
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
	_, err := c.pinger.Ping(ctx, &emptypb.Empty{}, c.selfCallOption)
	return err
}
