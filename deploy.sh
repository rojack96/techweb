#!/bin/bash

################################################################################
# TechWeb Deploy Script (Minimalist)
# Run services with docker compose
# Usage: ./deploy.sh [up|down|logs|status|help]
################################################################################

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
COMPOSE_FILE="${SCRIPT_DIR}/docker-compose.yml"

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

