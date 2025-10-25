# Makefile for Crypto Exchange Trading System

.PHONY: build run test clean docker docker-compose help

# Variables
APP_NAME=crypto-exchange
DOCKER_IMAGE=$(APP_NAME):latest
DOCKER_COMPOSE_FILE=docker-compose.yml

# Default target
help:
	@echo "Available commands:"
	@echo "  build          Build the application"
	@echo "  run            Run the application locally"
	@echo "  test           Run tests"
	@echo "  test-coverage  Run tests with coverage"
	@echo "  lint           Run linting"
	@echo "  clean          Clean build artifacts"
	@echo "  docker         Build Docker image"
	@echo "  docker-run     Run with Docker Compose"
	@echo "  docker-stop    Stop Docker Compose services"
	@echo "  docker-logs    Show Docker Compose logs"
	@echo "  deps           Download dependencies"
	@echo "  migrate        Run database migrations"
	@echo "  seed           Seed database with sample data"

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) ./cmd/server

# Run the application locally
run:
	@echo "Running $(APP_NAME)..."
	go run ./cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linting
lint:
	@echo "Running linting..."
	golangci-lint run

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	docker system prune -f

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Build Docker image
docker:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run with Docker Compose
docker-run:
	@echo "Starting services with Docker Compose..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker Compose services
docker-stop:
	@echo "Stopping Docker Compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Show Docker Compose logs
docker-logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# Database migrations (when running locally)
migrate:
	@echo "Running database migrations..."
	@echo "Migrations are handled automatically by the application"

# Seed database with sample data
seed:
	@echo "Seeding database..."
	@echo "Sample data is inserted via init.sql script"

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/air-verse/air@latest
	@echo "Development tools installed"

# Hot reload during development
dev:
	@echo "Starting development server with hot reload..."
	air

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	go test -tags=integration -v ./...

# Generate API documentation
docs:
	@echo "Generating API documentation..."
	swag init -g cmd/server/main.go

# Check for security vulnerabilities
security:
	@echo "Checking for security vulnerabilities..."
	go list -json -m all | nancy sleuth

# Performance benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

# Vet code
vet:
	@echo "Vetting code..."
	go vet ./...

# Complete check (format, vet, lint, test)
check: fmt vet lint test
	@echo "All checks completed successfully!"

# Production build
build-prod:
	@echo "Building production binary..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/$(APP_NAME) ./cmd/server

# Docker build for production
docker-prod:
	@echo "Building production Docker image..."
	docker build -f Dockerfile.prod -t $(APP_NAME):prod .

# Deploy to staging
deploy-staging:
	@echo "Deploying to staging..."
	docker-compose -f docker-compose.staging.yml up -d

# Deploy to production
deploy-prod:
	@echo "Deploying to production..."
	@echo "Please ensure you have proper deployment scripts for production"

# Backup database
backup-db:
	@echo "Creating database backup..."
	docker-compose exec postgres pg_dump -U user crypto_exchange > backup_$(shell date +%Y%m%d_%H%M%S).sql

# Restore database
restore-db:
	@echo "Restoring database..."
	@echo "Usage: make restore-db BACKUP_FILE=backup_20231215_120000.sql"
	@if [ -z "$(BACKUP_FILE)" ]; then \
		echo "Please specify BACKUP_FILE=<backup_file.sql>"; \
		exit 1; \
	fi
	docker-compose exec -T postgres psql -U user -d crypto_exchange < $(BACKUP_FILE)

# Monitor logs
logs:
	@echo "Monitoring application logs..."
	docker-compose logs -f crypto-exchange

# Health check
health:
	@echo "Checking application health..."
	curl -s http://localhost:8080/health | jq .

# Load test
load-test:
	@echo "Running load tests..."
	@echo "Please install and configure your preferred load testing tool"