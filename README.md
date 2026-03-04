# 💼 Govagas Microservice

<p align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Gin-008080?style=flat&logo=gin&logoColor=white" />
  <img src="https://img.shields.io/badge/SQLite-003B57?style=flat&logo=sqlite&logoColor=white" />
  <img src="https://img.shields.io/badge/GORM-blue?style=flat" />
</p>

O **Govagas** é um microsserviço para gerenciamento de oportunidades de emprego. Ele funciona como o core de um ecossistema de prova de conceito (PoC), comunicando-se com outros dois serviços:
1. **Auth-JWT Service**: Para validação de tokens e proteção de rotas.
2. **Email Service**: Para notificações de novas vagas e confirmações.

## 🛠️ Tecnologias e Dependências

* **Linguagem:** Go (v1.25.5)
* **Roteamento:** [Gin Web Framework](https://github.com/gin-gonic/gin)
* **Persistência:** SQLite (para desenvolvimento ágil)
* **ORM:** [GORM](https://gorm.io/)
* **Configuração:** `godotenv` para gerenciamento de variáveis de ambiente

## 🏗️ Arquitetura e Estrutura

O projeto utiliza uma estrutura modular para facilitar a manutenção:
* `handler/`: Gerenciamento de requisições e lógica de entrada/saída.
* `schemas/`: Definições de entidades de banco de dados (GORM Models).
* `config/`: Inicialização de banco de dados e logger customizado.
* `router/`: Configuração de grupos de rotas e middlewares.

## 📡 Endpoints Principais (API v1)

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| GET | `/api/v1/opening` | Busca uma vaga por ID |
| POST | `/api/v1/opening` | Cria uma nova oportunidade |
| PUT | `/api/v1/opening` | Atualiza dados de uma vaga |
| DELETE | `/api/v1/opening` | Remove uma vaga do sistema |
| GET | `/api/v1/openings` | Lista todas as vagas disponíveis |

## ⚙️ Configuração Local

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/Aaron-GMM/govagas](https://github.com/Aaron-GMM/govagas)
    ```
2.  **Variáveis de Ambiente:**
    Crie um arquivo `.env` na raiz (ignorado pelo Git) e defina:
    ```env
    DB_PATH=./db/main.db
    PORT=8080
    ```
3.  **Executar:**
    ```bash
    go run main.go
    ```

---
*Nota: Este projeto é uma Prova de Conceito (PoC) e está em desenvolvimento ativo.*