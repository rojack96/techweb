.PHONY: help setup verify clean build build-backend build-frontend build-images push deploy up down logs health status restart

# Colors
BLUE=\033[0;34m
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m

# Variables
REGISTRY ?= localhost:5000
NAMESPACE ?= techweb
VERSION ?= latest

help:
	@echo "$(BLUE)================================================$(NC)"
	@echo "$(BLUE)  TechWeb Build & Deployment Helper$(NC)"
	@echo "$(BLUE)================================================$(NC)"
	@echo ""
	@echo "$(GREEN)Setup Commands:$(NC)"
	@echo "  make setup              Setup project (make scripts executable, create .env)"
	@echo "  make verify             Verify prerequisites and configuration"
	@echo ""
	@echo "$(GREEN)Build Commands:$(NC)"
	@echo "  make build              Build everything (backend + frontend + images)"
	@echo "  make build-backend      Build backend only"
	@echo "  make build-frontend     Build frontend only"
	@echo "  make build-images       Build Docker images only"
	@echo "  make build-clean        Clean build artifacts"
	@echo ""
	@echo "$(GREEN)Docker Registry Commands:$(NC)"
	@echo "  make push               Push images to registry"
	@echo "  make list-images        List built images"
	@echo ""
	@echo "$(GREEN)Deployment Commands:$(NC)"
	@echo "  make up                 Start services (docker-compose up)"
	@echo "  make down               Stop services (docker-compose down)"
	@echo "  make restart            Restart all services"
	@echo "  make status             Show service status"
	@echo "  make logs               Show all service logs"
	@echo "  make logs-api           Show backend logs"
	@echo "  make logs-ui            Show frontend logs"
	@echo "  make health             Health check all services"
	@echo ""
	@echo "$(GREEN)Combined Workflows:$(NC)"
	@echo "  make full               Full pipeline (build + verify + deploy)"
	@echo "  make dev                Dev workflow (build + deploy + logs)"
	@echo ""
	@echo "$(YELLOW)Variables (can be overridden):$(NC)"
	@echo "  REGISTRY=$(REGISTRY)"
	@echo "  NAMESPACE=$(NAMESPACE)"
	@echo "  VERSION=$(VERSION)"
	@echo ""
	@echo "$(YELLOW)Examples:$(NC)"
	@echo "  make build"
	@echo "  make REGISTRY=gcr.io build"
	@echo "  make up"
	@echo "  make logs"
	@echo "  make dev"
	@echo ""

# Setup
setup:
	@echo "$(BLUE)[Setup]$(NC) Initializing project..."
	@bash setup.sh

verify:
	@echo "$(BLUE)[Verify]$(NC) Checking prerequisites..."
	@bash verify.sh

# Build targets
build: build-backend build-frontend build-images
	@echo "$(GREEN)[Build]$(NC) Complete!"

build-backend:
	@echo "$(BLUE)[Build]$(NC) Backend..."
	@DOCKER_REGISTRY=$(REGISTRY) DOCKER_NAMESPACE=$(NAMESPACE) ./build.sh --backend

build-frontend:
	@echo "$(BLUE)[Build]$(NC) Frontend..."
	@DOCKER_REGISTRY=$(REGISTRY) DOCKER_NAMESPACE=$(NAMESPACE) ./build.sh --frontend

build-images:
	@echo "$(BLUE)[Build]$(NC) Docker images..."
	@DOCKER_REGISTRY=$(REGISTRY) DOCKER_NAMESPACE=$(NAMESPACE) ./build.sh --images

build-clean:
	@echo "$(BLUE)[Cleanup]$(NC) Build artifacts..."
	@./build.sh --clean

# Registry commands
push:
	@echo "$(BLUE)[Push]$(NC) Images to registry..."
	@DOCKER_REGISTRY=$(REGISTRY) DOCKER_NAMESPACE=$(NAMESPACE) ./push.sh push $(REGISTRY)

list-images:
	@echo "$(BLUE)[Images]$(NC) Built images..."
	@DOCKER_REGISTRY=$(REGISTRY) DOCKER_NAMESPACE=$(NAMESPACE) ./push.sh list

# Deployment commands
up:
	@echo "$(BLUE)[Deploy]$(NC) Starting services..."
	@REGISTRY=$(REGISTRY) NAMESPACE=$(NAMESPACE) VERSION=$(VERSION) ./deploy.sh up $(VERSION)

down:
	@echo "$(BLUE)[Deploy]$(NC) Stopping services..."
	@./deploy.sh down

restart: down up
	@echo "$(GREEN)[Deploy]$(NC) Services restarted!"

status:
	@echo "$(BLUE)[Status]$(NC) Service status..."
	@./deploy.sh status

logs:
	@echo "$(BLUE)[Logs]$(NC) All services (Ctrl+C to exit)..."
	@./deploy.sh logs

logs-api:
	@echo "$(BLUE)[Logs]$(NC) Backend (Ctrl+C to exit)..."
	@./deploy.sh logs api

logs-ui:
	@echo "$(BLUE)[Logs]$(NC) Frontend (Ctrl+C to exit)..."
	@./deploy.sh logs ui

health:
	@echo "$(BLUE)[Health]$(NC) Checking services..."
	@./deploy.sh health

# Workflows
full: setup verify build up
	@echo "$(GREEN)[Complete]$(NC) Full pipeline done!"

dev: build up logs
	@echo "$(GREEN)[Dev]$(NC) Dev environment ready!"

# Docker commands for advanced usage
docker-clean:
	@echo "$(BLUE)[Docker]$(NC) Cleaning unused images and volumes..."
	@docker image prune -f
	@docker volume prune -f

docker-deep-clean:
	@echo "$(BLUE)[Docker]$(NC) Deep clean (removing dangling images)..."
	@docker image prune -a -f
	@docker volume prune -f
	@docker system prune -a -f

docker-ps:
	@docker ps -a

docker-images:
	@docker images | grep -E "(streetcats|techweb)" || echo "No images found"

# Development helpers
shell-api:
	@docker-compose exec api /bin/sh

shell-ui:
	@docker-compose exec ui /bin/sh

shell-db:
	@docker-compose exec postgis psql -U postgres -d gis

# Info
version:
	@echo "Build script version info:"
	@./build.sh --version

info:
	@echo "$(BLUE)Project Info:$(NC)"
	@echo "Backend: $(shell grep '^go' backend/streetcats-api/go.mod)"
	@echo "Frontend: $(shell grep '"version"' frontend/streetcats/package.json | head -1)"
	@echo "Git Branch: $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo 'N/A')"
	@echo "Git Commit: $(shell git rev-parse --short HEAD 2>/dev/null || echo 'N/A')"
	@echo ""

.DEFAULT_GOAL := help
