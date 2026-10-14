# Backend Technical Guardrails - Theta Backend (`theta_be`)

Technical guardrails and architecture enforcement for Go development in `theta_be`.

---

## 1. Stack & Toolchain
- **Language**: Go 1.22+ (tested on Go 1.26+)
- **HTTP Framework**: `github.com/gin-gonic/gin`
- **ORM & Database**: `gorm.io/gorm` with `gorm.io/driver/postgres`
- **UUID**: `github.com/google/uuid`
- **Configuration**: Environment variables with `.env` fallback via `github.com/joho/godotenv`

---

## 2. Clean Architecture Layering (SALT Pattern)

```
theta_be/
├── cmd/
│   └── app/                 <-- Composition root (dependency injection & server lifecycle)
├── domain/                  <-- Pure domain logic
│   ├── entity/              <-- Domain entities & invariant guards
│   ├── repository/          <-- Domain repository interfaces
│   ├── enum/                <-- Domain enums & types
│   └── valueobject/         <-- Domain value objects
├── application/
│   └── usecase/             <-- Business logic orchestrators & unit tests
├── delivery/
│   └── http/                <-- Gin handlers, request/response DTOs, routing, middleware
├── infrastructure/
│   ├── database/            <-- PostgreSQL GORM connection management
│   └── repository/          <-- GORM & in-memory repository implementations
└── pkg/
    ├── config/              <-- Configuration loader
    ├── errs/                <-- Domain/Application sentinel & code errors
    └── httpresponse/        <-- Standardized HTTP response envelopes
```

---

## 3. Strict Coding Invariants
1. **Dependency Inversion**: Outer layers (`delivery`, `infrastructure`) depend on inner layers (`application`, `domain`). Inner layers NEVER import outer layers.
2. **Domain Purity**: Never import Gin, GORM, or SQL packages into `domain/`.
3. **JSON Tagging**: All DTO struct fields in `delivery/` must use `json:"camelCase"`.
4. **Sentinel Errors**: Declare errors in `pkg/errs` or `domain/` and map them to HTTP status codes via `pkg/httpresponse`.
5. **No Naked Panic**: Always handle errors explicitly.
