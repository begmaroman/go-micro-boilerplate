// This package implements the store layer using in-memory data store.
// No business logic inside this package, only CRUD operations.
package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/store"
)

// Options contains the options to create a memory store
type Options struct {
	Log *logrus.Logger
}

// memory implements store.Store interface.
// Represents the store logic using in-memory data store.
type memory struct {
	sync.Mutex

	data map[string]*accountproto.User
	log  *logrus.Logger
}

// New is the constructor of memory
func New(opts *Options) store.Store {
	return &memory{
		data: make(map[string]*accountproto.User),
		log:  opts.Log,
	}
}

// CreateUser implements store.Store interface.
// This function stores the given user.
func (m *memory) CreateUser(ctx context.Context, input *accountproto.User) (*accountproto.User, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Generate a new user ID.
	input.Id = uuid.New()

	// Set timestamps
	now := ptypes.TimestampNow()
	input.CreatedAt = now
	input.UpdatedAt = now

	// Store the user
	m.data[input.Id] = input

	return input, nil
}

// ReadUser implements store.Store interface.
// This function reads an existing user by its ID.
func (m *memory) ReadUser(ctx context.Context, id string) (*accountproto.User, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Retrieve user with the given ID.
	user, ok := m.data[id]
	if !ok {
		// Return the not found errors.
		// Here should be custom not found error implementation
		// to convert in to the proto status instead of return this error.
		return nil, fmt.Errorf("user with ID '%s' doesn't found", id)
	}

	return user, nil
}

// ListUsers implements store.Store interface.
// This function lists all users.
func (m *memory) ListUsers(ctx context.Context) ([]*accountproto.User, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Prepare data to return.
	var users []*accountproto.User
	for _, user := range m.data {
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser implements store.Store interface.
// This function updates an existing user by its ID.
func (m *memory) UpdateUser(ctx context.Context, id string, input *accountproto.User) (*accountproto.User, error) {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Retrieve user with the given ID.
	if _, ok := m.data[id]; !ok {
		// Return the not found errors.
		// Here should be custom not found error implementation
		// to convert in to the proto status instead of return this error.
		return nil, fmt.Errorf("user with ID '%s' doesn't found", id)
	}

	// Update user record.
	input.UpdatedAt = ptypes.TimestampNow()
	m.data[id] = input

	return input, nil
}

// DeleteUser implements store.Store interface.
// This function deletes an existing user by its ID.
func (m *memory) DeleteUser(ctx context.Context, id string) error {
	// Protect the data from race condition and data race.
	m.Lock()
	defer m.Unlock()

	// Retrieve user with the given ID.
	if _, ok := m.data[id]; !ok {
		// Return the not found errors.
		// Here should be custom not found error implementation
		// to convert in to the proto status instead of return this error.
		return fmt.Errorf("user with ID '%s' doesn't found", id)
	}

	// Delete record.
	delete(m.data, id)

	return nil
}
