# 05-FEATURE-ADDITION-WORKFLOW.md - Feature Addition Protocol

## 1. Step-by-Step Layer Progression

When introducing a new feature or domain entity into `theta_be`, follow this mandatory execution order:

```
[ docs/ ] ➔ [ .context/ ] ➔ [ domain/ ] ➔ [ application/ ] ➔ [ infrastructure/ ] ➔ [ delivery/ ] ➔ [ cmd/app/ ]
```

### Phase 1: Specifications & Contracts
1. Document high-level requirements in `docs/` (e.g. `docs/features/<feature>.md`).
2. Add API endpoints to `.context/02-SHARED-CONTRACTS/openapi-spec.yaml`.
3. Add test scenarios to `.context/backend/business-matrix.md`.

### Phase 2: Domain Layer
1. Define entity struct & methods in `domain/entity/<feature>.go`.
2. Define repository interface in `domain/repository/<feature>_repository.go`.
3. Define domain enums / errors if needed.

### Phase 3: Application Layer (TDD)
1. Write table-driven test cases in `application/usecase/<feature>/<feature>_usecase_test.go` (**RED**).
2. Implement use case methods in `application/usecase/<feature>/<feature>_usecase.go` (**GREEN**).
3. Verify test execution: `go test -v ./application/usecase/...`.

### Phase 4: Infrastructure Layer
1. Implement persistence adapter in `infrastructure/repository/<feature>_gorm_repository.go`.
2. Implement in-memory mock/adapter in `infrastructure/repository/<feature>_memory_repository.go`.

### Phase 5: Delivery Layer
1. Create request & response DTOs in `delivery/http/<feature>handler/dto.go`.
2. Implement HTTP handlers in `delivery/http/<feature>handler/handler.go`.
3. Mount routes in `delivery/http/router.go`.

### Phase 6: Wiring & Verification
1. Instantiate repository, use case, and handler in `cmd/app/main.go`.
2. Execute full test suite and build verification:
   ```bash
   go test -v ./...
   go build -v -o bin/app.exe ./cmd/app
   ```
