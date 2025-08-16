# Conventions & Guidelines – Clinical Data Ingestion Service (CDIS)

**Versão:** 1.0  
**Autor:** Jason Silva / Principal Engineer  
**Data:** 15 de Agosto de 2025  

---

## 1. Estrutura de Branches

O repositório segue a convenção **Gitflow simplificada**:

- `main`: versão estável em produção.
- `release1/CDIS-v1`: branch de release (contém todos os códigos do dia 1).
- `release2/CDIS-v2`: branch de release do dia 2, etc.
- `task/<microtask>`: branch de tarefa dentro da release correspondente.  
  - Ex.: `task/CDIS-docs`, `task/CDIS-runbook`, `task/CDIS-sdr`.
- Cada microtask deve ser criada **a partir da branch de release correspondente**.
- Exclua branches remotas obsoletas para manter o repositório limpo.

---

## 2. Naming Conventions

### 2.1 Branches
Formato: `tipo/descrição-curta`  
Tipos: `release`, `task`, `hotfix`, `experiment`  
Exemplos:
- `feat/r1-t1.2-config-helpers`
- `feat/r1-t1.2-load-func`
- `feat/r1-t1.2-config-tests`

### 2.2 Commits (Conventional Commits)
Formato: `tipo(escopo): descrição`  
Tipos comuns: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`  
Escopo: módulo ou microtask  
Exemplo:
- `docs(adr): added initial ADRs v1.0 for Clinical Data Ingestion Service`
- `feat(logger): implemented Debug, Info, Warn, Error`
- `test(fhir-adapter): added retries unit tests`

### 2.3 Tags
- `vX.Y.Z` para releases.
- Ex.: `v1.0.0`, `v2.0.0-beta`.

---

## 3. Estrutura de Código e Estilo

### 3.1 Go
- `gofmt` obrigatório.
- Interfaces pequenas e explícitas.
- Evitar dependência circular entre pacotes.
- Estrutura típica:
  - `cmd/` → executáveis
  - `internal/` → código interno da aplicação
  - `pkg/` → pacotes reutilizáveis
  - `tests/` → testes unitários
- TDD obrigatório.

### 3.2 Observabilidade e Logging
- Logs estruturados (JSON) com `slog` ou equivalente.
- PHI nunca deve ser logado sem mascaramento.
- Métricas enviadas para Prometheus/Datadog:
  - Throughput
  - Latência
  - Failures / Errors
- Circuit Breakers e DLQs monitorados com alertas críticos.

### 3.3 Segurança e Compliance
- Todos os dados PHI criptografados (at-rest e in-transit).
- Audit logs obrigatórios para operações críticas.
- Seguir checklist HIPAA.
- Nunca commitar credenciais, tokens ou chaves.

### 3.4 Procedimentos Extras
- Scripts de deploy e migração versionados.
- Testes de integração devem rodar em ambiente isolado antes do merge.
- Documentos ADR, Runbooks e SDRs sempre versionados e revisados.

---

## 4. Pull Requests / Code Review

- PR deve ser criado a partir da branch da **microtask**.
- Descrição clara, referenciando ticket ou microtask.
- Checklist PR:
  - [ ] Código testado com TDD.
  - [ ] Testes unitários >= 80% coverage.
  - [ ] Revisão de pelo menos 1 colega.
  - [ ] ADR/Runbook/Docs atualizados se houver mudanças arquiteturais.
  - [ ] Nenhum segredo ou credencial exposta.

---

## 5. Processo de Tasks, Subtasks e Branches

### 5.1 Criando Tasks no ClickUp
1. Criar **Task** com nome claro.
2. Associar **Release** (ex.: Release 1) e **Módulo** (Backend, Infra, Observability, etc.).
3. Etiquetas sugeridas: `TDD`, `CleanCode`, `Adapter`, `UseCase`.
4. Status inicial: `Todo`.
5. Descrição: objetivos, critérios de aceitação e referência de arquitetura.

### 5.2 Criando Subtasks / Microtasks
- Quebrar tasks em microtasks granulares, ex.:
  - `Logger Estruturado` → `Definir interface Logger`, `Implementar métodos`, `Testes unitários`.
- Cada subtask tem **status independente** (`Todo`, `In Progress`, `Concluído`), mas **permanece vinculada à task principal**.
- Microtask iniciada → mover para `In Progress`.
- Task principal só vai para `Concluído` quando todas as subtasks forem finalizadas.

### 5.3 Branches para Microtasks
- Criar branch a partir da **release correspondente**:
  - git checkout release1/CDIS-v1
  - git checkout -b task/CDIS-logger-interface
- Branch nome: task/<microtask>.
- Cada branch corresponde a uma subtask.

### 5.4 Commits – Conventional Commits

- Mensagens imperativas, escopo claro:

  - feat(logger): define Logger interface
  - feat(logger): implement Debug, Info, Warn, Error
  - test(logger): add unit tests


### 5.5 Pull Requests (PR)

PR: branch da subtask → branch da release.
Descrição detalhando implementação.
Checklist PR:

 - Código testado com TDD
 - Cobertura >= 80%
 - Revisão de pelo menos 1 colega
 - ADR / Docs atualizados

## 6. Exemplo Prático de Task / Subtasks
| Task | Subtask | Branch | Commit | PR | Status |
|------|---------|--------|--------|----|--------|
| Logger Estruturado | Definir interface Logger | task/CDIS-logger-interface | feat(logger): define Logger interface | task/CDIS-logger-interface → release1/CDIS-v1 | In Progress |
| Logger Estruturado | Implementar métodos Debug/Info/Warn/Error | task/CDIS-logger-methods | feat(logger): implement logging methods | task/CDIS-logger-methods → release1/CDIS-v1 | Todo |
| Logger Estruturado | Implementar StructuredLogger JSON | task/CDIS-logger-json | feat(logger): implement StructuredLogger JSON | task/CDIS-logger-json → release1/CDIS-v1 | Todo |
| Logger Estruturado | Testes unitários TDD | task/CDIS-logger-tests | test(logger): add unit tests | task/CDIS-logger-tests → release1/CDIS-v1 | Todo |

Cada subtask tem progresso independente, mas está vinculada à task principal.

Seguindo estas convenções, garantimos código limpo, rastreabilidade de decisões, compliance com HIPAA e facilidade de manutenção em um sistema escalável.