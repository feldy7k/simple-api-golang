package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    DBPort     string
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
}

func LoadConfig() Config {
    // Load .env file (optional)
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: No .env file found")
    }

    return Config{
        DBPort:     os.Getenv("DB_PORT"),
        DBHost:     os.Getenv("DB_HOST"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
    }
}