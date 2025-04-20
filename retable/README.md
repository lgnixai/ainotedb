# Retable

A real-time collaborative table management system.

## Features

- Real-time collaboration
- Space management
- Table management with multiple field types
- Record management with CRUD operations
- View management with different display options
- Real-time updates via WebSocket
- Redis caching for improved performance
- Authentication and authorization
- API documentation

## Getting Started

1. Configure environment variables in `.env`
2. Start Redis server
3. Run `go mod download`
4. Run `go run cmd/server/main.go`

## Architecture

The system follows a clean architecture pattern with:

- Domain layer (core entities and business rules)
- Application layer (use cases)
- Infrastructure layer (external interfaces)
- Interface layer (controllers and presenters)

## Caching Strategy

The system uses Redis for:
- Record caching
- View configuration caching
- Session management
- Real-time message queuing

Cache invalidation is handled automatically on:
- Record updates
- Schema changes
- View modifications

## Directory Structure
```
retable/
├── cmd/                    # Entry points
│   └── server/            # Server application
├── internal/              # Private application code
│   ├── config/           # Configuration
│   ├── core/             # Core domain types
│   ├── auth/             # Authentication
│   ├── space/            # Space management
│   ├── table/            # Table management
│   ├── field/            # Field management
│   ├── record/           # Record management
│   ├── view/             # View management
│   ├── cache/            # Cache management
│   └── ws/               # WebSocket handling
├── pkg/                   # Public libraries
│   ├── logger/           # Logging
│   └── errors/           # Error handling
├── migrations/            # Database migrations
└── scripts/              # Build and deployment scripts
```

## Setup Instructions

1. Install dependencies:
   ```bash
   go mod init retable
   go mod tidy
   ```

2. Configure environment:
   Copy `.env.example` to `.env` and adjust settings.

3. Run migrations:
   ```bash
   go run scripts/migrate.go
   ```

4. Start server:
   ```bash
   go run cmd/server/main.go
   ```

## Development

- Go 1.21+
- PostgreSQL 14+
- Redis 6+

## License

MIT