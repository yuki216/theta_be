# Business Scenario Matrix & Domain Rules - Theta Backend

## 1. Scenario Matrix

| ID | Feature | Scenario | Input | Pre-condition | Expected Outcome | HTTP Status |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **SC-01** | Create Record | Valid input | `title="Record A"` | System up | Record created in `DRAFT` status | `201 Created` |
| **SC-02** | Create Record | Empty title | `title=""` | Any | Error `INVALID_INPUT` | `400 Bad Request` |
| **SC-03** | Get Record | Existing ID | Valid UUID | Record exists | Returns record details | `200 OK` |
| **SC-04** | Get Record | Unknown ID | Non-existent UUID | None | Error `RESOURCE_NOT_FOUND` | `404 Not Found` |
| **SC-05** | Activate | Valid transition | ID in `DRAFT` | Record in `DRAFT` | Status updated to `ACTIVE` | `200 OK` |
| **SC-06** | Activate | Already Archived | ID in `ARCHIVED` | Record is `ARCHIVED` | Error `INVALID_STATE_TRANSITION` | `422 Unprocessable` |
| **SC-07** | Delete Record | Existing ID | Valid UUID | Record exists | Record removed / soft deleted | `200 OK` |
| **SC-08** | Delete Record | Unknown ID | Non-existent UUID | None | Error `RESOURCE_NOT_FOUND` | `404 Not Found` |
