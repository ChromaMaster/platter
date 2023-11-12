#!/usr/bin/env bash

command -v docker >/dev/null 2>&1 || { echo "Docker is not installed. Aborting." >&2; exit 1; }

LINTER_VERSION="v1.50.1"
LINTER_ARGS="--fix"

# https://golangci-lint.run/usage/install/#docker
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:$LINTER_VERSION golangci-lint run $LINTER_ARGS $@
