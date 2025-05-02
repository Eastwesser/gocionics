# Socionics app using Golang

This is a simple app to show you what type of character you are!
We are using Golang and PostgreSQL to build a RESTful API.

## Project structure

```bash
gocionics/
├── cmd/
│   └── api/
│       └── v1/
│           └── main.go          # Точка входа + Swagger
├── config/
│   └── config.go               # Конфигурация (.env)
├── internal/
│   ├── app/                    # Инициализация приложения
│   │   └── app.go
│   ├── controllers/            # HTTP-контроллеры (Gin)
│   │   ├── auth/
│   │   │   └── auth.go
│   │   ├── user/
│   │   │   └── user.go
│   │   └── character/
│   │       └── character.go
│   ├── db/                     # Работа с БД
│   │   ├── database.go         # Подключение PostgreSQL
│   │   ├── migrations/         # Миграции Goose
│   │   └── sql.go              # SQL-запросы
│   ├── entities/               # Сущности
│   │   ├── user/
│   │   │   └── user.go
│   │   └── character/
│   │       └── character.go
│   ├── handlers/               # HTTP-обработчики (альтернатива контроллерам)
│   │   ├── user/
│   │   │   └── user.go
│   │   └── character/
│   │       └── character.go
│   ├── repositories/           # Репозитории (БД)
│   │   ├── user/
│   │   │   └── user.go
│   │   └── character/
│   │       └── character.go
│   ├── server/                 # HTTP-сервер
│   │   └── server.go
│   ├── services/               # Доп. сервисы (анализ теста)
│   │   └── test-analyzer.go
│   └── usecases/               # Бизнес-логика
│       ├── auth/
│       │   └── auth.go
│       ├── user/
│       │   └── user.go
│       └── character/
│           └── character.go
├── pkg/                        # Внешние утилиты
│   └── httpclient/             # HTTP-клиент с таймаутами
│       └── client.go
└── docs/                       # Swagger
    └── swagger/
        └── docs.go
```


To create folders:

```bash
    md "sub\sub dir\other dir"
```

To create multiple **embedded** folders:
```bash
    "entity", "usecases", "repositories", "handlers", "services", "app" | ForEach-Object { md -Force "internal/$_" }
```

To create a file:
```bash
    ni internal/entities/user.go -Type File
```

## Step 1: Create entities (for user and character)

Work with the directory "...\gocionics\internal\entities"

Create there 2 files, **user.go** and **character.go**, and put there your structures.

- What is an "entity" in this project?

- What are the "DTOs" of this project?

## Step 2: Create Use Case
We create a User structure and methods with panic("implement me")

## Step 3: Create Repository
Here we create exact interface to work with our User structure. 
So the structure does exactly the thing it's asked for, realising the interface.

## Step 4: Create a Controller
Create a controller for auth.

## Step 5: Database realisation
In database.go create your db implementation

## Step 6: Connect all layers
In main.go we connect all our layers 

## Step 7: Added Readmes for all folders to show what's there

## Step 8: After working with entities -> repos:
We code a constructor for our usecases

## Step 9: We realise the usecases
Set registration + set character

## Step 10: Handlers
Write code logic for methods and CRUDs

### TIPS:

✅ Clean Architecture:

    Entities → Use Cases → Repositories → Controllers.

Outer layers (DB, HTTP) depend on inner layers.

✅ Additional Components:

    HTTP client with timeouts.
    
    Swagger for documentation.
    
    Goose for migrations.

✅ Ready Use Cases:

    Authentication (auth).
    
    Working with socionic types (character).

## SQL Migrations

We are using **goose** for db migrations:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Add to go.mod:
```bash
go get github.com/pressly/goose/v3
```

Check version:
```bash
goose -version
```


To create first files for migrations use:
```bash
goose -dir internal/db/migrations create add_users_table sql
```

## Swagger
We use swaggo in this project

Type to generate docs:
```bash
swag init -g cmd/api/v1/main.go
```

More acute way:
```bash
swag init -g cmd/api/v1/main.go --output docs/swagger --parseDependency --parseInternal
```

## Docker

Use this command to run server:
```bash
docker-compose up --build
```

Also check migrations:
```bash
docker-compose exec app goose -dir ./migrations postgres "user=postgres password=postgres dbname=library host=db port=5432 sslmode=disable" status
```

## Final steps

Open your browser and enter **http://localhost:8080/swagger/index.html**

This is your documentation, congratulations!