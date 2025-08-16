# Release 3 — API e Observabilidade Completa

**Objetivo:**  
Expor a lógica de negócio do serviço através de uma API robusta, segura e totalmente instrumentada. Esta release foca na interface pública do serviço e em sua capacidade de ser monitorado em produção.

**Definição de "Concluído":**  
A API está funcional, com todos os **middlewares essenciais** implementados, e o serviço possui **observabilidade completa**, permitindo um monitoramento eficaz em ambiente de produção.

---

## Task 3.1 — Controllers / Handlers da API

**Branches Sugeridas:**
- `feat/r3-t3.1-patient-controller`
- `feat/r3-t3.1-error-handling-and-validation`

**Descrição:**  
Criar o controller para gerenciar as requisições do endpoint `/patient`. Implementar a **validação de parâmetros** de query e um **sistema padronizado de tratamento de erros**, garantindo que as mensagens de falha sejam claras e informativas.

**Critérios de Aceitação:**
- O endpoint deve processar requisições válidas com sucesso e retornar **status code 200**.
- Entradas inválidas (ex: parâmetros ausentes ou mal formatados) devem retornar **status code 4xx** com um corpo de resposta padronizado contendo a mensagem de erro.

---

## Task 3.2 — Middlewares Essenciais

**Branches Sugeridas:**
- `feat/r3-t3.2-middlewares`

**Descrição:**  
Implementar os **middlewares essenciais** para garantir a resiliência e a observabilidade da API. Os middlewares devem ser responsáveis por:  
- Logging de requisições  
- Coleta de métricas (Prometheus)  
- Rate limiting  
- Circuit breaker  

**Critérios de Aceitação:**
- Cada middleware deve ser testado individualmente (testes unitários) para garantir seu comportamento isolado.
- Os middlewares devem ser injetados na **cadeia de execução da API**, garantindo que cada requisição passe por todos os componentes na ordem correta.

---

## Task 3.3 — Integração e End-to-End

**Branches Sugeridas:**
- `feat/r3-t3.3-di-setup`
- `feat/r3-t3.3-e2e-pipeline`

**Descrição:**  
Finalizar a configuração da **Injeção de Dependências (DI)** para conectar os componentes da aplicação. Montar e testar o **pipeline completo da requisição**:  
`HTTP → Controller → UseCase → Adapter → Response`.

**Critérios de Aceitação:**
- A injeção de dependências deve ocorrer **sem erros em tempo de execução**.
- Testes de integração (end-to-end) devem validar o **fluxo completo**, utilizando **mocks** para simular serviços externos.

---

## Task 3.4 — Finalização e Release

**Branches Sugeridas:**
- `feat/r3-t3.4-observability-docs`
- `feat/r3-t3.4-release-v1`

**Descrição:**  
Expor o endpoint `/metrics` para que o Prometheus possa coletar dados. Finalizar a documentação do projeto, incluindo **diagramas de arquitetura**, e marcar a primeira versão de produção com a **tag v1.0.0**.

**Critérios de Aceitação:**
- O endpoint `/metrics` deve retornar as métricas configuradas de **latência, erros e taxa de transferência**.
- O `README.md` deve ser atualizado com instruções detalhadas sobre execução e monitoramento do serviço.
- A **tag v1.0.0** deve ser criada e enviada para o repositório principal, marcando a conclusão do projeto.
