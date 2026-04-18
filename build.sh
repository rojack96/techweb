#!/bin/bash

################################################################################
# TechWeb Build Script (Minimalist)
# Build backend, frontend, and Docker images locally
# Usage: ./build.sh [backend|frontend|images|all|clean]
################################################################################

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="${SCRIPT_DIR}"
BUILD_DIR="${PROJECT_ROOT}/.build"
VERSION_FILE="${PROJECT_ROOT}/.version"

BACKEND_NAME="streetcats-api"
FRONTEND_NAME="streetcats-ui"
BACKEND_DIR="${PROJECT_ROOT}/backend/${BACKEND_NAME}"
FRONTEND_DIR="${PROJECT_ROOT}/frontend/streetcats"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
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

# ============================================================================
# Version Management
# ============================================================================

init_version() {
    if [ ! -f "${VERSION_FILE}" ]; then
        echo "1.0.0" > "${VERSION_FILE}"
    fi
}

get_version() {
    cat "${VERSION_FILE}"
}

increment_version() {
    local version=$(get_version)
    local major=$(echo $version | cut -d. -f1)
    local minor=$(echo $version | cut -d. -f2)
    local patch=$(echo $version | cut -d. -f3)
    
    patch=$((patch + 1))
    local new_version="${major}.${minor}.${patch}"
    
    echo "${new_version}" > "${VERSION_FILE}"
    echo "${new_version}"
}

# ============================================================================
# Prerequisites
# ============================================================================

check_prerequisites() {
    log_info "Checking prerequisites..."
    
    local missing=()
    for tool in docker go npm; do
        if ! command -v "$tool" &> /dev/null; then
            missing+=("$tool")
        fi
    done
    
    if [ ${#missing[@]} -gt 0 ]; then
        log_error "Missing tools: ${missing[*]}"
        exit 1
    fi
    
    log_success "Prerequisites OK"
}

# ============================================================================
# Build Functions
# ============================================================================

build_backend() {
    log_info "Building backend..."
    cd "${BACKEND_DIR}"
    
    mkdir -p "${BUILD_DIR}/backend"
    
    # Download dependencies
    go mod download
    go mod tidy
    
    # Build binary
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
        -a -installsuffix cgo \
        -ldflags "-s -w" \
        -o "${BUILD_DIR}/backend/api" \
        ./cmd/rest_api/main.go
    
    if [ ! -f "${BUILD_DIR}/backend/api" ]; then
        log_error "Backend build failed"
        exit 1
    fi
    
    local size=$(du -h "${BUILD_DIR}/backend/api" | cut -f1)
    log_success "Backend built (${size})"
}

build_frontend() {
    log_info "Building frontend..."
    cd "${FRONTEND_DIR}"
    
    # Install and build
    npm ci
    npm run build
    
    if [ ! -d "${FRONTEND_DIR}/dist" ]; then
        log_error "Frontend build failed"
        exit 1
    fi

    mkdir -p "${BUILD_DIR}/frontend"
    rm -rf "${BUILD_DIR}/frontend/dist"
    mv "${FRONTEND_DIR}/dist" "${BUILD_DIR}/frontend/dist"

    local size=$(du -sh "${BUILD_DIR}/frontend/dist" | cut -f1)
    log_success "Frontend built (${size}) -> ${BUILD_DIR}/frontend/dist"
}

build_images() {
    log_info "Building Docker images..."
    local version=$(get_version)
    
    # Backend image
    log_info "Building backend image..."
    cd "${BACKEND_DIR}"
    docker build -t "${BACKEND_NAME}:${version}" -t "${BACKEND_NAME}:latest" -f Dockerfile .
    if [ $? -ne 0 ]; then
        log_error "Backend image build failed"
        exit 1
    fi
    log_success "Backend image: ${BACKEND_NAME}:${version}"
    
    # Frontend image
    log_info "Building frontend image..."
    cd "${FRONTEND_DIR}"
    docker build -t "${FRONTEND_NAME}:${version}" -t "${FRONTEND_NAME}:latest" -f Dockerfile .
    if [ $? -ne 0 ]; then
        log_error "Frontend image build failed"
        exit 1
    fi
    log_success "Frontend image: ${FRONTEND_NAME}:${version}"
}

clean_build() {
    log_info "Cleaning build directory..."
    rm -rf "${BUILD_DIR}"
    log_success "Build cleaned"
}

# ============================================================================
# Main
# ============================================================================

show_help() {
    cat <<EOF
TechWeb Build Script

Usage: ./build.sh [COMMAND]

Commands:
    backend     Build backend only
    frontend    Build frontend only
    images      Build Docker images only
    all         Build everything (default)
    clean       Clean build artifacts
    help        Show this help

Examples:
    ./build.sh
    ./build.sh backend
    ./build.sh images

EOF
}

main() {
    local command="${1:-all}"
    
    case "$command" in
        backend)
            init_version
            check_prerequisites
            build_backend
            ;;
        frontend)
            init_version
            check_prerequisites
            build_frontend
            ;;
        images)
            init_version
            check_prerequisites
            build_images
            ;;
        all)
            init_version
            check_prerequisites
            new_version=$(increment_version)
            log_info "Version: ${new_version}"
            mkdir -p "${BUILD_DIR}"
            build_backend
            build_frontend
            build_images
            log_success "Build completed"
            ;;
        clean)
            clean_build
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
