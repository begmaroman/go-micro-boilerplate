# Stage 1: Build artifacts.
# We use the following command "make build-base-image" to build the base image
FROM go-micro-boilerplate-base:latest AS builder

# Install the service binary.
RUN make build-account-svc

# Stage 2: Prepare all required data to run the service.
FROM alpine:3.9

# Install ca-certificates, bash
RUN apk -v --update add ca-certificates bash curl

# Copy entrypoint and service executable.
COPY ./services/account-svc/entrypoint.sh .
COPY --from=builder /go/src/github.com/begmaroman/go-micro-boilerplate/services/account-svc/dist/account-svc /go/bin/account-svc

# This is needed for healthcheck
EXPOSE 5678

ENTRYPOINT ["./entrypoint.sh"]
