package microservice

import "github.com/sirupsen/logrus"

// Options contains the configuration parameters of the service.
type Options struct {
	IsTest bool
}

// Validate applies the validation logic to the options.
func (opts *Options) Validate() error {
	return nil
}

// ClientOptions represent external dependencies.
type ClientOptions struct {
	Version string
	Log     *logrus.Logger
}
