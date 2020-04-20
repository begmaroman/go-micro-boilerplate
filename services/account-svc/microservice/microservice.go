package microservice

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	accountproto "github.com/begmaroman/go-micro-boilerplate/proto/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/proto/health"
	accountsvc "github.com/begmaroman/go-micro-boilerplate/services/account-svc"
	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/domain"
	"github.com/begmaroman/go-micro-boilerplate/services/account-svc/store/memory"
	"github.com/begmaroman/go-micro-boilerplate/utils/healthchecker"
	"github.com/begmaroman/go-micro-boilerplate/utils/rpc"
)

// MicroService is the micro-service.
type MicroService struct {
	svc     micro.Service
	handler *accountsvc.Handler
	opts    *Options
	log     *logrus.Logger
}

// Init initializes the service.
func Init(clientOpts *ClientOptions) (*MicroService, error) {
	var opts *Options

	// Create micro-service.
	svc := micro.NewService(
		micro.Name(rpc.AccountServiceName),
		micro.Version(clientOpts.Version),
		micro.Flags(flags...),
		micro.Action(func(c *cli.Context) {
			opts = buildOptions(c)
		}),
	)

	// Parse command-line arguments.
	svc.Init()

	return New(svc, opts, clientOpts)
}

// New is the constructor of the service.
func New(svc micro.Service, opts *Options, clientOpts *ClientOptions) (*MicroService, error) {
	// Validate options.
	if err := opts.Validate(); err != nil {
		return nil, errors.Wrap(err, "options validation failed")
	}

	// Create a self-pinger client.
	selfPingClient := health.NewSelfPingClient(svc, accountproto.NewAccountService(rpc.AccountServiceName, svc.Client()))

	// Create store layer using in-memory data store.
	// Here can be any implementation of the store layer.
	store := memory.New(&memory.Options{
		Log: clientOpts.Log,
	})

	// Create business layer.
	service := domain.New(&domain.Options{
		Store: store,
		Log:   clientOpts.Log,
	})

	// Create RPC handler.
	handler := accountsvc.NewHandler(&accountsvc.Options{
		Service:        service,
		SelfPingClient: selfPingClient,
		Log:            clientOpts.Log,
	})

	// Register the service.
	accountproto.RegisterAccountServiceHandler(svc.Server(), handler)

	return &MicroService{
		svc:     svc,
		handler: handler,
		opts:    opts,
		log:     clientOpts.Log,
	}, nil
}

// Run runs the service.
func (s *MicroService) Run() error {
	if s.opts.IsTest {
		s.log.Info("Running in test mode!")
	}

	// Run helathcheck endpoint.
	shutdown := healthchecker.Run(s.log, healthchecker.WrapRPC(s.handler.Health), nil)

	// Stop helathcheck endpoint after RPC service stop.
	s.svc.Init(micro.AfterStop(shutdown))

	// Start service.
	if err := s.svc.Run(); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}
