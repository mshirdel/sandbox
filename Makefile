.PHONY: build test test-single lint format vet clean help

# Build the application
build:
	go build .

# Run all tests
test:
	go test ./...

# Run a single test (usage: make test-single TEST=TestName PACKAGE=./path/to/package)
test-single:
	go test -run $(TEST) $(PACKAGE)

# Run linter
lint:
	golangci-lint run

# Format code
format:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Clean build artifacts
clean:
	go clean
	rm -f sandbox

# Run all checks (format, vet, lint, test)
check: format vet lint test

# Show help
help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  test        - Run all tests"
	@echo "  test-single - Run a single test (set TEST and PACKAGE variables)"
	@echo "  lint        - Run golangci-lint"
	@echo "  format      - Format code with go fmt"
	@echo "  vet         - Run go vet"
	@echo "  clean       - Clean build artifacts"
	@echo "  check       - Run all checks (format, vet, lint, test)"
	@echo "  help        - Show this help message"