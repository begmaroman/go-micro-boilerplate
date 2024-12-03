# Stage 1: Build artifacts.
# We use the following command "make build-base-image" to build the base image
FROM go-micro-boilerplate-base:latest AS builder

ARG PLUGINS=plugins_nats

# Install the service binary.
RUN PLUGINS=${PLUGINS} make build-rest-api-svc

# Stage 2: Prepare all required data to run the service.
FROM alpine:3.9

# Install ca-certificates, bash
RUN apk -v --update add ca-certificates bash

# Copy entrypoint and service executable.
COPY ./services/rest-api-svc/entrypoint.sh .
COPY --from=builder /go/src/github.com/begmaroman/go-micro-boilerplate/services/rest-api-svc/dist/rest-api-svc /go/bin/rest-api-svc

# This is needed for healthcheck
EXPOSE 5678

ENTRYPOINT ["./entrypoint.sh", "--server_address=:5678"]
