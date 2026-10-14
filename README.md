# Theta Backend (`theta_be`)

Production-ready Go backend boilerplate adhering to Context Engineering and Clean Architecture Layering (SALT standard: `cmd/domain/application/delivery/infrastructure/pkg`).

---

## 🏛️ Context Engineering Navigation

- **Central Hub**: `AGENTS.md`
- **Context Router**: `CLAUDE.md`
- **AI Operational Rules**: `.context/00-AI-PITFALLS-AND-RULES.md`
- **Domain Overview**: `.context/01-PROJECT-CONTEXT.md`
- **API Contracts**: `.context/02-SHARED-CONTRACTS/openapi-spec.yaml`
- **5-Layer Contract Architecture**: `.context/contracts/00-CONTRACTS-FRAMEWORK.md`
- **Backend Technical Guardrails**: `.context/backend/TECHNICAL-GUARDRAILS.md`
- **Scenario & Business Matrix**: `.context/backend/business-matrix.md`

---

## 📂 Clean Architecture Layering

```
theta_be/
├── cmd/
│   └── app/                 <-- Composition root (DI wiring & server lifecycle)
├── domain/                  <-- Pure enterprise rules
│   ├── entity/              <-- Domain entities & invariant methods
│   ├── repository/          <-- Domain repository interfaces (ports)
│   ├── enum/                <-- Domain enums & types
│   └── valueobject/         <-- Value objects
├── application/
│   └── usecase/             <-- Business orchestrators & table-driven unit tests
├── delivery/
│   └── http/                <-- HTTP handlers, DTOs, routing, middleware
├── infrastructure/
│   ├── database/            <-- PostgreSQL GORM connection management
│   └── repository/          <-- GORM & in-memory repository adapters
├── pkg/
│   ├── config/              <-- Configuration loader
│   ├── errs/                <-- Sentinel & typed errors
│   └── httpresponse/        <-- Standard response envelope helper
└── .context/                <-- Progressive Disclosure Context Engineering
```

---

## 🚀 Quick Start

### 1. Run Tests
```bash
go test -v ./...
```

### 2. Build Binary
```bash
go build -v -o bin/app.exe ./cmd/app
```

### 3. Run Locally (In-Memory by default, zero external DB required)
```bash
go run ./cmd/app
```

### 4. Health Check
```bash
curl http://localhost:8080/healthz
```
