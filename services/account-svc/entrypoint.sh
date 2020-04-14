#!/usr/bin/env bash

set -euo pipefail

if [ "${DOCKER_COMPOSE:-}" != "true" ]; then
  # Here you can do something if its production.
  # For example, retrieve values from AWS SecretsManager to pass them to the service.
  echo "Production mode!"
fi

exec /go/bin/account-svc "$@"
