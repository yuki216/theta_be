# 04-DATABASE-DESIGN.md - Database Schema & Persistence Strategy

## 1. Storage Strategy
`theta_be` supports dual persistence drivers:
1. **Memory Driver (`DB_DRIVER=memory`)**:
   - Zero-dependency thread-safe in-memory map protected by `sync.RWMutex`.
   - Designed for fast local development, unit tests, and instant hackathon demos.
2. **PostgreSQL Driver (`DB_DRIVER=postgres`)**:
   - Production-grade relational storage with GORM ORM.
   - Automatic migrations on boot and index management.

---

## 2. Table Schema DDL (PostgreSQL)

```sql
CREATE TABLE IF NOT EXISTS theta_records (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(32) NOT NULL DEFAULT 'DRAFT',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_theta_records_status ON theta_records (status);
CREATE INDEX IF NOT EXISTS idx_theta_records_created_at ON theta_records (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_theta_records_deleted_at ON theta_records (deleted_at);
```
