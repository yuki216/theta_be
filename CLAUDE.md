# CLAUDE.md - Theta Backend Context Router

Refer to `AGENTS.md` for full execution protocols.

## Context Source & SSOT
All project context and technical specifications are sourced from the `docs/` folder:
- `docs/01-system-overview.md`
- `docs/02-architecture-clean-layering.md`
- `docs/03-api-spec.md`
- `docs/04-database-design.md`
- `docs/05-feature-addition-workflow.md`

## Context Engineering Hub
- **Mandatory Anti-Pitfalls Rules**: `.context/00-AI-PITFALLS-AND-RULES.md`
- **5-Layer Contract Framework**: `.context/contracts/00-CONTRACTS-FRAMEWORK.md`
- **Domain Overview & Context**: `.context/01-PROJECT-CONTEXT.md` (Synced from `docs/`)
- **API & Network Contracts**: `.context/02-SHARED-CONTRACTS/`
- **Backend Technical Guardrails**: `.context/backend/TECHNICAL-GUARDRAILS.md`
- **Business Logic & State Scenarios**: `.context/backend/business-matrix.md`

## Clean Architecture Layers
- `cmd/app` — composition root & wiring
- `domain/` — pure entities, repository interfaces, enums (no Gin or GORM imports)
- `application/usecase/` — business orchestration & table-driven unit tests
- `delivery/http/` — Gin handlers, DTOs (`json:"camelCase"`), routing, middleware
- `infrastructure/` — GORM models, PostgreSQL connection, repository implementations
- `pkg/` — config, errs, httpresponse

## Claude Code Execution Protocol
When implementing features or refactoring:
1. Review `docs/` and `.context/02-SHARED-CONTRACTS/` first.
2. Follow TDD: write/update unit tests in `application/usecase/<feature>/` before finishing implementation.
3. Verify every change with:
   - `go test -v ./...`
   - `go build -v -o bin/app.exe ./cmd/app`
4. Never assume success without running verification commands.
