# Release 2 — Domínio e Lógica de Negócio

**Objetivo:**  
Implementar o núcleo do serviço, que inclui as regras de negócio e a interação com serviços externos. A arquitetura deve ser **hexagonal**, garantindo o desacoplamento entre a lógica central e os detalhes de infraestrutura.

**Definição de "Concluído":**  
As entidades de domínio estão definidas e validadas, os casos de uso estão implementados com 100% de cobertura de testes, e todos os adaptadores estão funcionando e desacoplados da lógica de negócio.

---

## Task 2.1 — Entidades de Domínio

**Branches Sugeridas:**
- `feat/r2-t2.1-entities`
- `feat/r2-t2.1-validations-factory`

**Descrição:**  
Criar as structs que representam as entidades de negócio (ex: `Patient`, `FHIRResource`). Adicionar validações básicas embutidas nas structs para garantir a integridade dos dados. Implementar um **Factory Pattern** para centralizar a lógica de criação e validação das instâncias.

**Critérios de Aceitação:**
- As entidades de domínio devem ser simples e **não conter lógica de infraestrutura**.
- O Factory deve encapsular a lógica de criação e retornar erros de validação claros e específicos, garantindo que objetos inválidos não sejam criados.

---

## Task 2.2 — Casos de Uso (UseCases)

**Branches Sugeridas:**
- `feat/r2-t2.2-interfaces-usecase`

**Descrição:**  
Definir as interfaces (também conhecidas como **Ports**) para os adaptadores externos (ex: `FHIRClient`, `DBRepository`, `CryptoService`). Implementar o caso de uso principal `ProcessFHIRResourceUseCase`, seguindo estritamente a metodologia **TDD**.

**Critérios de Aceitação:**
- O UseCase deve conter apenas **lógica de negócio pura**, sem qualquer dependência direta de implementações concretas de adaptadores.
- Os testes unitários devem cobrir todos os cenários de sucesso e falha, utilizando **mocks** para simular o comportamento das interfaces.

---

## Task 2.3 — Adapters Externos

**Branches Sugeridas:**
- `feat/r2-t2.3-fhir-adapter`
- `feat/r2-t2.3-db-adapter`
- `feat/r2-t2.3-crypto-adapter`

**Descrição:**  
Implementar a lógica para interagir com os serviços externos. Estes componentes são os **Adapters** que se comunicam com o mundo exterior.

**Critérios de Aceitação:**
- **FHIR Adapter:** Deve incorporar lógica de **retries com exponential backoff** e timeouts para resiliência. Os testes devem usar `httptest` para simular as respostas da API externa.
- **DB Adapter:** Deve implementar a conexão com o banco de dados e gerenciar o **connection pooling**. Os testes unitários devem garantir a correta interação via mocks, sem necessidade de um banco real.
- **Crypto Adapter:** Deve realizar operações de `encrypt`, `decrypt`, `hash` e gerenciar chaves. Os testes unitários devem cobrir todas as operações para garantir a segurança.

---

## Task 2.4 — Decorators para Observabilidade

**Branches Sugeridas:**
- `feat/r2-t2.4-observability-decorators`

**Descrição:**  
Criar **wrappers (decorators)** que envolvem os casos de uso para adicionar funcionalidades de observabilidade (**logs e métricas**) de forma transparente, sem modificar a lógica de negócio.

**Critérios de Aceitação:**
- Os decorators devem adicionar **logs de entrada e saída** das funções, bem como métricas de **tempo de execução**, sem alterar o comportamento ou a assinatura original dos casos de uso.
