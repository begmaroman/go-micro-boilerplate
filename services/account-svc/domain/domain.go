package domain

// Domain represents the behavior of the business/domain logic of the service
type Domain interface {
	// HealthCheck returns an error if there is a problem with the service
	HealthCheck() error
}
