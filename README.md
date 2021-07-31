# sqlboiler-restapi-example

An boielrplate of REST API server using following tools.

- [SQLBoiler](https://github.com/volatiletech/sqlboiler) for ORM
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) for REST API Interface 

This repository is an addendum of this post.

https://future-architect.github.io/articles/20210730a/


## Prerequisites

- [air](https://github.com/cosmtrek/air): Run the server with hot reloading.
- PostgreSQL server

If you are opening this project via VSCode Remote Container, these tools are already configured in [docker-compose.yml](.devcontainer/docker-compose.yml)

## Development

### Project structure

```
.
├── boiler          # Generated codes by SQLBoiler
├── restapi         # Generated codes by oapi-codegen
├── sql             # Database table definitions
├── openapi.yml     # API definition
├── cmd
│   └── server      # Run server
├── testdata        # Test data
├── app.go          # Server implementation
├── db.go           # Database configurations
├── event.go
└── user.go
```


### Generate code

Regenerate code when you modify following codes.

- `sql/*.sql`: Database table definitions.
- `openapi.yml`: API definitions.

```
make clean
make generate
go mod tidy
```

### Run local server

```
air
```

## App Configurations

**Environment variables**

- `DB_URL`: Database URL. Defaults connect to the container running in Remote Container service.
