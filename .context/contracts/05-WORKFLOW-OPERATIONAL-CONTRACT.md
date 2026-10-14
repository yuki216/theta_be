# 05-WORKFLOW-OPERATIONAL-CONTRACT.md: Workflow & Operational Contracts (Layer 5)

> **Level 5**: Governs developer/AI operational protocols, TDD enforcement, code review criteria, and Definition of Done.

---

## 1. Operational Invariants
1. **Mandatory Verification**: Every code modification must pass:
   ```bash
   go test -v ./...
   go build -v ./...
   ```
2. **Deterministic Time**: No uncontrolled `time.Now()` calls in assertions; pass clock providers or use deterministic fixtures.
3. **No Panic in Business Logic**: Return explicit `(result, error)` tuples across all layers.

---

## 2. Definition of Done (DoD)
- [ ] Contract documentation updated if domain/API schema altered.
- [ ] Unit tests written for usecases and business logic.
- [ ] Code formatted with `gofmt` / `goimports`.
- [ ] Build succeeds with 0 exit code.
- [ ] All tests pass with 0 exit code.
