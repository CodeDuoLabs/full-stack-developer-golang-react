package main

import (
	"log"
	"task_manager/internal/database"
	"task_manager/internal/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "task_manager/cmd/docs"
)

// @title        Task Manager API
// @version      1.0
// @description  A simple task management API built with Go and Gin.
// @host         localhost:8080
// @BasePath     /
func main() {
	db := database.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("Starting Api on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to initialize api server")
	}

}
