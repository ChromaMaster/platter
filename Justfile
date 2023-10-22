default:
    @just --list

run CMD="platter":
    go run ./cmd/{{ CMD }}/main.go

fmt:
    go fmt ./...
