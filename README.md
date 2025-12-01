# Sandbox API

A Go-based REST API for managing users and notes, built with Echo framework, Gorm ORM, and PostgreSQL.

## Features

- User management endpoints
- Note CRUD operations
- RESTful API design
- PostgreSQL database with Gorm
- CLI commands with Cobra
- Comprehensive testing with testify
- OpenAPI 3.0 specification

## Installation

1. Ensure you have Go 1.24 or later installed.
2. Clone the repository:
   ```bash
   git clone https://github.com/mshirdel/sandbox.git
   cd sandbox
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Build the application
```bash
go build .
```

### Run the server
```bash
./sandbox serve
```

The server will start on `http://localhost:8080`.

### CLI Commands
- `serve`: Start the HTTP server
- `run`: (Additional commands available via Cobra)

## API Documentation

The API is documented using OpenAPI 3.0. See `openapi.yaml` for the complete specification.

### Endpoints

- **Users**: `/v1/users`
  - GET `/v1/users/{id}` - Get user by ID
  - POST `/v1/users` - Save a message

- **Notes**: `/v1/notes`
  - GET `/v1/notes` - Get all notes
  - POST `/v1/notes` - Create a new note
  - GET `/v1/notes/{id}` - Get note by ID
  - PUT `/v1/notes/{id}` - Update a note
  - DELETE `/v1/notes/{id}` - Delete a note

## Testing

Run all tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test -run TestName ./path/to/package
```

## Deployment

### Docker Compose
Use the provided `docker-compose.yml` for local development with PostgreSQL.

### Helm Chart
Deployment configuration is available in `deployment/values.yaml`.

## Development

- **Lint**: `golangci-lint run`
- **Format**: `go fmt ./...`
- **Vet**: `go vet ./...`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## License

This project is licensed under the MIT License.