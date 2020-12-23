# Stage 1: Build artifacts.
# We use the base image build by ./scripts/build-base-image.sh script.
# If you don't have one, build it by executin ./scripts/build-base-image.sh in your terminal.
FROM go-micro-boilerplate-base:1.15-stretch-0 as builder

# Install the service binary.
RUN CGO_ENABLED=0 GOOS=linux go install ./services/rest-api-svc/cmd/rest-api-svc

# Stage 2: Prepare all required data to run the service.
FROM alpine:3.9

# Install ca-certificates, bash
RUN apk -v --update add ca-certificates bash

# Copy entrypoint and service executable.
COPY ./services/rest-api-svc/entrypoint.sh .
COPY --from=builder /go/bin/rest-api-svc /go/bin/rest-api-svc

# This is needed for healthcheck
EXPOSE 5678

ENTRYPOINT ["./entrypoint.sh", "--server_address=:5678"]
