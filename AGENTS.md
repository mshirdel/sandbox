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

### Testing
- Use `testify/assert` for assertions
- Test files: `*_test.go`
- Test functions: `Test_FunctionName_Scenario`