SDR – Software Design Record

Projeto: Clinical Data Ingestion Service
Versão: 2.0
Autor: Jason Silva / Principal Backend Engineer
Data: 2025-08-15

1. Contexto do Problema

O sistema legado de ingestão de dados clínicos é:

Monolítico e pouco observável

Propenso a falhas que corrompem ou perdem dados

Não suporta multi-tenancy confiável

Não possui rastreabilidade auditável

Não possui métricas ou tracing para operação

Necessidade: Um serviço resiliente, escalável, auditável e observável, com suporte a grandes volumes de dados clínicos e futura integração com ML/IA.

2. Objetivos

Processar grandes volumes de dados sem degradação de performance

Alta resiliência: ingestão contínua mesmo diante de falhas de APIs externas

Extensibilidade: fácil integração de novos pipelines ou modelos ML/IA

Visibilidade completa: logs estruturados, métricas, tracing

Segurança e conformidade HIPAA

3. Requisitos Funcionais e Não-Funcionais
Requisito	Implementação Atualizada
Multi-tenancy	Row-Level Security (RLS), tenant_id obrigatório em queries, Data Layer isolado
Idempotência	UPSERT PostgreSQL, hash estável FHIR, Crypto pkg para PHI
Observabilidade	Logs JSON (slog), métricas Prometheus, traces OpenTelemetry, dashboards Datadog
Resiliência	Circuit breaker (go-kit), retries com exponential backoff, DLQ para falhas
Deployment	Terraform IaC, ECS Fargate, autoescalonamento, SQS distribuído, RDS multi-tenant
4. Arquitetura

Hexagonal Architecture (Ports & Adapters)

HTTP/API Layer: Handlers REST, autenticação, rate-limiting, ETag, validation

Business Layer: Use cases, regras de negócio, normalização FHIR, hash estável

Data Layer: Repositórios, UPSERT idempotente, audit log, migrations multi-tenant

External Layer: FHIR Client resiliente, circuit breaker, retries, timeout

Observability Layer: Logs estruturados, métricas, tracing, dashboards

Infra Layer: ECS Fargate, RDS, SQS, KMS, Terraform

5. Fluxo de Ingestão

Scheduler publica mensagem SQS (tenant + cursor)

Workers consomem fila → invocam FHIR Client resiliente

Business Layer valida, normaliza e calcula hash

Data Layer persiste via UPSERT + audit log

Falhas → DLQ + alertas automáticos

Observabilidade → logs, métricas e traces

6. Melhorias para Performance

Batch inserts

Connection pooling otimizado

Cache local para recursos FHIR

Goroutines & Channels para paralelismo

DLQ & alerting automático

Autoescalonamento baseado em métricas