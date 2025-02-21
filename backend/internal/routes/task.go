package routes

import (
	"task_manager/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/tasks", handler.CreateTaskHandler(db))
	r.GET("/tasks", handler.GetTasksHandler(db))
	r.GET("/tasks/:id", handler.GetTaskHandler(db))
	r.PUT("/tasks/:id", handler.UpdateTaskHandler(db))
	r.DELETE("/tasks/:id", handler.DeleteTaskHandler(db))
}
