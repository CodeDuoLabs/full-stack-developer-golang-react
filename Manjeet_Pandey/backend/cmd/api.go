package main

import (
	"task_manager/internal/middleware"
	"task_manager/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func StartApi(db *gorm.DB) {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Print("Starting Api on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize the server")
	}
}
