# Architecture Decision Records (ADR) - Clinical Data Ingestion Service v1.0

**Autor:** Jason Silva / Principal Backend Engineer
**Data:** 15 de Agosto de 2025  
**Versão:** 2.0  

---

## ADR-001: Arquitetura Hexagonal (Ports & Adapters)

**Contexto:**  
O serviço de ingestão precisa ser desacoplado da infraestrutura para facilitar manutenção, testes e a substituição de implementações (como banco de dados, cliente FHIR, filas, etc.). O acoplamento rígido entre a lógica de negócio e os detalhes de infraestrutura dificulta o desenvolvimento e a resiliência do sistema.

**Decisão:**  
Adotar a Arquitetura Hexagonal (Ports & Adapters). O núcleo da aplicação (lógica de negócio) será isolado e se comunicará com o mundo externo (infraestrutura) através de interfaces bem definidas (ports). As implementações dessas interfaces (adapters) serão responsáveis por interagir com os serviços de infraestrutura reais.

**Justificativa:**  
Essa arquitetura garante que as regras de negócio sejam agnósticas à infraestrutura, facilitando testes unitários (através de mocks), permitindo a substituição de implementações (ex: SQS para Kafka) e mantendo o código mais limpo e modular.

**Consequências:**  
- **Positivas:** Maior manutenibilidade e flexibilidade do sistema. Testes unitários do core se tornam mais rápidos e confiáveis.  
- **Negativas:** A complexidade inicial de projeto e implementação é um pouco maior. Exige entendimento claro de Inversão de Dependência e Dependency Injection.

---

## ADR-002: Idempotência, Paralelismo e Performance

**Contexto:**  
A alta ingestão de dados clínicos requer eficiência e a garantia de que reprocessamentos de mensagens não causarão duplicação ou inconsistência. Inserções unitárias são lentas e race conditions podem ocorrer durante o processamento concorrente.

**Decisão:**  
- Utilizar Goroutines e channels com um Pool de Workers para processamento concorrente de mensagens, garantindo alto throughput e controlando a carga no banco de dados.  
- Implementar inserções em lote (bulk inserts) com a cláusula UPSERT do PostgreSQL para garantir atomicidade e idempotência. Um hash estável de cada recurso FHIR será utilizado como chave para o UPSERT.

**Justificativa:**  
Essa abordagem combinada melhora drasticamente o desempenho da ingestão ao processar múltiplos recursos em paralelo e ao inserir grandes volumes de dados de uma só vez. O UPSERT e o hashing garantem que o sistema seja robusto o suficiente para lidar com reprocessamentos e falhas sem duplicar dados.

**Consequências:**  
- **Positivas:** Alto throughput, latência reduzida e consistência de dados garantida.  
- **Negativas:** Exige monitoramento cuidadoso do paralelismo para evitar sobrecarga no banco de dados.

---

## ADR-003: Multi-Tenant

**Contexto:**  
O serviço deve suportar múltiplos clientes (tenants) de forma segura e isolada, compartilhando a infraestrutura de maneira eficiente.

**Decisão:**  
Implementar a arquitetura multi-tenant com Row-Level Security (RLS). As tabelas do banco de dados serão compartilhadas, mas cada linha será associada a um `tenant_id` obrigatório.

**Justificativa:**  
Essa abordagem otimiza o uso de recursos e simplifica a manutenção (não é necessário gerenciar múltiplos bancos de dados). O RLS garante que as consultas de um tenant só possam acessar os dados a ele pertencentes, mantendo a segurança e o isolamento dos dados.

**Consequências:**  
- **Positivas:** Redução de custos e simplificação da administração.  
- **Negativas:** Todas as consultas e operações de repositório devem obrigatoriamente validar o `tenant_id`, exigindo atenção extra no desenvolvimento.

---

## ADR-004: Resiliência e Recuperação de Falhas

**Contexto:**  
O serviço de ingestão depende de APIs externas que podem falhar e utiliza um sistema de filas para processamento assíncrono. O sistema precisa ser autônomo na recuperação de falhas e garantir o reprocessamento de mensagens sem perda de dados.

**Decisão:**  
- Implementar retry com exponential backoff e circuit breaker para todas as chamadas HTTP a APIs externas.  
- Utilizar filas SQS com Dead-Letter Queues (DLQ) e alertas automáticos.

**Justificativa:**  
A combinação de retry e circuit breaker protege a aplicação contra falhas intermitentes e evita sobrecarga de serviços externos. O uso de DLQs garante que mensagens com falhas permanentes não sejam perdidas e possam ser inspecionadas e reprocessadas manualmente, assegurando a integridade do dado.

**Consequências:**  
- **Positivas:** Maior estabilidade, confiabilidade e proteção do sistema.  
- **Negativas:** Configuração precisa de timeouts e thresholds para circuit breakers e DLQs.

---

## ADR-005: Observabilidade e Telemetria

**Contexto:**  
A operação proativa do serviço de ingestão requer visibilidade completa em tempo real sobre seu estado, desempenho e comportamento.

**Decisão:**  
Adotar uma estratégia de observabilidade abrangente, incluindo Logs Estruturados (JSON), Métricas Customizadas (Prometheus) e Tracing Distribuído (OpenTelemetry). Todas as telemetrias serão enviadas para um painel de monitoramento centralizado (Datadog).

**Justificativa:**  
Essa decisão permite a criação de dashboards completos, a configuração de alertas automáticos e a auditoria detalhada das operações. O tracing distribuído é fundamental para depurar gargalos de performance em um sistema distribuído.

**Consequências:**  
- **Positivas:** Operação proativa, tempo de resposta a incidentes reduzido e capacidade de rastrear latência e throughput.  
- **Negativas:** Exige o desenvolvimento de middlewares e coletores, além de garantir que dados sensíveis (PHI) sejam mascarados nos logs.

---

## ADR-006: Segurança e Conformidade HIPAA

**Contexto:**  
O serviço de ingestão lida com dados de saúde sensíveis (PHI), exigindo a máxima confidencialidade, integridade e rastreabilidade para garantir a conformidade com a HIPAA.

**Decisão:**  
- Todos os dados sensíveis serão criptografados em repouso (at-rest) usando KMS (Key Management Service).  
- Nenhum dado sensível (PHI) será exposto em logs ou dashboards. Um processo de mascaramento será aplicado a todas as telemetrias.  
- Toda operação crítica será registrada em um sistema de Audit Logging para garantir rastreabilidade e auditoria.

**Justificativa:**  
Essas medidas garantem o cumprimento dos rigorosos requisitos da HIPAA, protegendo a privacidade dos dados e minimizando os riscos de vazamento de informações.

**Consequências:**  
- **Positivas:** Conformidade regulatória, alta segurança e sistema robusto contra ameaças.  
- **Negativas:** Implementação da criptografia, mascaramento e auditoria exige cuidado e atenção aos detalhes.

---

## ADR-007: Infraestrutura e Deployment

**Contexto:**  
Precisamos de uma infraestrutura que suporte escalabilidade, alta disponibilidade e operação confiável com o mínimo de intervenção manual.

**Decisão:**  
Utilizar ECS Fargate para orquestração de containers, Terraform para Infrastructure as Code (IaC), SQS para filas, RDS multi-tenant para o banco de dados e KMS para gerenciamento de chaves.

**Justificativa:**  
Essa arquitetura gerenciada e automatizada permite o auto-escalonamento do serviço com base na carga, elimina a necessidade de gerenciamento de servidores e garante a consistência do ambiente através do IaC.

**Consequências:**  
- **Positivas:** Redução da complexidade operacional, maior uptime e capacidade de escalar rapidamente em resposta à demanda.  
- **Negativas:** Custos de infraestrutura podem ser maiores do que em soluções dedicadas, mas são compensados pela redução no custo operacional.

