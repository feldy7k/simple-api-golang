package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

var DB *pgx.Conn

func ConnectDatabase() {
	dsn := "postgres://postgres:pgfeldymaster@localhost:5432/coursedb"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	DB = conn
	fmt.Println("Connected to PostgreSQL!")
}