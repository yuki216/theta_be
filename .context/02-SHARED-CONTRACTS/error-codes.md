# Error Codes Registry & Contract

Standardized error codes returned across all HTTP responses in `theta_be`.

| HTTP Status | Error Code | Description |
| :--- | :--- | :--- |
| `400 Bad Request` | `INVALID_INPUT` | Request body or query parameters failed validation. |
| `401 Unauthorized` | `UNAUTHORIZED` | Authentication header missing or invalid. |
| `403 Forbidden` | `FORBIDDEN` | Caller lacks permission to perform this action. |
| `404 Not Found` | `RESOURCE_NOT_FOUND` | Target entity does not exist. |
| `409 Conflict` | `RESOURCE_ALREADY_EXISTS` | Unique constraint violated. |
| `422 Unprocessable Entity` | `INVALID_STATE_TRANSITION` | Requested state change violates business lifecycle rules. |
| `500 Internal Server Error` | `INTERNAL_SERVER_ERROR` | Unhandled failure or unexpected infrastructure error. |
