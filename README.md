# Backend - Chá de Casa Nova

API REST desenvolvida em Go (Gin) para gerenciar lista de presentes, confirmações de presença e mensagens para chá de casa nova.

## 🚀 Tecnologias

-   [Go](https://golang.org/)
-   [Gin](https://gin-gonic.com/)
-   [PostgreSQL](https://www.postgresql.org/) (via [Supabase](https://supabase.com/))
-   [GORM](https://gorm.io/) - ORM para Go

## 📋 Pré-requisitos

-   Go 1.16 ou superior
-   Docker e Docker Compose (opcional)
-   Conta no Supabase (para banco de dados)

## 🔧 Instalação

1. Clone o repositório
2. Entre na pasta do backend:

```bash
cd backend
```

3. Instale as dependências:

```bash
go mod download
```

4. Crie um arquivo `.env` baseado no exemplo:

```bash
cp .env.example .env
```

## 🏃‍♂️ Rodando o projeto

### Localmente

```bash
go run main.go
```

### Com Docker

```bash
docker-compose up --build
```

O servidor estará disponível em `http://localhost:8080`

## 🛣️ Rotas da API

### Públicas

-   `GET /api/v1/items` - Lista todos os itens
-   `GET /api/v1/items/:id` - Obtém um item específico
-   `GET /api/v1/stats` - Obtém estatísticas gerais
-   `GET /api/v1/evento` - Obtém informações do evento
-   `POST /api/v1/mensagens` - Cria uma nova mensagem
-   `GET /api/v1/mensagens` - Lista mensagens aprovadas
-   `POST /api/v1/confirmacoes` - Cria uma confirmação de presença
-   `POST /api/v1/items/:id/resgate` - Resgata um item
-   `POST /api/v1/items/:id/cancela-resgate` - Cancela o resgate de um item

### Administrativas (requer autenticação)

-   `GET /api/v1/admin/items` - Lista todos os itens (admin)
-   `POST /api/v1/admin/items` - Cria um novo item
-   `PUT /api/v1/admin/items/:id` - Atualiza um item
-   `DELETE /api/v1/admin/items/:id` - Remove um item
-   `GET /api/v1/admin/mensagens` - Lista todas as mensagens
-   `DELETE /api/v1/admin/mensagens/:id` - Remove uma mensagem
-   `POST /api/v1/admin/mensagens/:id/aprovar` - Aprova uma mensagem
-   `GET /api/v1/admin/confirmacoes` - Lista todas as confirmações
-   `PUT /api/v1/admin/confirmacoes/:id` - Atualiza uma confirmação
-   `DELETE /api/v1/admin/confirmacoes/:id` - Remove uma confirmação
-   `GET /api/v1/admin/stats/detalhadas` - Obtém estatísticas detalhadas
-   `PUT /api/v1/admin/evento` - Atualiza informações do evento

## 🔐 Autenticação

Para rotas administrativas, é necessário enviar o token de autenticação no header:

```
Authorization: Bearer <admin_password>
```

## 🔧 Variáveis de Ambiente

-   `PORT` - Porta do servidor (padrão: 8080)
-   `ADMIN_PASSWORD` - Senha para acesso administrativo
-   `DB_HOST` - Host do banco de dados Supabase
-   `DB_PORT` - Porta do banco de dados (padrão: 5432)
-   `DB_USER` - Usuário do banco de dados
-   `DB_PASSWORD` - Senha do banco de dados
-   `DB_NAME` - Nome do banco de dados

## 📦 Estrutura do Projeto

```
backend/
├── controllers/     # Controladores da API
├── database/       # Configuração do banco de dados
├── models/         # Modelos de dados
├── data/          # Dados estáticos
├── main.go        # Ponto de entrada da aplicação
└── docker-compose.yml
```
