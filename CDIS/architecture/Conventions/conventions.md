# Conventions & Guidelines – Clinical Data Ingestion Service (CDIS)

**Version:** 1.0  
**Author:** Jason Silva / Principal Engineer
**Date:** August 15, 2025  

---

## 1. Branch Structure

The repository follows the **simplified Gitflow** convention:

- `main`: stable version in production.
- `release1/CDIS-v1`: release branch (contains all code from day 1).
- `task/<microtask>`: task branch within the release.
  - Ex.: `task/CDIS-docs`, `task/CDIS-runbook`, `task/CDIS-sdr`.
- Each microtask must be created **from the corresponding release branch**.
- Delete obsolete remote branches to keep the repository clean.

---

## 2. Naming Conventions

### Branches
type/short-description
- Types: `release`, `task`, `hotfix`, `experiment`.
- Example: `task/CDIS-adr-docs`, `hotfix/CDIS-dlq-fix`.

### Commits (Conventional Commits)
type(scope): description
- Common types: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`.
- Scope: module or microtask.
- Example: `docs(adr): add initial ADRs v1.0 for Clinical Data Ingestion Service`
- Always write **imperative and clear**.

### Tags
- `vX.Y.Z` for releases.
- Example: `v1.0.0`, `v2.0.0-beta`.

---

## 3. Code Style

### Python
- PEP8 with `black` + `isort`.
- Mandatory typing with `mypy`.
- Clear directory structure: `handlers/`, `services/`, `shared/`, `utils/`.

### Go
- `gofmt` mandatory.
- Small and explicit interfaces.
- Avoid circular dependencies between packages.

### JavaScript / React
- ESLint + Prettier.
- Functional components whenever possible.
- Hooks following standard rules.

---

## 4. Pull Requests / Code Review

- PR must be created from the task branch.
- PR checklist:
- [ ] Code tested with TDD.
- [ ] Unit tests >= 80% coverage.
- [ ] Code reviewed by at least 1 colleague.
- [ ] ADR/Runbook/Docs updated if there are architectural changes.
- [ ] No secrets or credentials exposed.

- PRs must have a **clear description**, referencing ticket or microtask.

---

## 5. Observability and Logging

- Structured logs (JSON), using `slog` or equivalent.
- PHI should never be logged without masking.
- Metrics should be sent to Prometheus/Datadog:
  - Throughput
  - Latency
  - Failures / Errors
- Circuit Breakers and DLQs monitored with critical alerts.

---

## 6. Security and Compliance

- All PHI data must be encrypted (at rest and in transit).
- Audit logs are mandatory for critical operations.
- Follow HIPAA compliance checklist.
- Never commit credentials, tokens, or keys.

---

## 7. Extra Procedures

- Deployment and migration scripts must be versioned.
- Integration tests must run in an isolated environment before merging.
- ADR documents, Runbooks, and SDRs must always be versioned and reviewed.

---

> By following these conventions, we ensure clean code, decision traceability, HIPAA compliance, and ease of maintenance in a scalable system.

