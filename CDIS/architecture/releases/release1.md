# Release 1 — Setup Base e Infraestrutura

**Objetivo:**  
Estabelecer a base do projeto, incluindo a estrutura de arquivos, o pipeline de CI/CD e os componentes essenciais de configuração e logging.

**Definição de "Concluído":**  
O repositório está configurado, o pipeline de CI/CD está funcionando e os componentes de configuração e logging estão prontos para uso, com testes unitários que garantem a qualidade do código.

---

## Task 1.1 — Estrutura Base do Projeto

**Branches Sugeridas:**
- `feat/r1-t1.1-setup-folders`
- `feat/r1-t1.1-init-mod`
- `feat/r1-t1.1-gitignore-readme-ci`

**Descrição:**  
Criar a estrutura padrão de diretórios (`cmd`, `internal`, `pkg`, `__tests__`), inicializar o módulo Go, configurar o `.gitignore` para ignorar arquivos desnecessários, criar o `README.md` inicial e configurar o workflow de GitHub Actions para testes e linting.

**Critérios de Aceitação:**
- As pastas e arquivos padrão devem estar na raiz do projeto.
- O comando `go mod tidy` deve ser executado sem erros.
- O workflow do GitHub Actions deve executar `go test ./...` e `go vet ./...` com sucesso.

---

## Task 1.2 — Configuração de Ambiente

**Branches Sugeridas:**
- `feat/r1-t1.2-config-helpers`
- `feat/r1-t1.2-load-func`
- `feat/r1-t1.2-config-tests`

**Descrição:**  
Implementar a lógica para carregar as configurações do ambiente. Seguir a metodologia **TDD (Test-Driven Development)**: escrever testes antes de cada implementação de código para garantir o comportamento esperado.

**Critérios de Aceitação:**
- A função `Load()` deve carregar variáveis de ambiente com sucesso, aplicando valores padrão quando necessário.
- A função deve retornar um erro explícito se uma variável obrigatória estiver ausente.
- Testes unitários devem cobrir todas as funções auxiliares e a função principal de carregamento.

---

## Task 1.3 — Logger Estruturado

**Branches Sugeridas:**
- `feat/r1-t1.3-logger-interface-impl`
- `feat/r1-t1.3-logger-methods`
- `feat/r1-t1.3-logger-tests`

**Descrição:**  
Definir a interface `Logger` para desacoplamento, criar a implementação concreta `StructuredLogger` para logging em formato JSON e adicionar os métodos `Debug`, `Info`, `Warn`, `Error` e `Fatal`.

**Critérios de Aceitação:**
- A interface e a implementação devem estar claras e seguir os padrões de programação em Go.
- O logger deve formatar as mensagens em JSON para facilitar a análise em sistemas de logging centralizados.
- Testes unitários devem validar a saída do logger, os diferentes níveis de log e os campos estruturados.
