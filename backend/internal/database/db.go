package database

import (
	"log"
	"task_manager/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"),  &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&model.Task{})
	return db
}
