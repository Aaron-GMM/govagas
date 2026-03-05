# 💼 Govagas Microservice

<p align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Gin-008080?style=flat&logo=gin&logoColor=white" />
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white" />
  <img src="https://img.shields.io/badge/GORM-blue?style=flat" />
</p>

O **Govagas** é um microsserviço para gerenciamento de oportunidades de emprego. Ele funciona como o core de um ecossistema de prova de conceito (PoC), comunicando-se com outros serviços como Auth-JWT e Email.

O projeto foi recentemente refatorado para seguir os princípios de **Clean Architecture** e **Injeção de Dependências**, garantindo um código altamente escalável, testável e de fácil manutenção, além de rodar em um ambiente totalmente containerizado.

## 🛠️ Tecnologias e Dependências

* **Linguagem:** Go (v1.25.5)
* **Roteamento:** [Gin Web Framework](https://github.com/gin-gonic/gin)
* **Persistência:** PostgreSQL (migrado de SQLite para maior robustez)
* **ORM:** [GORM](https://gorm.io/)
* **Infraestrutura:** Docker e Docker Compose (Multi-stage builds)
* **Configuração:** `godotenv` para gerenciamento de variáveis de ambiente

## 🏗️ Arquitetura e Estrutura (Clean Architecture)

O projeto utiliza uma estrutura modular baseada em camadas para isolar responsabilidades e aplicar injeção de dependências:

* `router/`: Configuração de grupos de rotas e orquestração de dependências.
* `handler/` (Apresentação): Gerenciamento de requisições HTTP (Gin) e formatação de respostas.
* `service/` (Regras de Negócio): Camada central isolada, responsável pelas regras da aplicação.
* `repository/` (Acesso a Dados): Abstração do banco de dados utilizando GORM.
* `schemas/`: Definições de entidades do banco e structs de requisição/resposta.
* `config/`: Inicialização de banco de dados e logger customizado.

## 📡 Endpoints Principais (API v1)

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| GET | `/api/v1/opening` | Busca uma vaga por ID (via query param `?id=`) |
| POST | `/api/v1/opening` | Cria uma nova oportunidade |
| PUT | `/api/v1/opening` | Atualiza dados de uma vaga (via query param `?id=`) |
| DELETE | `/api/v1/opening` | Remove uma vaga do sistema (via query param `?id=`) |
| GET | `/api/v1/openings` | Lista todas as vagas disponíveis |

## ⚙️ Configuração Local com Docker

A aplicação foi desenhada para rodar facilmente através do Docker, sem a necessidade de instalar o PostgreSQL localmente na sua máquina.

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/Aaron-GMM/govagas](https://github.com/Aaron-GMM/govagas)
    cd govagas
    ```

2.  **Variáveis de Ambiente:**
    Crie um arquivo `.env` na raiz do projeto e defina as credenciais do banco e a porta da API:
    ```env
    PORT=8080
    DB_HOST=db
    DB_PORT=5432
    DB_USER=admin
    DB_PASSWORD=adminpassword
    DB_NAME=govagas_db
    ```

3.  **Executar a Aplicação:**
    Levante os containers do PostgreSQL e da API utilizando o comando:
    ```bash
    docker compose up --build -d
    ```

4.  **Acompanhar os Logs (Opcional):**
    ```bash
    docker compose logs -f api
    ```

A API estará disponível e escutando em `http://localhost:8080`.

---
*Nota: Este projeto é uma Prova de Conceito (PoC) focada em boas práticas de arquitetura de software (Clean Arch, DI) no ecossistema Go.*