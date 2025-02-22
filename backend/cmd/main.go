package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "task_manager/cmd/docs"
	"task_manager/internal/database"
	"task_manager/internal/middleware"
	"task_manager/internal/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        Task Manager API
// @version      1.0
// @description  A simple task management API built with Go and Gin.
// @host         localhost:8080
// @BasePath     /
func main() {
	db := database.InitDB()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Print("Starting Api on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize the server")
	}

}
