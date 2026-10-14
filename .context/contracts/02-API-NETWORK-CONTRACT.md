# 02-API-NETWORK-CONTRACT.md: API & Network Contracts (Layer 2)

> **Level 2**: Governs external HTTP boundaries, JSON serialization, status codes, and DTO specifications.

---

## 1. Response Envelope Format

### Standard Success Response (2xx)
```json
{
  "success": true,
  "data": {},
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 1
  }
}
```

### Standard Error Response (4xx / 5xx)
```json
{
  "success": false,
  "error": {
    "code": "RESOURCE_NOT_FOUND",
    "message": "The requested item does not exist",
    "details": null
  }
}
```

---

## 2. Invariants
1. All JSON keys in DTOs must use `camelCase`.
2. HTTP status codes must accurately reflect domain errors:
   - `400 Bad Request` -> `INVALID_INPUT`
   - `401 Unauthorized` -> `UNAUTHORIZED`
   - `403 Forbidden` -> `FORBIDDEN`
   - `404 Not Found` -> `RESOURCE_NOT_FOUND`
   - `409 Conflict` -> `RESOURCE_ALREADY_EXISTS`
   - `500 Internal Server Error` -> `INTERNAL_SERVER_ERROR`
3. Nil arrays in JSON responses must serialize to `[]`, not `null`.
