Simple API using Golang

Created: June 2025

Run:
- go run cmd/main.go

Swagger:
- http://localhost:8080/swagger/index.html

![image](https://github.com/user-attachments/assets/415b6a46-ba7b-4a2f-8144-b09813371cd6)


Modify & Update Swagger:
- swag init --generalInfo cmd/main.go


Database: postgresql-17.5-1

Golang ver: go1.24.4

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
