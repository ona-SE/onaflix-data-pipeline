# OnaFlix Data Pipeline

Batch data processing pipeline for the OnaFlix platform. Reads movies from PostgreSQL, enriches them with external data, and writes results back.

## Stack

- **Language:** Go 1.21
- **Database:** PostgreSQL (lib/pq)
- **Logging:** logrus

## Setup

```bash
go mod download
DATABASE_URL="postgres://gitpod:gitpod@localhost:5432/onaflix?sslmode=disable" go run cmd/pipeline/main.go
```

## Testing

```bash
go test ./...
```

## Build

```bash
go build -o bin/pipeline cmd/pipeline/main.go
```
