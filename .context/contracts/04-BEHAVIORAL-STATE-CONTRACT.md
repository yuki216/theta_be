# 04-BEHAVIORAL-STATE-CONTRACT.md: Behavioral & State Machine Contracts (Layer 4)

> **Level 4**: Governs state transitions, pre/post-conditions, and invariant validations across business workflows.

---

## 1. Lifecycle State Machine

```
              ┌─────────┐
              │  DRAFT  │
              └────┬────┘
                   │ Activate()
                   ▼
              ┌─────────┐
    ┌─────────┤ ACTIVE  ├─────────┐
    │ Suspend │         │ Archive │
    ▼         └─────────┘         ▼
┌───────────┐                ┌──────────┐
│ SUSPENDED │                │ ARCHIVED │
└─────┬─────┘                └──────────┘
      │ Reactivate
      ▼
   (ACTIVE)
```

---

## 2. Transition Rules Matrix

| Current State | Target State | Action | Pre-condition | Post-condition |
| :--- | :--- | :--- | :--- | :--- |
| `DRAFT` | `ACTIVE` | `Activate` | Title not empty, valid schema | Status becomes `ACTIVE` |
| `ACTIVE` | `SUSPENDED` | `Suspend` | Reason provided | Status becomes `SUSPENDED` |
| `SUSPENDED` | `ACTIVE` | `Reactivate` | Operator authorized | Status becomes `ACTIVE` |
| `ACTIVE` | `ARCHIVED` | `Archive` | None | Status becomes `ARCHIVED`, immutable |
| `ARCHIVED` | Any | Disallowed | Terminal state | Returns `ErrInvalidStateTransition` |
