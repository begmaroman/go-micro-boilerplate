package healthchecker

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/begmaroman/go-micro-boilerplate/proto/health"
)

// Wrap transforms a RPC healthcheck into a Check
func WrapRPC(check func(_ context.Context, _ *emptypb.Empty, res *health.HealthResponse) error) Check {
	return func() error {
		var res health.HealthResponse
		if err := check(context.TODO(), nil, &res); err != nil {
			return err
		}

		if res.GetStatus() != health.HealthResponse_SERVING {
			return fmt.Errorf("unhealthy: %v", res.GetMessage())
		}

		return nil
	}
}
