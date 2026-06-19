# Gopportunities API

API REST em Go para cadastrar, listar, consultar, atualizar e remover oportunidades de emprego.

## Tecnologias

- Go
- Gin
- GORM
- PostgreSQL
- Swagger / Swaggo
- Docker Compose

## Requisitos

- Go compativel com a versao definida em `go.mod`
- Docker e Docker Compose, para subir o PostgreSQL local
- Opcional: `swag`, para regenerar a documentacao Swagger

## Configuracao

Crie o arquivo `.env` a partir do exemplo:

```powershell
Copy-Item .env.example .env
```

Valores padrao usados pelo projeto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=app
DB_SSLMODE=disable
```

O `docker-compose.yaml` ja sobe um PostgreSQL com esses mesmos dados de conexao.

## Como executar

Suba o banco de dados:

```bash
docker compose up -d postgres
```

Instale as dependencias:

```bash
go mod download
```

Execute a API:

```bash
go run main.go
```

Ou use o Makefile:

```bash
make run
```

A API ficara disponivel em:

```text
http://localhost:8080
```

Ao iniciar, a aplicacao usa o GORM `AutoMigrate` para criar/atualizar a tabela `openings`. Os arquivos SQL em `migrations/` tambem documentam a estrutura esperada da tabela.

## Swagger

Com a API em execucao, acesse:

```text
http://localhost:8080/swagger/index.html
```

Para regenerar os arquivos em `docs/`:

```bash
make docs
```

Esse comando requer o CLI `swag` instalado.

## Recursos da API

Base URL local:

```text
http://localhost:8080/api/v1
```

### Modelo de oportunidade

Campos aceitos no corpo JSON:

| Campo | Tipo | Obrigatorio no cadastro | Descricao |
| --- | --- | --- | --- |
| `role` | string | Sim | Cargo da oportunidade |
| `company` | string | Sim | Empresa contratante |
| `location` | string | Sim | Localizacao da vaga |
| `remote` | boolean | Sim | Indica se a vaga e remota |
| `link` | string | Sim | Link da vaga |
| `salary` | integer | Sim | Salario da vaga |

## Endpoints

| Metodo | Rota | Descricao |
| --- | --- | --- |
| `POST` | `/api/v1/opening` | Cria uma oportunidade |
| `GET` | `/api/v1/openings` | Lista todas as oportunidades |
| `GET` | `/api/v1/opening?id={id}` | Busca uma oportunidade por ID |
| `PUT` | `/api/v1/opening?id={id}` | Atualiza uma oportunidade por ID |
| `DELETE` | `/api/v1/opening?id={id}` | Remove uma oportunidade por ID |

### Criar oportunidade

```http
POST /api/v1/opening
Content-Type: application/json
```

Body:

```json
{
  "role": "Backend Developer",
  "company": "Acme",
  "location": "Sao Paulo",
  "remote": true,
  "link": "https://example.com/jobs/backend",
  "salary": 9000
}
```

Exemplo com cURL:

```bash
curl -X POST http://localhost:8080/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Backend Developer",
    "company": "Acme",
    "location": "Sao Paulo",
    "remote": true,
    "link": "https://example.com/jobs/backend",
    "salary": 9000
  }'
```

### Listar oportunidades

```http
GET /api/v1/openings
```

Exemplo com cURL:

```bash
curl http://localhost:8080/api/v1/openings
```

### Buscar oportunidade por ID

```http
GET /api/v1/opening?id=1
```

Exemplo com cURL:

```bash
curl "http://localhost:8080/api/v1/opening?id=1"
```

### Atualizar oportunidade

```http
PUT /api/v1/opening?id=1
Content-Type: application/json
```

Body com os campos que deseja atualizar:

```json
{
  "role": "Senior Backend Developer",
  "salary": 12000,
  "remote": true
}
```

Exemplo com cURL:

```bash
curl -X PUT "http://localhost:8080/api/v1/opening?id=1" \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Senior Backend Developer",
    "salary": 12000,
    "remote": true
  }'
```

### Remover oportunidade

```http
DELETE /api/v1/opening?id=1
```

Exemplo com cURL:

```bash
curl -X DELETE "http://localhost:8080/api/v1/opening?id=1"
```

## Formato das respostas

Respostas de sucesso retornam HTTP `200`:

```json
{
  "message": "operation from handler create-opening successfull",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-06-19T08:00:00Z",
    "UpdatedAt": "2026-06-19T08:00:00Z",
    "DeletedAt": null,
    "Role": "Backend Developer",
    "Company": "Acme",
    "Location": "Sao Paulo",
    "Remote": true,
    "Link": "https://example.com/jobs/backend",
    "Salary": 9000
  }
}
```

Respostas de erro seguem o formato atual implementado em `handler/response.go`:

```json
{
  "messagg": "param: id (type: queryParameter) is required",
  "errorCode": 400
}
```

Possiveis codigos de erro:

| Codigo | Quando acontece |
| --- | --- |
| `400` | Corpo invalido ou parametro obrigatorio ausente |
| `404` | Oportunidade nao encontrada |
| `500` | Erro interno ou erro ao acessar o banco |

## Comandos uteis

```bash
make run      # executa a API
make build    # gera o binario gopportunities
make docs     # regenera a documentacao Swagger
go test ./... # executa os testes
```

## Colecoes de API

As colecoes para testar os endpoints estao em:

```text
docs/api-collections/
```
