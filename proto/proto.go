package proto

//go:generate protoc --experimental_allow_proto3_optional --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/status/status.proto
//go:generate protoc --experimental_allow_proto3_optional --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/health/health.proto
//go:generate protoc --experimental_allow_proto3_optional --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/common/error_response.proto
//go:generate protoc --experimental_allow_proto3_optional --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/common/types.proto

//go:generate protoc --experimental_allow_proto3_optional --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-micro --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/begmaroman/go-micro-boilerplate/proto/account-svc/account.proto
