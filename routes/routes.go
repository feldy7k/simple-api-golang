package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-api/controllers"
	_ "go-api/docs"
	"go-api/config"
)

func SetupAndRun() {
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users/:id", controllers.GetUser)
		v1.POST("/users", controllers.CreateUser)
		v1.PUT("/users", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)

		// auth
		v1.POST("/auth", controllers.AuthLogin)
	}

	cfg := config.LoadConfig()
	r.Run(":" + cfg.HTTPPort)
}