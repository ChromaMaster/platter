#!/usr/bin/env bash

command -v podman >/dev/null 2>&1 || { echo "Podman is not installed. Aborting." >&2; exit 1; }

IMAGE="docker.io/golangci/golangci-lint"
IMAGE_VERSION="v1.50.1"

LINTER_ARGS="--fix"

# https://golangci-lint.run/usage/install/#docker
podman run --rm -v $(pwd):/app -w /app $IMAGE:$IMAGE_VERSION golangci-lint run $LINTER_ARGS $@
