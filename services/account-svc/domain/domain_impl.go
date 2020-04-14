package domain

import (
	"github.com/sirupsen/logrus"

	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/store"
)

// Options contains options to create a domain
type Options struct {
	Store store.Store
	Log   *logrus.Logger
}

// domain implements the business/domain logic of the service
type domain struct {
	store store.Store
	log   *logrus.Logger
}

// New is the constructor of domain
func New(opts *Options) Domain {
	return &domain{
		store: opts.Store,
		log:   opts.Log,
	}
}

// HealthCheck implements Domain interface
func (d *domain) HealthCheck() error {
	return nil
}
