coverage-dir := "$PWD/coverage"

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
    mkdir -p {{coverage-dir}}/unit
    go test -failfast -shuffle on -parallel $(nproc) -cover -short ./... -args -test.gocoverdir={{coverage-dir}}/unit

# Run integration tests
test-int:
    mkdir -p {{coverage-dir}}/int
    go test -failfast -shuffle on -parallel $(nproc) -cover ./... -args -test.gocoverdir={{coverage-dir}}/int

# Show coverage information in HTML format
coverage: test && coverage-clean
    go tool covdata percent -i={{coverage-dir}}/unit,{{coverage-dir}}/int -o={{coverage-dir}}/c.out
    go tool cover -html={{coverage-dir}}/c.out

# Remove coverage information
coverage-clean:
    rm -rf {{coverage-dir}}

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
clean: coverage-clean
