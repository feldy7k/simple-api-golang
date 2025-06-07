package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

var DB *pgx.Conn

func ConnectDatabase() {
	cfg := LoadConfig()

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
    )

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	DB = conn
	fmt.Println("Connected to PostgreSQL!")
}