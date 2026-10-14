# 01-SYSTEM-OVERVIEW.md - Theta Backend System Overview

## 1. Executive Summary & Purpose
`Theta Backend (theta_be)` is an enterprise-grade Go backend microservice designed for hackathon rapid prototyping and seamless evolution into production. Built upon Clean Architecture Layering and Context Engineering, it provides:
- **Pluggable Persistence**: Zero-setup In-Memory storage engine out-of-the-box, switchable to production-ready PostgreSQL with GORM via a single environment variable (`DB_DRIVER=postgres`).
- **Standardized Communication**: Deterministic JSON response envelopes with predictable success/error payloads.
- **Contract-Driven Boundaries**: OpenAPI 3.0 specification and 5-layer contracts ensuring zero-guesswork integration.

---

## 2. Core Capabilities
1. **Lifecycle Management**: Entity state machine transitions (`DRAFT` -> `ACTIVE` -> `SUSPENDED` / `ARCHIVED`).
2. **Context Engineering Compliance**: Integrated guardrails preventing AI hallucination, state corruption, and architectural erosion.
3. **Graceful Lifecycle**: HTTP server lifecycle management with OS signal interception (`SIGINT`, `SIGTERM`) and drain timeout.
4. **Resilient Middleware**: Structured request logging, CORS preflight handling, and panic recovery preventing unexpected process exits.

---

## 3. High-Level System Architecture

```
                    +---------------------------+
                    |   Clients / External APIs |
                    +-------------+-------------+
                                  | HTTP / JSON
                                  v
+-----------------------------------------------------------------------+
|  THETA BACKEND (theta_be)                                             |
|                                                                       |
|  +-----------------------------------------------------------------+  |
|  |  DELIVERY LAYER (delivery/http)                                 |  |
|  |  - Gin HTTP Router & Route Groups                               |  |
|  |  - CORS, Recovery, Request Logger Middleware                    |  |
|  |  - CamelCase DTOs & Validation Binding                          |  |
|  +--------------------------------+--------------------------------+  |
|                                   |                                   |
|                                   v Usecase Interface (Ports)         |
|  +--------------------------------+--------------------------------+  |
|  |  APPLICATION LAYER (application/usecase)                        |  |
|  |  - Business Orchestration & Workflow Validation                |  |
|  |  - Clock & ID Injection for Deterministic Testing               |  |
|  +--------------------------------+--------------------------------+  |
|                                   |                                   |
|                                   v Entity & Repo Interface           |
|  +--------------------------------+--------------------------------+  |
|  |  DOMAIN LAYER (domain)                                          |  |
|  |  - Pure Entities (Record) & Invariant Guards                    |  |
|  |  - Repository Interfaces (Ports)                                |  |
|  |  - Enums (Status) & Domain Errors                               |  |
|  +--------------------------------+--------------------------------+  |
|                                   ^                                   |
|                                   | Implements Domain Port            |
|  +--------------------------------+--------------------------------+  |
|  |  INFRASTRUCTURE LAYER (infrastructure)                          |  |
|  |  - Database Adapter: PostgreSQL GORM Driver                     |  |
|  |  - In-Memory Adapter: Thread-Safe In-Memory Store               |  |
|  +-----------------------------------------------------------------+  |
+-----------------------------------------------------------------------+
```
