// This package contains business logic of the service.
// No store or transport logic inside.
package domain

import (
	"context"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
)

// Domain represents the behavior of the business/domain logic of the service.
// Currently, proto models are used in the domain layer.
// This is not really the right way because proto models are more about the transport layer, not the domain one.
// So if you would like to implement domain representation of proto models, you can do so.
// Don't forget to implement a conversion logic between proto models and domain ones :)
type Domain interface {
	// HealthCheck returns an error if there is a problem with the service.
	HealthCheck() error

	// CreateUser creates a new user by the given input.
	CreateUser(context.Context, *accountproto.User) (*accountproto.User, error)

	// ReadUser reads an existing user by its ID.
	ReadUser(context.Context, string) (*accountproto.User, error)

	// ListUsers lists all users.
	ListUsers(context.Context) ([]*accountproto.User, error)

	// UpdateUser updates an existing user by its ID using the given input.
	UpdateUser(context.Context, string, *accountproto.User) (*accountproto.User, error)

	// DeleteUser deletes an existing user by its ID.
	DeleteUser(context.Context, string) error
}
