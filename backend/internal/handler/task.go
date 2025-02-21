package handler

import (
	"net/http"
	"task_manager/internal/response"
	"task_manager/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func sendResponse(c *gin.Context, resp response.Response) {
	c.JSON(resp.Status, resp)
}

// CreateTaskHandler creates a new task in the system.
// @Summary      Create a new task
// @Description  Creates a task with the provided name and status.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      model.Task  true  "Task object to create"
// @Success      201   {object}  response.Response{data=model.Task}  "Created task"
// @Failure      400   {object}  response.Response  "Invalid request payload"
// @Failure      500   {object}  response.Response  "Failed to create task"
// @Router       /tasks [post]
func CreateTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}
		if err := db.Create(&task).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusInternalServerError, "Failed to create task"))
			return
		}
		sendResponse(c, response.NewSuccessResponse(http.StatusCreated, task))
	}
}

// GetTasksHandler retrieves all tasks from the system.
// @Summary      List all tasks
// @Description  Retrieves a list of all tasks stored in the database.
// @Tags         tasks
// @Produce      json
// @Success      200   {object}  response.Response{data=[]model.Task}  "List of tasks"
// @Failure      500   {object}  response.Response  "Failed to retrieve tasks"
// @Router       /tasks [get]
func GetTasksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks []model.Task
		if err := db.Find(&tasks).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusInternalServerError, "Failed to retrieve tasks"))
			return
		}
		sendResponse(c, response.NewSuccessResponse(http.StatusOK, tasks))
	}
}

// GetTaskHandler retrieves a specific task by ID.
// @Summary      Get a task
// @Description  Retrieves a task by its unique identifier.
// @Tags         tasks
// @Produce      json
// @Param        id   path      string       true  "Task ID (UUID)"
// @Success      200  {object}  response.Response{data=model.Task}  "Task details"
// @Failure      404  {object}  response.Response  "Task not found"
// @Router       /tasks/{id} [get]
func GetTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := db.First(&task, c.Param("id")).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusNotFound, "Task not found"))
			return
		}
		sendResponse(c, response.NewSuccessResponse(http.StatusOK, task))
	}
}

// UpdateTaskHandler updates an existing task.
// @Summary      Update a task
// @Description  Updates the details of a task identified by its ID.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "Task ID (UUID)"
// @Param        task  body      model.Task  true  "Updated task object"
// @Success      200   {object}  response.Response{data=model.Task}  "Updated task"
// @Failure      400   {object}  response.Response  "Invalid request payload"
// @Failure      404   {object}  response.Response  "Task not found"
// @Failure      500   {object}  response.Response  "Failed to update task"
// @Router       /tasks/{id} [put]
func UpdateTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := db.First(&task, c.Param("id")).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusNotFound, "Task not found"))
			return
		}
		if err := c.ShouldBindJSON(&task); err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}
		if err := db.Save(&task).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusInternalServerError, "Failed to update task"))
			return
		}
		sendResponse(c, response.NewSuccessResponse(http.StatusOK, task))
	}
}

// DeleteTaskHandler deletes a task by ID.
// @Summary      Delete a task
// @Description  Deletes a task identified by its unique identifier.
// @Tags         tasks
// @Produce      json
// @Param        id   path      string       true  "Task ID (UUID)"
// @Success      204  {object}  response.Response  "Task deleted successfully"
// @Failure      404  {object}  response.Response  "Task not found"
// @Failure      500  {object}  response.Response  "Failed to delete task"
// @Router       /tasks/{id} [delete]
func DeleteTaskHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := db.First(&task, c.Param("id")).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusNotFound, "Task not found"))
			return
		}
		if err := db.Delete(&task).Error; err != nil {
			sendResponse(c, response.NewErrorResponse(http.StatusInternalServerError, "Failed to delete task"))
			return
		}
		sendResponse(c, response.NewSuccessResponse(http.StatusNoContent, nil))
	}
}
