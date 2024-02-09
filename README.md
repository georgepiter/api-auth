# :key: API-Auth - API de autenticação e validação do token JWT :lock:

## 💻 O projeto foi desenvolvido com:

- [x] GO 1.20
- [x] Token JWT
- [x] Gin 1.9.1
- [x] GORM 1.9.16
- [x] Swaggo 1.16 - Após iniciar o projeto acesse: http://localhost:8080/swagger/index.html.
- [x] Banco de dados - MYSQL last version
- [x] Docker

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:

* Você instalou GO `< 1.20 / requeridos>`
* Você instalou a versão`< DOCKER / requeridos>`

## 🚀 Executando

Após efetuar o clone do Back End execute os comandos:

- [x] go mod init api-auth  (para criar o go.mod)
- [x] go mod tidy  (para baixar as dependências)

Em seguida via terminal navegue até a pasta raiz do projeto onde encontra-se o arquivo docker-compose.yml e execute o comando:

- [x] docker-compose up

O Docker vai iniciar subindo um container com o MySQL, após isso, inicie o projeto pela classe `main.go`. 

A func `run` será responsável executar um AutoMigrate criando a table no MySql.

Para encerrar o container, utilize o comando `docker-compose down`. Esse comando remove o container e o banco de dados.
