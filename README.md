Simple API using Golang

Run:
- go run cmd/main.go

Swagger:
- http://localhost:8080/swagger/index.html
![image](https://github.com/user-attachments/assets/dbcd38d8-0d35-4a1f-89c7-f74dd07bd3ba)


Modify & Update Swagger:
- swag init --generalInfo cmd/main.go


Database: PostgreSQL

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
