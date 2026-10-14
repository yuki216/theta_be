# 01-CODE-INTERFACE-CONTRACT.md: Code & Interface Contracts (Layer 1)

> **Level 1**: Governs in-memory signatures, return values, interface purity, and sentinel errors within Go packages.

---

## 1. Core Principles
- **Explicit Interfaces**: Consumers define interfaces where applicable; domain defines repository and usecase contracts.
- **Purity of Domain**: `internal/domain` must not import infrastructure or transport packages (no `gorm.io`, no `gin-gonic`).
- **Sentinel Errors**: Explicit error variables (`ErrNotFound`, `ErrConflict`, `ErrInvalidInput`) must be returned rather than generic formatted strings.

---

## 2. Invariants
1. Functions returning errors must return `(nil, err)` or zero-value representations on failure.
2. Context (`context.Context`) must be the first parameter in all repository and usecase methods.
3. No struct fields with unexported mutability that violate encapsulation.
