package main

import (
	"go-api/config"
	"go-api/routes"
	"fmt"
)
// Author:			feldy j kambey

// @title           Go User API
// @version         1.0
// @description     A simple user CRUD API using Gin.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	config.ConnectDatabase()
	routes.SetupAndRun()
	cfg := config.LoadConfig()
    fmt.Println("Server DB Postgresql running on port:", cfg.DBPort)
}