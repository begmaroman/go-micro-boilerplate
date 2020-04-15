// This package contains the store representation of this service.
// No business logic inside of this package.
package store

import (
	"context"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
)

// Store represents the behavior of the store layer.
// Currently, proto models are used in the store layer as well.
// The same comment as for domain.Domain interface.
type Store interface {
	// CreateUser creates a new user by the given input in the store.
	// This function only creates a new record using the given input. No business logic there.
	CreateUser(context.Context, *accountproto.User) (*accountproto.User, error)

	// ReadUser reads an existing user by its ID from the store.
	ReadUser(context.Context, string) (*accountproto.User, error)

	// ListUsers lists all users from the store.
	ListUsers(context.Context) ([]*accountproto.User, error)

	// UpdateUser updates an existing user in the store by its ID using the given input.
	// This function only updates the record using the given input. No business logic there.
	UpdateUser(context.Context, string, *accountproto.User) (*accountproto.User, error)

	// DeleteUser deletes an existing user from the store by its ID.
	// This function only deletes the record using the given input. No business logic there.
	DeleteUser(context.Context, string) error
}
