
# Retable

A modern table management system built with Go, PostgreSQL, GORM, Redis and WebSocket.

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
