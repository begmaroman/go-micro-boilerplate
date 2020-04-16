package microservice

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	restapisvc "github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/account"
	"github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/swaggergen/restapi/operations"
	"github.com/begmaroman/go-micro-boilerplate/utils/rpc"
)

// MicroService is the micro-service.
type MicroService struct {
	svc  web.Service
	opts *Options
	log  *logrus.Logger
}

// Init initializes the service.
func Init(clientOpts *ClientOptions) (*MicroService, error) {
	var opts *Options

	// Create micro-service.
	svc := web.NewService(
		web.Name(rpc.AccountServiceName),
		web.Version(clientOpts.Version),
		web.Flags(flags...),
		web.Action(func(c *cli.Context) {
			opts = buildOptions(c)
		}),
	)

	// Parse command-line arguments.
	if err := svc.Init(); err != nil {
		return nil, errors.Wrap(err, "unable to initialize web service")
	}

	return New(svc, opts, clientOpts)
}

// New is the constructor of the service.
func New(svc web.Service, opts *Options, clientOpts *ClientOptions) (*MicroService, error) {
	// Validate options.
	if err := opts.Validate(); err != nil {
		return nil, errors.Wrap(err, "options validation failed")
	}

	// Init dependencies. Here we create a client to send RPC requests to account-svc.
	accountClient := accountproto.NewAccountService(rpc.AccountServiceName, client.DefaultClient)

	// Create handler of REST endpoints.
	accountHandler := account.NewRestHandler(&account.RestHandlerOptions{
		AccountService: accountClient,
		Logger:         clientOpts.Log,
	})

	// Create API.
	restAPI := restapisvc.NewRestAPI(clientOpts.Log)

	// Create healthcheck handler. This is needed for LB.
	restAPI.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
		return operations.NewGetHealthOK()
	})

	// Register API.
	accountHandler.Register(restAPI)

	// Setup handler.
	svc.Handle("/", restAPI.Serve(nil))

	// Initialize service with updated configuration.
	if err := svc.Init(); err != nil {
		return nil, errors.Wrap(err, "unable to initialize web service")
	}

	return &MicroService{
		svc:  svc,
		opts: opts,
		log:  clientOpts.Log,
	}, nil
}

// Run runs the service.
func (s *MicroService) Run() error {
	if s.opts.IsTest {
		s.log.Info("Running in test mode!")
	}

	// Start service.
	if err := s.svc.Run(); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}
