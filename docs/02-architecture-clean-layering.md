# 02-ARCHITECTURE-CLEAN-LAYERING.md - Clean Architecture Layering Specification

## 1. Principles & Dependency Inversion
`theta_be` implements Uncle Bob's Clean Architecture structured around the SALT Go standard:
`cmd/` -> `delivery/` -> `application/` -> `domain/` <- `infrastructure/` + `pkg/`.

### Dependency Rules:
1. **Source code dependencies point inward**: Nothing in an inner circle can know anything at all about something in an outer circle.
2. **Domain is pure**: `domain/` contains zero framework imports (`net/http`, `gin`, `gorm`, `sql`).
3. **Application depends only on Domain & Shared Pkg**: `application/` orchestrates use cases through domain interfaces.
4. **Infrastructure and Delivery depend on Application and Domain**: Adapters implement domain interfaces and translate transport protocols.

---

## 2. Directory Layout & Layer Responsibilities

| Layer | Directory | Responsibility | Invariant Rules |
| :--- | :--- | :--- | :--- |
| **Composition Root** | `cmd/app/` | Wire dependency injection, read config, handle OS signals, start server. | No business logic; only instantiation and wiring. |
| **Domain** | `domain/` | Enterprise rules, entities, value objects, domain enums, repository interfaces. | Pure Go stdlib only. No external ORM or HTTP tags. |
| **Application** | `application/` | Application use cases, orchestration, transaction boundaries, unit tests. | Operates on domain interfaces. Uses injected clocks/UUIDs. |
| **Delivery** | `delivery/` | Transport adapters (HTTP Gin handlers, routing, middleware, request/response DTOs). | Translates requests to DTOs; never accesses database directly. |
| **Infrastructure** | `infrastructure/`| Database drivers, GORM models, external client adapters, persistence implementations. | Implements repository interfaces declared in `domain/repository/`. |
| **Shared Pkg** | `pkg/` | Cross-cutting helpers (`config`, `errs`, `httpresponse`). | Stateless, reusable, domain-agnostic utilities. |
