package restapisvc

import (
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"

	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi/operations"
)

// NewRestAPI sets up the swagger API of the service.
func NewRestAPI(log *logrus.Logger) *operations.RestAPISvcAPI {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	return operations.NewRestAPISvcAPI(swaggerSpec)
}
