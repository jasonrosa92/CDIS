# Runbook Operacional – Clinical Data Ingestion Service (CDIS)

**Versão:** 2.0  
**Autor:** Jason Silva / Principal Engineer 
**Data:** 15 de Agosto de 2025  

---

## 1. Objetivo e Escopo
Este runbook serve como guia técnico para a equipe de Operações e Engenharia, fornecendo procedimentos padronizados para o CDIS (Clinical Data Ingestion Service).  
O objetivo é garantir um manuseio seguro, eficiente e rastreável do serviço durante **deployments**, **monitoramento diário** e **resposta a incidentes**.

> Qualquer procedimento executado deve ser documentado em um ticket no Jira ou plataforma de rastreamento de incidentes.

---

## 2. Componentes Críticos e Observabilidade

| Componente       | Descrição                                                            | Tecnologias Chave                              | Pontos de Observabilidade |
|-----------------|---------------------------------------------------------------------|-----------------------------------------------|--------------------------|
| **API Layer**    | Gateway de entrada. Recebe requisições de ingestão via HTTP e valida dados básicos | Go (net/http, Chi Router)                     | Health Check: GET /health; Métricas: latência, taxa de sucesso/erro (HTTP 2xx, 4xx, 5xx) |
| **Business Layer** | Coração da aplicação. Processa, valida e normaliza dados FHIR      | Arquitetura Hexagonal                          | Logs estruturados (slog) indicando status de processamento de cada recurso |
| **Data Layer**   | Gerencia persistência no banco de dados                             | PostgreSQL, sqlx, Row-Level Security          | Métricas: latência de queries (UPSERT), erros de conexão; Logs: auditoria e erros do banco |
| **External Layer** | Cliente de APIs externas de dados clínicos (FHIR)                 | go-kit/circuitbreaker, go-retryablehttp       | Alertas: Circuit breaker aberto, falhas repetidas; Métricas: tempo de resposta, contagem de retries |
| **Workers/Queue** | Processamento assíncrono dos dados                                 | AWS SQS, Go Goroutines, Channels              | Métricas: tamanho da fila SQS, tempo de processamento; Alertas: mensagens na DLQ |
| **Observability** | Coleta e visualização de dados operacionais                        | zerolog, Prometheus, OpenTelemetry, Datadog   | Dashboards: latência, throughput, erros, recursos por segundo; Alertas: DLQ, Circuit Breaker |
| **Infraestrutura** | Orquestra e gerencia recursos de nuvem                             | AWS ECS Fargate, Terraform, RDS, KMS          | Logs do ECS Fargate; Métricas: CPU/Memória, instâncias ativas |

---

## 3. Procedimentos Operacionais Padrão

### 3.1 Procedimento de Deployment

**Pré-deployment:**
- Verificar status do serviço no Datadog, confirmando que não há incidentes ativos.
- Confirmar que o health check (`GET /health`) da versão atual retorna `200 OK`.
- Validar que as imagens Docker da nova versão foram construídas e estão no registro.

**Deployment:**
- Executar migrações do banco de dados (staging) via CLI `migrate`.
- Aplicar mudanças de infraestrutura via Terraform e orquestrar rollout no ECS Fargate.

**Pós-deployment:**
- Health Check: `GET /health` deve retornar `200 OK`.
- Observabilidade: Confirmar métricas fluindo para Prometheus e dashboards Datadog populados.
- Teste de sanidade: Processar pequeno volume de dados (ex.: 100 recursos FHIR) para validar ingestão sem erros.

---

### 3.2 Operação Diária e Monitoramento

**KPIs principais (Datadog dashboards):**
- **Throughput:** recursos processados por minuto.
- **Latência média:** tempo de processamento de um lote.
- **Taxa de erro:** percentual de requisições com falha (deve ser 0%).
- **Tamanho da fila SQS:** fila processada em tempo hábil.

**Alertas críticos:**
- DLQ > 0 mensagens.
- Circuit Breaker aberto por > 1 minuto.
- Erros de API externa (HTTP 5xx) acima do threshold.

---

## 4. Procedimentos de Resposta a Incidentes (Troubleshooting)

### 4.1 Mensagens na Dead-Letter Queue (DLQ)
**Identificação:** Consultar logs da DLQ para `tenant_id`, payload e motivo da falha.  
**Diagnóstico:** Analisar se falha é de reprocessamento ou dados mal formatados (hash FHIR).  
**Correção:**
- Problema de dados: corrigir payload na origem ou manualmente.  
- Problema transitório: reprocessamento resolve.  
**Reprocessamento:**
```bash
python reprocess_dlq.py --tenant <tenant_id> --message_id <message_id>
```
### 4.2 Falha de Conexão com API Externa

**Diagnóstico:** Verificar logs e alertas do Circuit Breaker no Datadog.
**Análise:** Checar retries e exponential backoff.
**Ação:**

Se falha persistente, abrir ticket de alta prioridade para time da API FHIR externa.

Documentar incidente no Jira/Confluence.

### 4.3 Erros de Banco de Dados (UPSERT)

**Diagnóstico:** Consultar logs JSON (zerolog) da requisição com falha.
**Causas comuns:** tenant_id incorreto, dados mal formatados ou inconsistência no UPSERT.
**Solução:**

Corrigir tenant_id na fonte, se necessário.

Verificar se hash do recurso já existe; se inconsistência, usar script de reconciliação.

**Importante:** nunca apagar dados em produção sem autorização explícita.

## 5. Escalonamento e Handover

Registrar incidente com: tenant_id, bundle ou recurso FHIR, timestamp e stacktrace.

Escalar para Tech Lead se:

Circuit Breaker aberto > 10 min.

DLQ > 50 mensagens.

Latência média de lote > 5 segundos de forma consistente.

Handover: ao final do turno, revisar todos incidentes e confirmar documentação e escalonamento.

## 6. Checklist de Verificação Operacional

 Health check da API retorna 200 OK.

 Métricas do Prometheus coletadas e dashboards do Datadog atualizados.

 Dead-Letter Queue vazia.

 Circuit Breaker fechado (closed).

 Logs de auditoria registrando corretamente todas as operações.