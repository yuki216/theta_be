# OpenCode Rules - Theta Backend (theta_be)

This repository follows Context Engineering and Clean Architecture Layering.

## Key Directives
- **Routing**: Read `AGENTS.md` and `docs/` before implementing features.
- **Contract Enforcement**: Check `.context/02-SHARED-CONTRACTS/openapi-spec.yaml` and `.context/backend/business-matrix.md`.
- **Layer Integrity**:
  - `domain/` -> Entities and repository interfaces (zero external dependencies).
  - `application/` -> Usecase business logic + unit tests.
  - `delivery/` -> Gin HTTP handlers and camelCase DTOs.
  - `infrastructure/` -> PostgreSQL GORM and In-Memory repository adapters.
  - `pkg/` -> Standard response envelopes, error handling, and config.
  - `cmd/app/` -> Wiring and startup.

## Validation
```bash
go test -v ./...
go build -v -o bin/app.exe ./cmd/app
```
