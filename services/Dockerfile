# This is the base image of the services.
# In this stage we prepare all required common data to dockerize our services.
FROM golang:1.15

RUN apt-get update                                                        && \
  DEBIAN_FRONTEND=noninteractive apt-get install -yq --no-install-recommends \
    curl                                                                     \
    git                                                                      \
    unzip                                                                    \
  && rm -rf /var/lib/apt/lists/*

# Install Protobuf.
ARG PROTOBUF_VERSION=3.12.4
RUN curl -sOL "https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip" && \
  unzip protoc-*.zip              && \
  mv bin/protoc /usr/local/bin    && \
  mv include/* /usr/local/include && \
  rm -f protoc-*.zip

# Go to the root of the project.
WORKDIR /go/src/github.com/begmaroman/go-micro-boilerplate/

# Switch on the Go modules. Go modules will be switched on by default from Go 1.13 onwards.
ENV GO111MODULE=on

# Copy module files.
COPY go.mod .
COPY go.sum .

# Download project dependencies.
RUN go mod download

# Install go tools.
RUN go install \
  github.com/go-openapi/runtime \
  github.com/tylerb/graceful \
  github.com/jessevdk/go-flags \
  github.com/micro/protoc-gen-micro/v2 \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/go-swagger/go-swagger/cmd/swagger

# Copy the project on build trigger.
ONBUILD COPY . .

# Generate protobuf go files on build trigger.
ONBUILD RUN ./scripts/generate.sh .
