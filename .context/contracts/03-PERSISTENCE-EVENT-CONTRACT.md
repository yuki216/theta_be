# 03-PERSISTENCE-EVENT-CONTRACT.md: Persistence & Storage Contracts (Layer 3)

> **Level 3**: Governs PostgreSQL DDL, database integrity constraints, indexing, and migration standards.

---

## 1. Schema Invariants
1. **Primary Keys**: UUID v4 or ULID strings (standardized on 36-character UUID representation).
2. **Timestamps**:
   - `created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP`
   - `updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP`
   - `deleted_at TIMESTAMPTZ` (for soft-deletes if enabled)
3. **Foreign Keys**: Must have explicit `ON DELETE` actions (`CASCADE` or `RESTRICT`).

---

## 2. Table Definition Standard (PostgreSQL)

```sql
CREATE TABLE IF NOT EXISTS theta_records (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(32) NOT NULL DEFAULT 'DRAFT',
    metadata JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_theta_records_status ON theta_records (status);
CREATE INDEX IF NOT EXISTS idx_theta_records_created_at ON theta_records (created_at DESC);
```
