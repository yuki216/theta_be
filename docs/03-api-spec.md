# 03-API-SPEC.md - REST API Specification

## 1. Protocol & Envelope Standard
All endpoints communicate over HTTP/HTTPS with JSON payloads. Every response is encapsulated in a standardized envelope.

### Success Envelope (HTTP 200 / 201)
```json
{
  "success": true,
  "data": { ... },
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 100
  }
}
```

### Error Envelope (HTTP 4xx / 5xx)
```json
{
  "success": false,
  "error": {
    "code": "RESOURCE_NOT_FOUND",
    "message": "Resource with the specified ID was not found",
    "details": null
  }
}
```

---

## 2. Endpoints Summary

| Method | Endpoint | Description | Request Body | Response Body |
| :--- | :--- | :--- | :--- | :--- |
| `GET` | `/healthz` | Health check probe | None | `{"status":"ok","timestamp":"..."}` |
| `POST` | `/api/v1/records` | Create new record | `{"title":"...","description":"..."}` | `RecordResponseDTO` |
| `GET` | `/api/v1/records` | List records (paginated) | Query: `?page=1&limit=10` | `[]RecordResponseDTO` + `meta` |
| `GET` | `/api/v1/records/:id` | Get record by ID | None | `RecordResponseDTO` |
| `DELETE` | `/api/v1/records/:id` | Delete record by ID | None | `{"message":"Record deleted successfully"}` |
| `PATCH` | `/api/v1/records/:id/activate` | Activate draft record | None | `RecordResponseDTO` |
