package memory

import (
	"github.com/sirupsen/logrus"

	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/store"
)

// Options contains the options to create a memory store
type Options struct {
	Log *logrus.Logger
}

// memory implements store.Store interface.
// Represents the store logic using in-memory data store.
type memory struct {
	log *logrus.Logger
}

// New is the constructor of memory
func New(opts *Options) store.Store {
	return &memory{
		log: opts.Log,
	}
}
