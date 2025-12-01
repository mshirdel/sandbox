# AGENTS.md

## Build/Test Commands
- **Build**: `go build .`
- **Test all**: `go test ./...`
- **Test single**: `go test -run TestName ./path/to/package`
- **Lint**: `golangci-lint run`
- **Format**: `go fmt ./...`
- **Vet**: `go vet ./...`

## Code Style Guidelines

### Project Structure
- Echo framework for HTTP server
- Cobra for CLI commands
- testify for testing
- logrus for logging

### Naming Conventions
- Exported types/functions: PascalCase
- Unexported: camelCase
- Struct fields: PascalCase

### Imports
- Standard library first
- Third-party packages second
- Local packages last

### Struct Tags
- JSON: `json:"fieldName"`
- Params: `param:"id"`
- Query: `query:"sort"`
- Validation: `validate:"required"`

### Error Handling
- Return errors from functions
- Use `ctx.JSON(status, response)` for HTTP responses
- Log errors with `logrus.Errorf()`

### Controllers
- Define Request/Response structs
- Controller struct with methods
- Routes in separate Router struct

### Models
- Use Gorm for ORM with PostgreSQL
- Include soft delete with `gorm.DeletedAt`
- Use proper Gorm tags: `gorm:"primaryKey"`, `gorm:"not null"`, etc.
- JSON tags for API responses
- Validation tags for input validation
- Custom table names with `TableName()` method

### Repository
- Repository pattern for data access layer
- Interface-based design for testability
- Gorm-based implementations
- Repository files: `*_repository.go`
- Repository test files: `*_repository_test.go`

### Testing
- Use `testify/assert` for assertions
- Test files: `*_test.go`
- Test functions: `Test_FunctionName_Scenario`
- Use in-memory SQLite for repository tests

### Database
- PostgreSQL for production
- Auto-migration with Gorm
- Connection handling in application layer