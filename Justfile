cover-file := "cover.out"

# Show this help
default:
    @just --list

# Run the application
run CMD="platter":
    go run ./cmd/{{ CMD }}/main.go

# Run tests
test: test-unit test-int

# Run unit tests
test-unit:
    go test -short -failfast -shuffle on -parallel $(nproc) -cover -coverprofile {{cover-file}} ./...

# Run integration tests
test-int:
    go test -failfast -shuffle on -parallel $(nproc) -cover -coverprofile {{cover-file}} ./...

# Show coverage information in HTML format
coverage: test
    go tool cover -html={{cover-file}}

# Run the formatter (go fmt)
fmt:
    go fmt ./...

# Make sure everything is up to date
tidy:
    go mod tidy

# Upgrade all dependencies
upgrade:
    go get -u ./...

# Run the linter (golangci-lint)
lint *args="":
    ./tools/lint.sh {{ args }}

# Remove everything that is generated
clean:
    rm {{cover-file}}
