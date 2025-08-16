# Conventions & Guidelines – Clinical Data Ingestion Service (CDIS)

**Versão:** 1.0  
**Autor:** Jason Silva / Tech Lead  
**Data:** 15 de Agosto de 2025  

---

## 1. Estrutura de Branches

O repositório segue a convenção **Gitflow simplificada**:

- `main`: versão estável em produção.
- `release1/CDIS-v1`: branch de release (contém todos os códigos do dia 1).
- `task/<microtask>`: branch de tarefa dentro da release.
  - Ex.: `task/CDIS-docs`, `task/CDIS-runbook`, `task/CDIS-sdr`.
- Cada microtask deve ser criada **a partir da branch de release** correspondente.
- Exclua branches remotas obsoletas para manter o repositório limpo.

---

## 2. Naming Conventions

### Branches
<tipo>/<descrição-curta>
- Tipos: `release`, `task`, `hotfix`, `experiment`.
- Exemplo: `task/CDIS-adr-docs`, `hotfix/CDIS-dlq-fix`.

### Commits (Conventional Commits)
<tipo>(<escopo>): <descrição>
- Tipos comuns: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`.
- Escopo: módulo ou microtask.
- Exemplo: `docs(adr): add initial ADRs v1.0 for Clinical Data Ingestion Service`
- Sempre escrever **imperativo e claro**.

### Tags
- `vX.Y.Z` para releases.
- Exemplo: `v1.0.0`, `v2.0.0-beta`.

---

## 3. Code Style

### Python
- PEP8 com `black` + `isort`.
- Tipagem obrigatória com `mypy`.
- Estrutura de diretórios clara: `handlers/`, `services/`, `shared/`, `utils/`.

### Go
- `gofmt` obrigatório.
- Interfaces pequenas e explícitas.
- Evitar dependência circular entre pacotes.

### JavaScript / React
- ESLint + Prettier.
- Componentes funcionais sempre que possível.
- Hooks seguindo regras padrão.

---

## 4. Pull Requests / Code Review

- PR deve ser criado a partir da branch de task.
- Checklist PR:
  - [ ] Código testado com TDD.
  - [ ] Testes unitários >= 80% coverage.
  - [ ] Código revisado por pelo menos 1 colega.
  - [ ] ADR/Runbook/Docs atualizados se houver mudanças arquiteturais.
  - [ ] Nenhum segredo ou credencial exposta.

- PRs devem ter **descrição clara**, referenciando ticket ou microtask.

---

## 5. Observabilidade e Logging

- Logs estruturados (JSON), usando `slog` ou equivalente.
- PHI nunca deve ser logado sem mascaramento.
- Métricas devem ser enviadas para Prometheus/Datadog:
  - Throughput
  - Latência
  - Failures / Errors
- Circuit Breakers e DLQs monitorados com alertas críticos.

---

## 6. Segurança e Compliance

- Todos os dados PHI devem ser criptografados (at-rest e in-transit).
- Audit logs obrigatórios para operações críticas.
- Follow HIPAA compliance checklist.
- Nunca commitar credenciais, tokens ou chaves.

---

## 7. Procedimentos Extras

- Scripts de deploy e migração devem estar versionados.
- Testes de integração devem rodar em ambiente isolado antes do merge.
- Documentos ADR, Runbooks e SDRs sempre versionados e revisados.

---

> Seguindo essas convenções, garantimos código limpo, rastreabilidade de decisões, compliance com HIPAA e facilidade de manutenção em um sistema escalável.

