package main

import (
	"os"
	_ "task_manager/cmd/docs"
	"task_manager/internal/database"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title        Task Manager API
// @version      1.0
// @description  A simple task management API built with Go and Gin.
// @host         localhost:8080
// @BasePath     /
func main() {
	db := database.InitDB()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// var args []string
	args := os.Args
	if len(args) < 2 {
		log.Fatal().Msg("Not enough argument")
	}
	switch args[1] {
	case "api":
		StartApi(db)
	default:
		log.Fatal().Msg("Don;t know what to do")
	}
}
