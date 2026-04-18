# Use pre-compiled binary from build.sh
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from build artifacts (context is root directory)
COPY .build/backend/api .

# Copy config files
COPY backend/streetcats-api/configs/config.production.yml ./configs/config.yml

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./api"]
