package proto

//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/status/status.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/health/health.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/common/error_response.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/common/types.proto

//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/account-svc/account.proto
