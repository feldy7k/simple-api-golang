Simple API using Golang

Run:
- go run cmd/main.go

Swagger:
- http://localhost:8080/swagger/index.html

Modify & Update Swagger:
- swag init --generalInfo cmd/main.go


Database: PostgreSQL

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);