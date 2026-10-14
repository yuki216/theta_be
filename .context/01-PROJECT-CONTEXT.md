# Layer 1: Theta Backend Project Context & Domain Overview

> **Context Origin**: Sourced directly from technical and system specifications in the `docs/` folder.
> 
> Primary Source Documents:
> - System Overview: `docs/01-system-overview.md`
> - Clean Architecture Layering: `docs/02-architecture-clean-layering.md`
> - REST API Specifications: `docs/03-api-spec.md`
> - Database Design & DDL: `docs/04-database-design.md`
> - Feature Addition Protocol: `docs/05-feature-addition-workflow.md`

---

## 1. Vision & Purpose
`Theta Backend (theta_be)` is an enterprise-grade Go backend microservice designed for hackathons and scalable production backends. It integrates Context Engineering with Clean Architecture Layering (SALT standard: `cmd/domain/application/delivery/infrastructure/pkg`).

---

## 2. Core Architecture Principles (Distilled from `docs/02-architecture-clean-layering.md`)
- **Inward Dependency Rule**: Inner layers (`domain`, `application`) never depend on outer layers (`delivery`, `infrastructure`).
- **Domain Purity**: Enterprise business rules reside solely in `domain/` without any external transport or database framework dependencies.
- **Pluggable Storage**: Dual-mode storage (In-Memory for zero-DB setup, PostgreSQL GORM for production).
- **Single Source of Truth (SSOT)**: Shared API contracts define the network boundary. No field invention allowed.
- **Stateless & Scalable**: HTTP handlers remain stateless; persistence is delegated to repository adapters.

---

## 3. Ubiquitous Language (Domain Glossary)

- **Record / Item**: Primary domain entity representing a managed resource with lifecycle states (`docs/01-system-overview.md`).
- **Status**: Lifecycle condition of a resource (`DRAFT`, `ACTIVE`, `ARCHIVED`, `SUSPENDED`).
- **AuditMetadata**: Standard timestamps (`CreatedAt`, `UpdatedAt`, `DeletedAt`).
- **CodeError**: Standard domain and application error carrying error codes mapped to HTTP status codes (`pkg/errs/errs.go`).
- **ResponseEnvelope**: Standard JSON payload format returned by all HTTP endpoints (`docs/03-api-spec.md`).

---

## 4. Context Synchronization Invariant
Whenever product requirements, domain rules, or database models change:
1. Update the authoritative document in `docs/`.
2. Sync this `.context/01-PROJECT-CONTEXT.md` to reflect the updated domain context.
3. Update contract schemas in `.context/02-SHARED-CONTRACTS/`.
