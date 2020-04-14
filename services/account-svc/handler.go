package accountsvc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/proto/health"
	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/domain"
)

// To make sure Handler implements accountproto.AccountService interface
var _ accountproto.AccountServiceHandler = &Handler{}

// Options serves as the dependency injection container to create a new handler
type Options struct {
	Service        domain.Domain
	SelfPingClient *health.SelfPingClient
	Log            *logrus.Logger
}

// Handler implements authzproto.AuthorizationServiceHandler interface
type Handler struct {
	service        domain.Domain
	selfPingClient *health.SelfPingClient
	log            *logrus.Logger
}

// NewHandler returns a new handler for the account-svc
func NewHandler(opts *Options) *Handler {
	return &Handler{
		service:        opts.Service,
		selfPingClient: opts.SelfPingClient,
		log:            opts.Log,
	}
}

// CreateUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to create a new user by the given input.
func (h *Handler) CreateUser(ctx context.Context, req *accountproto.CreateUserRequest, resp *accountproto.CreateUserResponse) error {

}

// ReadUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to read an existing user by the given ID.
func (h *Handler) ReadUser(ctx context.Context, req *accountproto.ReadUserRequest, resp *accountproto.ReadUserResponse) error {

}

// ListUsers implements accountproto.AccountServiceHandler interface.
// Calls the service's method to list all users.
func (h *Handler) ListUsers(ctx context.Context, req *accountproto.ListUsersRequest, resp *accountproto.ListUsersResponse) error {

}

// UpdateUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to update an existing user by the given ID and input.
func (h *Handler) UpdateUser(ctx context.Context, req *accountproto.UpdateUserRequest, resp *accountproto.UpdateUserResponse) error {

}

// DeleteUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to delete an existing user by the given ID.
func (h *Handler) DeleteUser(ctx context.Context, req *accountproto.DeleteUserRequest, resp *accountproto.DeleteUserResponse) error {

}

// Health implements accountproto.AccountServiceHandler interface
func (h *Handler) Health(ctx context.Context, _ *empty.Empty, res *health.HealthResponse) error {
	// Check database
	if err := h.service.HealthCheck(); err != nil {
		res.Status = health.HealthResponse_NOT_SERVING
		return err
	}

	// Check nats, i.e. call ourselves (exactly this node) via nats
	err := h.selfPingClient.Ping(ctx)
	if err != nil {
		res.Status = health.HealthResponse_NOT_SERVING
		return errors.Wrapf(err, "unable to ping ourselves")
	}

	res.Status = health.HealthResponse_SERVING
	return nil
}

// Ping implements accountproto.AccountServiceHandler interface
func (h *Handler) Ping(ctx context.Context, _ *empty.Empty, _ *empty.Empty) error {
	return nil
}
