module github.com/begmaroman/go-micro-boilerplate

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-openapi/errors v0.19.8
	github.com/go-openapi/loads v0.19.6
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/spec v0.19.14
	github.com/go-openapi/strfmt v0.19.11
	github.com/go-openapi/swag v0.19.12
	github.com/go-openapi/validate v0.19.14
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats/v2 v2.9.1
	github.com/micro/go-plugins/registry/nats/v2 v2.9.1
	github.com/micro/go-plugins/transport/nats/v2 v2.9.1
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/tools v0.0.0-20191218040434-6f9e13bbec44 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/protobuf v1.22.0
)
