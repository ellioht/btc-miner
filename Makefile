.PHONY: test build

test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Tests passed."

build:
	@echo "Building..."
	@go build -o bin/ ./...
	@echo "Build complete."