# AGENTS.md - Theta Backend Context Engineering Hub

> **Project**: `Theta Backend (theta_be)`  
> **Scope**: Backend Service Only (Go, Gin, GORM, PostgreSQL)  
> **Architecture**: Clean Architecture Layering (SALT standard: `cmd/domain/application/delivery/infrastructure/pkg`)  
> **Documentation SSOT**: `docs/` folder (Source of truth for all domain & system requirements)  
> **Supported AI Tools**: Cursor (`.cursorrules`, `.cursor/rules/`), Claude (`CLAUDE.md`), OpenCode (`.opencode/`, `AGENTS.md`), Trae (`.traerules`, `.trae/rules/`)

---

## 🧭 Progressive Disclosure Routing

Before executing any instruction, the AI agent must review the following rules:

0. **MANDATORY ANTI-AI PITFALL RULES** ➔ Read `.context/00-AI-PITFALLS-AND-RULES.md`.
0.1. **5-LAYER CONTRACT ARCHITECTURE** ➔ Read `.context/contracts/00-CONTRACTS-FRAMEWORK.md`:
   - L1: `.context/contracts/01-CODE-INTERFACE-CONTRACT.md` (Go Interfaces, Strict Types, Sentinel Errors)
   - L2: `.context/contracts/02-API-NETWORK-CONTRACT.md` (OpenAPI, JSON Envelopes, DTOs)
   - L3: `.context/contracts/03-PERSISTENCE-EVENT-CONTRACT.md` (PostgreSQL DDL, Constraints, Invariants)
   - L4: `.context/contracts/04-BEHAVIORAL-STATE-CONTRACT.md` (State Machine & Business Scenario Matrix)
   - L5: `.context/contracts/05-WORKFLOW-OPERATIONAL-CONTRACT.md` (TDD, DoD, Circuit Breakers)
0.2. **PROJECT SPECIFICATIONS & CONTEXT SOURCE** ➔ Sourced from `docs/`:
   - `docs/01-system-overview.md`
   - `docs/02-architecture-clean-layering.md`
   - `docs/03-api-spec.md`
   - `docs/04-database-design.md`
   - `docs/05-feature-addition-workflow.md`

Execution Flow for Backend Tasks:

```
                          ┌─────────────────────────────┐
                          │    AGENTS.md (Entrypoint)   │
                          └──────────────┬──────────────┘
                                         │
                                         ▼
                          ┌─────────────────────────────┐
                          │         docs/ (SSOT)        │
                          └──────────────┬──────────────┘
                                         │
                                         ▼
                     ┌───────────────────────────────────┐
                     │ .context/01-PROJECT-CONTEXT.md    │
                     ├───────────────────────────────────┤
                     │ .context/02-SHARED-CONTRACTS/     │
                     ├───────────────────────────────────┤
                     │ .context/backend/                 │
                     │  - TECHNICAL-GUARDRAILS.md        │
                     │  - business-matrix.md             │
                     └─────────────────┬─────────────────┘
                                       │
                                       ▼
                     ┌───────────────────────────────────┐
                     │ Clean Architecture Layers         │
                     │  1. domain/ (entity, repo interface)
                     │  2. application/ (usecase orchestration)
                     │  3. delivery/ (http handlers, routes)
                     │  4. infrastructure/ (db, gorm repo)
                     │  5. pkg/ (config, errs, httpresponse)
                     │  6. cmd/app/ (DI wiring)          │
                     └─────────────────┬─────────────────┘
                                       │
                                       ▼
                     ┌───────────────────────────────────┐
                     │ Implementation & Verification     │
                     │  - go test -v ./...               │
                     │  - go build -v ./...              │
                     └───────────────────────────────────┘
```

---

## 🛠️ Tool Integration Registry

| Tool | Rule / Config File | Behavior |
| :--- | :--- | :--- |
| **Cursor** | `.cursorrules` & `.cursor/rules/context-engineering.mdc` | Automatically loads project guardrails, clean architecture constraints, and testing protocols. |
| **Claude (Code / Desktop)** | `CLAUDE.md` | Provides deterministic routing, TDD protocols, and terminal verification instructions. |
| **OpenCode** | `AGENTS.md` & `.opencode/rules.md` | Enforces open-agent spec standards and context-first development. |
| **Trae (ByteDance)** | `.traerules` & `.trae/rules/project_rules.md` | Guides Trae agent to follow layer isolation and run verification checks. |

---

## 🛡️ Non-Negotiable Operational Rules

1. **Context Sourced from `docs/`**: All domain logic, functional scopes, and technical constraints must trace back to specifications in the `docs/` folder. Never invent domain rules without documenting them in `docs/` first.
2. **Clean Layering Boundary Integrity**:
   - `domain/`: Pure business entities, value objects, domain repository interfaces, and enums. Zero external imports (no Gin, no GORM, no SQL).
   - `application/`: Application usecases orchestrating domain logic. Depends only on domain interfaces and `pkg/errs`.
   - `delivery/`: Delivery mechanisms (HTTP Gin handlers, routing, middleware). Translates HTTP requests to usecase calls and maps errors via `pkg/httpresponse`.
   - `infrastructure/`: External drivers, database connections (PostgreSQL GORM), and repository implementations adhering to domain repository interfaces.
   - `pkg/`: Reusable packages (`config`, `errs`, `httpresponse`).
   - `cmd/app/`: Composition root (main.go). Initializes config, connects infrastructure, wires usecases, and launches delivery servers.
3. **Contract-First Development**: All API changes, schema alterations, or domain field additions must be specified in `.context/02-SHARED-CONTRACTS/` before modifying code.
4. **Deterministic Verification**:
   - Every modification must be validated with:
     ```bash
     go test -v ./...
     go build -v ./...
     ```
   - No task is marked complete without zero-error test output.
5. **Standardized Response Envelope**:
   - Success: `{"success": true, "data": ..., "meta": ...}`
   - Error: `{"success": false, "error": {"code": "...", "message": "...", "details": ...}}`
