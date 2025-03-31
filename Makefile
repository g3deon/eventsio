.PHONY: all fmt lint test clean

all: fmt lint test

fmt:
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "Code formatted successfully."

lint:
	@echo "Running Go linter..."
	@golangci-lint run
	@echo "Linting completed successfully."

test:
	@echo "Running Go tests..."
	@go test ./tests/... -v
	@echo "All tests passed successfully."

clean:
	@echo "Cleaning up..."
	@rm -rf ./bin ./coverage.out
	@echo "Cleanup complete."

install-hooks:
	@echo "Installing Git pre-commit hook..."
	@cp githooks/pre-commit.sh .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "Pre-commit hook installed."

docker-up:
	@echo "Starting AMQP test environment..."
	@docker-compose -f deployments/amqp-test_docker-compose.yml up -d
	@echo "AMQP environment started."

docker-down:
	@echo "Stopping AMQP test environment..."
	@docker-compose -f deployments/amqp-test_docker-compose.yml down
	@echo "AMQP environment stopped."

coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

