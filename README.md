# Bids App with Redis

A Go application for managing bids with Redis as the data store.

## Folder Structure

The project follows a simplified Go project layout:

- `/cmd`: Main applications for this project
  - `/app`: Main application entry point
- `/internal`: Private application and library code
  - `/config`: Configuration management
  - `/services`: Application-specific services
    - `/redis`: Redis service implementation
  - `/api`: API-specific code
    - `/routes`: API route definitions
  - `/util`: Application-specific utility functions and helpers

## Redis Client

The Redis client is implemented in `internal/services/redis/client.go` and provides a wrapper around the go-redis library. The service layer in `internal/services/redis/service.go` provides application-specific Redis functionality.

Features:
- Connection pooling and management
- Automatic connection testing
- Bidirectional error handling
- Type-safe Redis operations

## Configuration

The application can be configured using environment variables:

- `SERVER_PORT`: HTTP server port (default: 8080)
- `REDIS_HOST`: Redis host (default: localhost)
- `REDIS_PORT`: Redis port (default: 6379)
- `REDIS_DB`: Redis database number (default: 0)
- `REDIS_PASSWORD`: Redis password (default: empty)

## Development

To run the application:

```bash
go run cmd/app/main.go
```

To build the application:

```bash
go build -o bin/app cmd/app/main.go
```

## API Endpoints

- `GET /`: Welcome message
- `GET /health`: Health check endpoint 