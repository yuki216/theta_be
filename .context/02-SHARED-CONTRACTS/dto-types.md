# DTO Types & Schema Invariants

All JSON properties MUST adhere to `camelCase` naming conventions.

## Record DTO
```json
{
  "id": "string (uuid)",
  "title": "string (1-255 chars)",
  "description": "string (optional)",
  "status": "DRAFT | ACTIVE | SUSPENDED | ARCHIVED",
  "createdAt": "string (ISO 8601)",
  "updatedAt": "string (ISO 8601)"
}
```

## Create Record Request
```json
{
  "title": "string (required, min 1, max 255)",
  "description": "string (optional)"
}
```
