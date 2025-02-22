package database

import (
	"task_manager/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	db.AutoMigrate(&model.Task{})
	return db
}
