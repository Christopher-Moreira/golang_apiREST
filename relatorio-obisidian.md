2026-06-07 02:19

Status: #adult

Tags: [[golang]] [[api]] [[programação]]

---

# Decorrer

Este projeto implementa a base de uma [[API REST]] escrita em [[Golang]] para gerenciamento de vagas de emprego, representadas pelo modelo `Opening`.

A aplicação utiliza o [[framework Gin]] para expor rotas [[HTTP]], o [[GORM]] para comunicação com o banco de dados e o [[PostgreSQL]] como banco relacional.

A ideia principal do projeto é construir uma API simples, organizada em camadas, seguindo uma estrutura comum em aplicações Go:

* `main`: ponto de entrada da aplicação;
* `config`: configurações gerais, banco de dados e logger;
* `router`: definição das rotas HTTP;
* `handler`: funções responsáveis por lidar com as requisições;
* `schemas`: modelos e estruturas relacionadas ao banco de dados.

Durante o desenvolvimento, o projeto foi executado localmente com:

```bash
go run main.go
```

Já o comando:

```bash
go build
```

é utilizado para compilar a aplicação e gerar um binário executável.

---

# Estrutura inicial da aplicação

Inicialmente foi criado um módulo Go para a aplicação, utilizando o arquivo `go.mod`.

```go
module github.com/Christopher-Moreira/golang_apiREST
```

Esse arquivo define o nome do módulo, a versão da linguagem e permite que o Go gerencie as dependências externas do projeto, como Gin, GORM e o driver do PostgreSQL.

No arquivo `main.go` ficou concentrada a inicialização principal da aplicação, chamando as configurações necessárias e subindo o servidor HTTP.

Também foram separados arquivos e pacotes específicos para melhorar a organização do projeto, evitando deixar toda a lógica concentrada em um único arquivo.

---

# Servidor HTTP com Gin

Foi utilizado o [[Gin]] para criar o servidor HTTP da aplicação.

A API foi configurada para rodar na porta `:8080`, expondo as rotas principais do sistema.

Também foi criado um grupo de rotas versionado em:

```txt
/api/v1
```

A partir desse grupo, foram declaradas as rotas principais relacionadas ao CRUD de vagas:

```txt
GET    /api/v1/opening
POST   /api/v1/opening
PUT    /api/v1/opening/:id
DELETE /api/v1/opening/:id
GET    /api/v1/openings
```

Essas rotas representam as operações básicas da API:

* criar uma vaga;
* buscar uma vaga;
* listar vagas;
* atualizar uma vaga;
* remover uma vaga.

Para testar e documentar as requisições HTTP, foi utilizado o [[Bruno]], funcionando como alternativa ao Postman.

---

# Integração com PostgreSQL usando GORM

Depois da estrutura inicial da API, foi configurada a conexão com o banco [[PostgreSQL]] utilizando o [[GORM]].

O GORM permite trabalhar com o banco de dados usando structs Go, facilitando operações como criação, busca, atualização e remoção de registros.

Também foi utilizado o recurso de [[AutoMigrate]], que permite criar ou atualizar automaticamente a tabela relacionada ao modelo `Opening`.

Além disso, foi criado um arquivo `docker-compose.yaml` para subir um banco PostgreSQL localmente durante o desenvolvimento.

Também foi criado um arquivo `.env.example` contendo as variáveis de ambiente necessárias para a conexão com o banco:

```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
```

Para evitar versionar informações sensíveis, como credenciais reais do banco, foi criado um `.gitignore` incluindo o arquivo `.env`.

---

# Logger customizado

Também foi criado um logger próprio para padronizar as mensagens da aplicação.

Esse logger é utilizado para registrar mensagens de:

* debug;
* informação;
* aviso;
* erro.

Isso ajuda bastante durante o desenvolvimento, principalmente para entender o fluxo da aplicação, identificar erros de validação e acompanhar operações feitas na API.

---

# Modelo de dados

O modelo principal da aplicação está no arquivo:

```txt
schemas/opening.go
```

Ele representa uma vaga de emprego dentro do sistema.

```go
type Opening struct {
    gorm.Model

    Role     string
    Company  string
    Location string
    Remote   bool
    Link     string
    Salary   int64
}
```

Esse modelo possui os seguintes campos:

* `Role`: cargo da vaga;
* `Company`: empresa responsável pela vaga;
* `Location`: localização da vaga;
* `Remote`: indica se a vaga é remota;
* `Link`: link da vaga;
* `Salary`: salário da vaga.

O campo embutido `gorm.Model` adiciona automaticamente alguns campos padrão do GORM:

* `ID`;
* `CreatedAt`;
* `UpdatedAt`;
* `DeletedAt`.

Também foi criada uma estrutura `OpeningResponse`, usada para representar a resposta JSON da API de forma mais controlada, utilizando tags como:

```go
json:"role"
json:"salary"
```

---

# Conceitos de Go aplicados

## Módulos Go

O projeto utiliza o arquivo `go.mod` para declarar o módulo da aplicação e suas dependências.

```go
module github.com/Christopher-Moreira/golang_apiREST
```

Esse arquivo permite que o Go saiba qual é o módulo principal do projeto e quais bibliotecas externas precisam ser instaladas.

---

## Pacotes

O código foi dividido em pacotes para separar responsabilidades.

Exemplos de pacotes utilizados:

```go
package main
package config
package router
package handler
package schemas
```

Essa separação melhora a organização, facilita a manutenção e deixa o projeto mais próximo de uma estrutura real de API.

---

## Imports

Cada arquivo importa apenas os pacotes necessários para sua responsabilidade.

Alguns exemplos:

```go
"github.com/gin-gonic/gin"
```

Usado para criar rotas, lidar com contexto HTTP e retornar respostas JSON.

```go
"gorm.io/gorm"
```

Usado para trabalhar com o ORM.

```go
"gorm.io/driver/postgres"
```

Usado como driver de conexão com PostgreSQL.

```go
"fmt"
"os"
```

Usados para formatação de strings, erros e leitura de variáveis de ambiente.

```go
"log"
"io"
```

Usados na criação e configuração do logger.

---

# CRUD e lógica base da API

Na segunda parte do desenvolvimento, foi iniciada a criação da lógica do CRUD da API.

Para isso, foi criado um handler responsável por receber as requisições, validar os dados e interagir com o banco de dados.

Também foi criado um arquivo `request.go`, responsável por definir a estrutura dos dados esperados no corpo da requisição.

---

# Inicialização do handler

No pacote `handler`, foram declaradas variáveis globais para o logger e para a conexão com o banco.

```go
package handler

import (
    "github.com/Christopher-Moreira/golang_apiREST/config"
    "gorm.io/gorm"
)

var (
    logger *config.Logger
    db     *gorm.DB
)

func InitializeHandler() {
    logger = config.GetLogger("handler")
    db = config.GetPostgres()
}
```

A função `InitializeHandler` centraliza a inicialização das dependências usadas pelos handlers.

Com isso, os handlers conseguem acessar tanto o logger quanto a conexão com o PostgreSQL.

---

# Recebendo JSON com Gin

No Gin, para receber o corpo de uma requisição JSON, usamos o método:

```go
ctx.BindJSON(&request)
```

Esse método lê o corpo da requisição e tenta preencher a struct informada.

O primeiro teste foi feito com uma struct simples contendo apenas o campo `Role`.

```go
package handler

import (
    "github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
    request := struct {
        Role string `json:"role"`
    }{}

    ctx.BindJSON(&request)

    logger.Infof("request received: %v", request)
}
```

O payload enviado foi:

```json
{
  "role": "Senior da vida"
}
```

O retorno no log foi semelhante a:

```txt
>INFO: 2026/06/09 07:18:52 request received: {Senior da vida}
[GIN] 2026/06/09 - 07:18:52 | 200 | 1.06ms | 127.0.0.1 | POST "/api/v1/opening"
```

Com isso, foi validado que a API já conseguia receber dados via JSON no corpo da requisição.

---

# Estrutura de request

Depois do primeiro teste, foi criada uma struct específica para representar os dados esperados na criação de uma vaga.

```go
package handler

type CreateOpeningRequest struct {
    Role     string `json:"role"`
    Company  string `json:"company"`
    Location string `json:"location"`
    Remote   *bool  `json:"remote"`
    Link     string `json:"link"`
    Salary   int64  `json:"salary"`
}
```

O campo `Remote` foi definido como ponteiro para `bool`:

```go
Remote *bool `json:"remote"`
```

Isso permite diferenciar quando o valor foi enviado como `false` e quando ele simplesmente não foi enviado na requisição.

Depois disso, a função do handler passou a receber essa struct:

```go
func CreateOpeningHandler(ctx *gin.Context) {
    request := CreateOpeningRequest{}

    ctx.BindJSON(&request)

    logger.Infof("request received: %v", request)
}
```

---

# Validação dos dados

Após receber os dados, foi criada uma função `Validate` para validar os campos obrigatórios da request.

```go
func (r *CreateOpeningRequest) Validate() error {
    if r.Role == "" {
        return errParamIsRequired("role", "string")
    }

    if r.Company == "" {
        return errParamIsRequired("company", "string")
    }

    if r.Location == "" {
        return errParamIsRequired("location", "string")
    }

    if r.Link == "" {
        return errParamIsRequired("link", "string")
    }

    if r.Remote == nil {
        return errParamIsRequired("remote", "bool")
    }

    if r.Salary <= 0 {
        return errParamIsRequired("salary", "int64")
    }

    return nil
}
```

Essa função verifica se os campos obrigatórios foram enviados corretamente.

Caso algum campo esteja ausente ou inválido, a função retorna um erro indicando qual parâmetro é obrigatório.

---

# Criando uma vaga no banco

Depois da validação, os dados da request são usados para criar um registro no banco.

O ideal é converter a request para o schema `Opening`, que representa a tabela no banco de dados.

```go
opening := schemas.Opening{
    Role:     request.Role,
    Company:  request.Company,
    Location: request.Location,
    Remote:   *request.Remote,
    Link:     request.Link,
    Salary:   request.Salary,
}
```

Depois disso, o registro pode ser salvo com o GORM:

```go
if err := db.Create(&opening).Error; err != nil {
    logger.Errorf("error creating opening: %v", err.Error())
    sendError(ctx, http.StatusInternalServerError, "error creating opening")
    return
}
```

---

# Handler de criação

A função final do `CreateOpeningHandler` fica responsável por:

1. receber o JSON da requisição;
2. validar os dados;
3. criar o modelo `Opening`;
4. salvar no banco;
5. retornar uma resposta para o cliente.

```go
package handler

import (
    "net/http"

    "github.com/Christopher-Moreira/golang_apiREST/schemas"
    "github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
    request := CreateOpeningRequest{}

    if err := ctx.BindJSON(&request); err != nil {
        logger.Errorf("invalid JSON: %v", err.Error())
        sendError(ctx, http.StatusBadRequest, "invalid JSON")
        return
    }

    if err := request.Validate(); err != nil {
        logger.Errorf("validation error: %v", err.Error())
        sendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

    opening := schemas.Opening{
        Role:     request.Role,
        Company:  request.Company,
        Location: request.Location,
        Remote:   *request.Remote,
        Link:     request.Link,
        Salary:   request.Salary,
    }

    if err := db.Create(&opening).Error; err != nil {
        logger.Errorf("error creating opening: %v", err.Error())
        sendError(ctx, http.StatusInternalServerError, "error creating opening")
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "opening created successfully",
        "data":    opening,
    })
}
```

Um ponto importante é que a request não deve ser salva diretamente no banco usando:

```go
db.Create(&request)
```

Isso faria o GORM tentar criar uma tabela baseada na struct `CreateOpeningRequest`, o que não é o objetivo.

A request serve apenas para representar os dados recebidos pela API. Quem representa a tabela do banco é o schema `Opening`.

---

# Respostas da API

Para padronizar as respostas de erro, foi criado um arquivo `response.go`.

```go
package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, code int, msg string) {
    ctx.Header("Content-Type", "application/json")

    ctx.JSON(code, gin.H{
        "msg":       msg,
        "errorCode": code,
    })
}
```

Essa função facilita o retorno de erros para o cliente, mantendo um padrão nas respostas JSON.

Exemplo de resposta de erro:

```json
{
  "msg": "role is required",
  "errorCode": 400
}
```
---

# Referências
