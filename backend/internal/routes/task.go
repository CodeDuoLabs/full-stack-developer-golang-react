package routes

import (
	"task_manager/internal/handler"
	"task_manager/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	taskService := service.NewTaskService(db)
	taskHandler := handler.NewTaskHandler(taskService)
	r.POST("/tasks", taskHandler.CreateTaskHandler())
	r.GET("/tasks", taskHandler.GetTasksHandler())
	r.GET("/tasks/:id", taskHandler.GetTaskHandler())
	r.PUT("/tasks/:id", taskHandler.UpdateTaskHandler())
	r.DELETE("/tasks/:id", taskHandler.DeleteTaskHandler())
}
