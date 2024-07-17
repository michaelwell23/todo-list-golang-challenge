# Dasafio ToDo-List da MoonLabs

- [Dasafio ToDo-List da MoonLabs](#dasafio-todo-list-da-moonlabs)
  - [Sobre](#sobre)
    - [Regras de Negócio](#regras-de-negócio)
    - [Requisitos não funcionais](#requisitos-não-funcionais)
  - [Como Executar a Aplicação](#como-executar-a-aplicação)
    - [Configuração e Execução](#configuração-e-execução)
    - [Proteção da aplicação](#proteção-da-aplicação)
    - [Rotas da Aplicação](#rotas-da-aplicação)
  - [Documentação](#documentação)

## Sobre

O Objetivo do projeto é desenvolver o back-end de uma lista de tarefas simples usando Golang, com um banco de dados na memória, e expor os microsserviços usando uma ferramenta OpenAPI de sua escolha.

### Regras de Negócio

- O front-end pode acessar tanto a lista de tarefas pendentes quanto as concluídas.
- Para adicionar uma nova tarefa, são necessários: descrição, responsável e e-mail.
- Tarefas pendentes podem ser marcadas como concluídas e vice-versa, com um limite máximo de 3 reversões.
- A exclusão de tarefas é permanente.

### Requisitos não funcionais

- O login não é necessário para usar o ToDo.
- Alguma forma de proteção deve ser usada para proteger os endpoints do acesso público. E para solucionar esse requisito foi adicionado um middleware de autenticação foi adicionado ao projeto.
- Os métodos HTTP correspondentes para cada operação devem ser usados.

---

## Como Executar a Aplicação

### Configuração e Execução

1. Clone o repositório:

```bash
   git clone https://github.com/seu-usuario/todo-list.git
   cd todo-list
```

2. Instale as dependências:

```bash
go mod tidy

```

3. Modifique o arquivo `.env` na raiz do projeto e adicione sua chave de API

```bash
API_KEY=mysecretkey

```

4. Compile e execute a aplicação:

```bash
go run cmd/main.go
```

---

### Proteção da aplicação

A aplicação requer uma chave de API (API_KEY) no cabeçalho das requisições para proteger os endpoints de acesso não autorizado. Certifique-se de que a chave de API configurada no arquivo `.env` corresponda à chave usada nas requisições.

### Rotas da Aplicação

A seguir estão as principais rotas da API:

- `POST` /tasks: Adicionar uma nova tarefa.
- `GET` /tasks/pending: Obter todas as tarefas pendentes.
- `GET` /tasks/completed: Obter todas as tarefas concluídas.
- `PUT` /tasks/{id}/complete: Marcar uma tarefa como concluída.
- `PUT` /tasks/{id}/revert: Reverter uma tarefa concluída para pendente.
- `DELETE` /tasks/{id}: Excluir uma tarefa.

O arquvivo de [rotas](/routes_todo.json) pode ser utilizados em aplicativos como Insomnia ou Postman.

---

## Documentação

Para mais detalhes sobre a API, consulte a documentação localizada no diretório [\_doc](/_doc/openapi.yaml).
