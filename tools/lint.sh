#!/usr/bin/env bash

command -v podman >/dev/null 2>&1 || { echo "Podman is not installed. Aborting." >&2; exit 1; }

IMAGE="docker.io/golangci/golangci-lint"
IMAGE_VERSION="v1.56.2"

LINTER_ARGS="--fix"

# https://golangci-lint.run/usage/install/#docker
podman run --rm \
	-v $(pwd):/app \
	-v $(go env GOCACHE):/cache/go \
	-v $(go env GOPATH)/pkg:/go/pkg:ro \
	-e GOCACHE=/cache/go \
	-e GOLANGCI_LINT_CACHE=/cache/go \
	-w /app \
	$IMAGE:$IMAGE_VERSION golangci-lint run --verbose $LINTER_ARGS $@
