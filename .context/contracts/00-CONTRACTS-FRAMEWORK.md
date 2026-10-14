# 00-CONTRACTS-FRAMEWORK.md: The 5-Layer Contract Architecture

> **Definition**: A contract is a formal agreement governing interactions between systems, modules, human engineers, and AI agents. It defines pre-conditions, post-conditions, and immutable invariants.

---

## 🏛️ The Contract Pyramid

```
                            ┌───────────────────────────────┐
                            │   5. WORKFLOW & OPERATIONAL   │  ➔ Enforcer: CI, Linters, Human Review
                            ├───────────────────────────────┤
                            │   4. BEHAVIORAL & STATE       │  ➔ Enforcer: Automated Unit & Scenario Tests
                            ├───────────────────────────────┤
                            │   3. PERSISTENCE & EVENT      │  ➔ Enforcer: PostgreSQL Engine & Migrations
                            ├───────────────────────────────┤
                            │   2. API & BOUNDARY NETWORK   │  ➔ Enforcer: OpenAPI Validator & Schema Tests
                            ├───────────────────────────────┤
                            │   1. CODE & INTERFACE         │  ➔ Enforcer: Go Compiler (`gc`) & `go vet`
                            └───────────────────────────────┘
```

---

## 📋 Summary Table

| Level | Document | Governed Scope | Enforcer |
| :--- | :--- | :--- | :--- |
| **1. Code & Interface** | `01-CODE-INTERFACE-CONTRACT.md` | Method signatures, return types, sentinel errors, domain purity | Go Compiler (`gc`) |
| **2. API & Network** | `02-API-NETWORK-CONTRACT.md` | HTTP methods, JSON envelopes, DTO validation, status codes | OpenAPI & Integration Tests |
| **3. Persistence & Event** | `03-PERSISTENCE-EVENT-CONTRACT.md` | SQL DDL, indexes, foreign keys, migrations | PostgreSQL Engine |
| **4. Behavioral & State** | `04-BEHAVIORAL-STATE-CONTRACT.md` | State machines, business invariant rules, scenarios | Automated Table-Driven Tests |
| **5. Workflow & Operational** | `05-WORKFLOW-OPERATIONAL-CONTRACT.md`| TDD mandatory cycle, scope isolation, DoD | CI/CD Pipeline & Review |
