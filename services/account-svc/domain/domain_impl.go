package domain

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/store"
)

// Options contains options to create a domain.
type Options struct {
	Store store.Store
	Log   *logrus.Logger
}

// domain implements the business/domain logic of the service.
type domain struct {
	store store.Store
	log   *logrus.Logger
}

// New is the constructor of domain.
func New(opts *Options) Domain {
	return &domain{
		store: opts.Store,
		log:   opts.Log,
	}
}

// HealthCheck implements Domain interface.
func (d *domain) HealthCheck() error {
	return nil
}

// CreateUser implements Domain interface.
// The business logic of the user creation operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *domain) CreateUser(ctx context.Context, input *accountproto.User) (*accountproto.User, error) {
	// Call the store directly.
	createdUser, err := d.store.CreateUser(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create user in the store layer")
	}

	return createdUser, nil
}

// ReadUser implements Domain interface.
// The business logic of the user reading operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *domain) ReadUser(ctx context.Context, id string) (*accountproto.User, error) {
	// Call the store directly.
	user, err := d.store.ReadUser(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read user in the store layer with ID '%s'", id)
	}

	return user, nil
}

// ListUsers implements Domain interface.
// The business logic of the user listing operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *domain) ListUsers(ctx context.Context) ([]*accountproto.User, error) {
	// Call the store directly.
	users, err := d.store.ListUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list users in the store layer")
	}

	return users, nil
}

// UpdateUser implements Domain interface.
// The business logic of the user updating operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *domain) UpdateUser(ctx context.Context, id string, input *accountproto.User) (*accountproto.User, error) {
	// Call the store directly.
	updatedUser, err := d.store.UpdateUser(ctx, id, input)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to update user in the store layer with ID '%s'", id)
	}

	return updatedUser, nil
}

// DeleteUser implements Domain interface.
// The business logic of the user deletion operation can be implemented within this function.
// For now, it's not implemented because this is just an example of an architecture.
func (d *domain) DeleteUser(ctx context.Context, id string) error {
	// Call the store directly.
	if err := d.store.DeleteUser(ctx, id); err != nil {
		return errors.Wrapf(err, "unable to delete user in the store layer with ID '%s'", id)
	}

	return nil
}
