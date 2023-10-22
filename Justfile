default:
    @just --list

run CMD="platter":
    go run ./cmd/{{ CMD }}/main.go

test:
    go test -failfast -shuffle on -parallel $(nproc) -cover ./...

fmt:
    go fmt ./...

tidy:
    go mod tidy
