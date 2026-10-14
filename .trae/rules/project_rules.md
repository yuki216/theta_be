# Project Rules - Trae IDE

## Architecture & Code Conventions
- Follow the SALT Go clean architecture standard (`cmd/domain/application/delivery/infrastructure/pkg`).
- `domain/` must remain pure: no external packages, web frameworks, or ORMs.
- Response payloads must use the standard envelope defined in `pkg/httpresponse`.
- DTO fields must use `json:"camelCase"`.
- Always write table-driven unit tests for use cases before marking tasks complete.

## Verification
Execute in terminal:
```bash
go test -v ./...
go build -v -o bin/app.exe ./cmd/app
```
