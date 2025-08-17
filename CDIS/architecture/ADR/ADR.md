# Architecture Decision Records (ADR) - Clinical Data Ingestion Service v1.0

**Author:** Jason Silva / Principal Backend Engineer
**Date:** August 15, 2025  
**Version:** 1.0  

---

## ADR-001: Hexagonal Architecture (Ports & Adapters)

**Context:**  
The ingestion service needs to be decoupled from the infrastructure to facilitate maintenance, testing, and replacement of implementations (such as database, FHIR client, queues, etc.). The tight coupling between business logic and infrastructure details hinders development and system resilience.

**Decision:**  
Adopt Hexagonal Architecture (Ports & Adapters). The core of the application (business logic) will be isolated and will communicate with the outside world (infrastructure) through well-defined interfaces (ports). The implementations of these interfaces (adapters) will be responsible for interacting with the actual infrastructure services.

**Justification:**  
This architecture ensures that business rules are infrastructure-agnostic, facilitating unit testing (through mocks), allowing for the replacement of implementations (e.g., SQS for Kafka), and keeping the code cleaner and more modular.

**Consequences:**  
- **Positive:** Greater system maintainability and flexibility. Core unit tests become faster and more reliable.  
- **Negative:** The initial design and implementation complexity is slightly higher. Requires a clear understanding of Dependency Inversion and Dependency Injection.

---

## ADR-002: Idempotence, Parallelism, and Performance

**Context:**  
High clinical data ingestion requires efficiency and assurance that message reprocessing will not cause duplication or inconsistency. Unit insertions are slow, and race conditions can occur during concurrent processing.

**Decision:**  
- Use Goroutines and channels with a Worker Pool for concurrent message processing, ensuring high throughput and controlling the load on the database.  
- Implement bulk inserts with the PostgreSQL UPSERT clause to ensure atomicity and idempotence. A stable hash of each FHIR resource will be used as the key for UPSERT.

**Justification:**  
This combined approach dramatically improves ingestion performance by processing multiple resources in parallel and inserting large volumes of data at once. UPSERT and hashing ensure that the system is robust enough to handle reprocessing and failures without duplicating data.

**Consequences:**  
- **Positive:** High throughput, reduced latency, and guaranteed data consistency.  
- **Negative:** Requires careful monitoring of parallelism to avoid database overload.

---

## ADR-003: Multi-Tenant

**Context:**  
The service must support multiple clients (tenants) in a secure and isolated manner, sharing the infrastructure efficiently.

**Decision:**  
Implement multi-tenant architecture with Row-Level Security (RLS). Database tables will be shared, but each row will be associated with a mandatory `tenant_id`.

**Justification:**  
This approach optimizes resource usage and simplifies maintenance (no need to manage multiple databases). RLS ensures that a tenant's queries can only access data belonging to that tenant, maintaining data security and isolation.

**Consequences:**  
- **Positive:** Cost reduction and simplified administration.  
- **Negative:** All queries and repository operations must validate the `tenant_id`, requiring extra attention during development.

---

## ADR-004: Resilience and Failure Recovery

**Context:**  
The ingestion service depends on external APIs that can fail and uses a queuing system for asynchronous processing. The system needs to be autonomous in recovering from failures and ensure message reprocessing without data loss.

**Decision:**  
- Implement retry with exponential backoff and circuit breaker for all HTTP calls to external APIs.  
- Use SQS queues with Dead-Letter Queues (DLQ) and automatic alerts.

**Justification:**  
The combination of retry and circuit breaker protects the application against intermittent failures and prevents overload of external services. The use of DLQs ensures that messages with permanent failures are not lost and can be inspected and reprocessed manually, ensuring data integrity.

**Consequences:**  
- **Positive:** Increased system stability, reliability, and protection.  
- **Negative:** Precise configuration of timeouts and thresholds for circuit breakers and DLQs.

---


## ADR-005: Observability and Telemetry

**Context:**  
Proactive operation of the ingestion service requires complete real-time visibility into its status, performance, and behavior.

**Decision:**  
Adopt a comprehensive observability strategy, including Structured Logs (JSON), Custom Metrics (Prometheus), and Distributed Tracing (OpenTelemetry). All telemetry will be sent to a centralized monitoring dashboard (Datadog).

**Justification:**  
This decision allows for the creation of comprehensive dashboards, the configuration of automatic alerts, and detailed auditing of operations. Distributed tracing is essential for debugging performance bottlenecks in a distributed system.

**Consequences:**  
- **Positive:** Proactive operation, reduced incident response time, and the ability to track latency and throughput.  
- **Negative:** Requires the development of middleware and collectors, in addition to ensuring that sensitive data (PHI) is masked in the logs.

---

## ADR-006: HIPAA Security and Compliance

**Context:**  
The ingestion service handles sensitive health data (PHI), requiring maximum confidentiality, integrity, and traceability to ensure HIPAA compliance.

**Decision:**  
- All sensitive data will be encrypted at rest using KMS (Key Management Service).  
- No sensitive data (PHI) will be exposed in logs or dashboards. A masking process will be applied to all telemetry.  
- All critical operations will be recorded in an Audit Logging system to ensure traceability and auditing.

**Justification:**  
These measures ensure compliance with strict HIPAA requirements, protecting data privacy and minimizing the risk of information leaks.

**Consequences:**  
- **Positive:** Regulatory compliance, high security, and robust system against threats.  
- **Negative:** Implementation of encryption, masking, and auditing requires care and attention to detail.

---

## ADR-007: Infrastructure and Deployment

**Context:**  
We need an infrastructure that supports scalability, high availability, and reliable operation with minimal manual intervention.

**Decision:**  
Use ECS Fargate for container orchestration, Terraform for Infrastructure as Code (IaC), SQS for queues, multi-tenant RDS for the database, and KMS for key management.

**Justification:**  
This managed and automated architecture allows for load-based auto-scaling of the service, eliminates the need for server management, and ensures consistency of the environment through IaC.

**Consequences:**  
- **Positive:** Reduced operational complexity, increased uptime, and the ability to scale quickly in response to demand.  
- **Negative:** Infrastructure costs may be higher than dedicated solutions, but are offset by reduced operating costs.

