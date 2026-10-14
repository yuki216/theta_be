# 00-AI-PITFALLS-AND-RULES.md: AI Operational Guardrails

> **Purpose**: Convert mechanical and cognitive weaknesses of LLM code generation into rigid operational guardrails to ensure robust, production-grade Go backend systems.

---

## 🛑 1. Anti-Sycophancy Rule (Echo Chamber Mitigation)

### Trap:
LLMs tend to blindly agree with user proposals, even when the requested design introduces security vulnerabilities, data races, or unmaintainable complexity.

### Operational Guardrails:
1. **Mandatory Critical Analysis**:
   - The AI must evaluate architectural proposals for:
     - Concurrency races (goroutine leaks, unbuffered channel deadlocks, unprotected shared state).
     - Database bottlenecks ($O(N)$ query loops, missing indexes, unindexed foreign keys).
     - Edge-case breakage (zero values, nil pointers, connection timeouts).
2. **Explicit Trade-off Presentation**:
   - When suggesting alternatives, the AI must outline: `User Proposal (Risks) vs Architecture Recommendation (Benefits)`.

---

## 🛑 2. "Don't Imitate, Understand!" (Codebase Alienation Mitigation)

### Trap:
Generating dense boilerplate with generic abstraction layers (e.g., factories of factories) that neither the engineer nor future maintainers can easily understand or debug.

### Operational Guardrails:
1. **Explainability by Default**:
   - Functions must have clear declarative flow and single responsibilities.
   - Do not create unrequested abstractions (no interface with only one private consumer unless required for testing).
2. **Computational Complexity Documentation**:
   - State Time and Space complexity ($O(N)$) in function header comments for non-trivial algorithms.
3. **Targeted Diffs**:
   - Produce minimal, high-signal code edits.

---

## 🛑 3. Anti-Confabulation Rule (Hallucination & Version Drift Mitigation)

### Trap:
The AI invents non-existent functions, struct tags, or packages due to probabilistic pattern completion.

### Operational Guardrails:
1. **Strict Symbol Verification**:
   - Never import fictitious third-party packages or reference methods not in Go stdlib or declared project dependencies.
2. **Locked Tech Stack**:
   - Go 1.22+ standard patterns.
   - Gin HTTP (`github.com/gin-gonic/gin`).
   - GORM (`gorm.io/gorm`) with PostgreSQL driver.
3. **Negative Proof Requirement**:
   - If a library lacks a feature, explicitly acknowledge it rather than assuming an imaginary API exists.

---

## 🛑 4. Anti-Cognitive Laziness (Engineering Discipline & TDD)

### Trap:
Skipping requirements validation, omitting edge-case tests, and declaring tasks complete without proof of execution.

### Operational Guardrails:
1. **Spec-First Invariant**:
   - Define contracts in `.context/contracts/` or scenario matrices before modifying implementation.
2. **Mandatory Verification**:
   - All tasks must be backed by real execution output: `go test -v ./...` and `go build ./...`.
3. **Deterministic Testing**:
   - Never use unpredictable `time.Now()` directly in test assertions. Inject clock functions (`nowFunc func() time.Time`).
