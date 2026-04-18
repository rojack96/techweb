#!/bin/bash

################################################################################
# TechWeb Deploy Script (Minimalist)
# Run services with docker compose
# Usage: ./deploy.sh [up|down|logs|status|help]
################################################################################

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
COMPOSE_FILE="${SCRIPT_DIR}/docker-compose.yml"
BUILD_DIR="${SCRIPT_DIR}/.build"

# Colors
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

log_info() {
    echo -e "${BLUE}ℹ${NC} $*"
}

log_success() {
    echo -e "${GREEN}✓${NC} $*"
}

log_error() {
    echo -e "${RED}✗${NC} $*"
}

check_build_artifacts() {
    if [ ! -f "${BUILD_DIR}/backend/api" ]; then
        log_error "Build artifact missing: ${BUILD_DIR}/backend/api"
        log_error "Run ./build.sh all or ./build.sh backend first."
        exit 1
    fi

    if [ ! -d "${BUILD_DIR}/frontend/dist" ]; then
        log_error "Frontend build artifact missing: ${BUILD_DIR}/frontend/dist"
        log_error "Run ./build.sh all or ./build.sh frontend first."
        exit 1
    fi
}

show_help() {
    cat <<EOF
TechWeb Deploy Script

Usage: ./deploy.sh [COMMAND]

Commands:
    up          Start services
    down        Stop services
    logs        Show logs
    logs-api    Show backend logs
    logs-ui     Show frontend logs
    status      Show service status
    help        Show this help

Examples:
    ./deploy.sh up
    ./deploy.sh logs
    ./deploy.sh down

EOF
}

main() {
    local command="${1:-help}"
    
    case "$command" in
        up)
            log_info "Starting services..."
            check_build_artifacts
            docker compose -f "${COMPOSE_FILE}" up -d
            log_success "Services started"
            docker compose -f "${COMPOSE_FILE}" ps
            echo ""
            echo "Access:"
            echo "  Frontend:  http://localhost:3000"
            echo "  Backend:   http://localhost:8080"
            ;;
        down)
            log_info "Stopping services..."
            docker compose -f "${COMPOSE_FILE}" down
            log_success "Services stopped"
            ;;
        logs)
            docker compose -f "${COMPOSE_FILE}" logs -f
            ;;
        logs-api)
            docker compose -f "${COMPOSE_FILE}" logs -f api
            ;;
        logs-ui)
            docker compose -f "${COMPOSE_FILE}" logs -f ui
            ;;
        status)
            docker compose -f "${COMPOSE_FILE}" ps
            ;;
        help|--help|-h)
            show_help
            ;;
        *)
            log_error "Unknown command: $command"
            show_help
            exit 1
            ;;
    esac
}

main "$@"

