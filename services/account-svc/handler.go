// This is the transport layer of the service. Here we use RPC communication type.
// You can implement any transport you want, e.g. HTTP, WS, etc.
package accountsvc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/proto/health"
	proto "github.com/begmaroman/go-micro-boilerplate/proto/status"
	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/domain"
	"github.com/begmaroman/go-micro-boilerplate/utils/rpc"
)

// To make sure Handler implements accountproto.AccountService interface.
var _ accountproto.AccountServiceHandler = &Handler{}

// Options serves as the dependency injection container to create a new handler.
type Options struct {
	Service        domain.Domain
	SelfPingClient *health.SelfPingClient
	Log            *logrus.Logger
}

// Handler implements authzproto.AuthorizationServiceHandler interface.
type Handler struct {
	service        domain.Domain
	selfPingClient *health.SelfPingClient
	log            *logrus.Logger
}

// NewHandler returns a new handler for the account-svc.
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
	// Create user.
	createdUser, err := h.service.CreateUser(ctx, req.GetUser())
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &accountproto.CreateUserResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrap(err, "unable to create user")
	}

	// Prepare RPC response data.
	resp.Result = &accountproto.CreateUserResponse_User{
		User: createdUser,
	}
	return nil
}

// ReadUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to read an existing user by the given ID.
func (h *Handler) ReadUser(ctx context.Context, req *accountproto.ReadUserRequest, resp *accountproto.ReadUserResponse) error {
	// Read user.
	user, err := h.service.ReadUser(ctx, req.GetUserId())
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &accountproto.ReadUserResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to read user with ID '%s'", req.GetUserId())
	}

	// Prepare RPC response data.
	resp.Result = &accountproto.ReadUserResponse_User{
		User: user,
	}
	return nil
}

// ListUsers implements accountproto.AccountServiceHandler interface.
// Calls the service's method to list all users.
func (h *Handler) ListUsers(ctx context.Context, req *accountproto.ListUsersRequest, resp *accountproto.ListUsersResponse) error {
	// List all users.
	users, err := h.service.ListUsers(ctx)
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &accountproto.ListUsersResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to list users")
	}

	// Prepare RPC response data.
	resp.Result = &accountproto.ListUsersResponse_Data{
		Data: &accountproto.ListUsersResponseOK{
			Users: users,
		},
	}
	return nil
}

// UpdateUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to update an existing user by the given ID and input.
func (h *Handler) UpdateUser(ctx context.Context, req *accountproto.UpdateUserRequest, resp *accountproto.UpdateUserResponse) error {
	// Update user by its ID.
	user, err := h.service.UpdateUser(ctx, req.GetUserId(), req.GetUser())
	if err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &accountproto.UpdateUserResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to update user with ID '%s'", req.GetUserId())
	}

	// Prepare RPC response data.
	resp.Result = &accountproto.UpdateUserResponse_User{
		User: user,
	}
	return nil
}

// DeleteUser implements accountproto.AccountServiceHandler interface.
// Calls the service's method to delete an existing user by the given ID.
func (h *Handler) DeleteUser(ctx context.Context, req *accountproto.DeleteUserRequest, resp *accountproto.DeleteUserResponse) error {
	// Delete user by its ID.
	if err := h.service.DeleteUser(ctx, req.GetUserId()); err != nil {
		// Try to convert the given error to the proto status.
		if resStatus, ok := h.errorAsStatus(ctx, err); ok {
			resp.Result = &accountproto.DeleteUserResponse_Error{
				Error: resStatus,
			}
			return nil
		}

		// Otherwise just return this error wrapped to a description.
		return errors.Wrapf(err, "unable to delete user with ID '%s'", req.GetUserId())
	}

	// Prepare RPC response data.
	resp.Result = &accountproto.DeleteUserResponse_Empty{
		Empty: &empty.Empty{},
	}
	return nil
}

// Health implements accountproto.AccountServiceHandler interface
func (h *Handler) Health(ctx context.Context, _ *empty.Empty, res *health.HealthResponse) error {
	// Check database
	if err := h.service.HealthCheck(); err != nil {
		res.Status = health.HealthResponse_NOT_SERVING
		return err
	}

	// Check nats, i.e. call ourselves (exactly this node) via nats
	if err := h.selfPingClient.Ping(ctx); err != nil {
		res.Status = health.HealthResponse_NOT_SERVING
		return errors.Wrapf(err, "unable to ping ourselves")
	}

	res.Status = health.HealthResponse_SERVING
	return nil
}

// Ping implements accountproto.AccountServiceHandler and helath.Pinger interface.
// This is needed to implement self-pinger functionality.
func (h *Handler) Ping(ctx context.Context, _ *empty.Empty, _ *empty.Empty) error {
	return nil
}

// errorAsStatus converts the given error to the proto status.
// This function have to be implemented according to the logic of your project.
// For now, it always returns the ErrAborted RPC status code.
// What will be returned:
// - the first parameter if the proto status of the error;
// - the second boolean value is true, if the error has been matched with one of RPC statuses;
func (h *Handler) errorAsStatus(ctx context.Context, err error) (*proto.Status, bool) {
	return rpc.ErrAbortedf(err.Error()), true
}
