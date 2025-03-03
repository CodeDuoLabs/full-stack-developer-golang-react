package handler

import (
	"net/http"
	"task_manager/internal/response"
	"task_manager/internal/service"
	"task_manager/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func sendResponse(c *gin.Context, resp response.Response) {
	c.JSON(resp.Status, resp)
}

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(service service.TaskService) TaskHandler {
	return TaskHandler{
		taskService: service,
	}
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
// @ID CreateTask
func (t *TaskHandler) CreateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			log.Err(err).Msg("Invalid payload")
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}
		if err := t.taskService.CreateTask(task); err != nil {
			log.Err(err).Msg("Error creating task")
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
// @ID ListTasks
func (t *TaskHandler) GetTasksHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := t.taskService.ListTask()
		if err != nil {
			log.Err(err).Msg("Error retreiving tasks")
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
// @ID GetTaskByID
func (t *TaskHandler) GetTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			log.Err(err).Msg("Error parsing uuid")
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, "Invalid task id"))
		}
		task, err := t.taskService.GetTask(id)
		if err != nil {
			log.Err(err).Msg("Task not found")
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
// @ID UpdateTask
func (t *TaskHandler) UpdateTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			log.Err(err).Msg("Error parsing uuid")
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, "Invalid task id"))
		}
		if err := c.ShouldBindJSON(&task); err != nil {
			log.Err(err).Msg("Error binding payload")
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}
		if err := t.taskService.UpdateTask(id, task); err != nil {
			log.Err(err).Msg("Failed to update Task")
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
// @ID DeleteTask
func (t *TaskHandler) DeleteTaskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			log.Err(err).Msg("Error parsing uuid")
			sendResponse(c, response.NewErrorResponse(http.StatusBadRequest, "Invalid task id"))
		}
		if err := t.taskService.DeleteTask(id); err != nil {
			log.Err(err).Msg("Failed to delete task")
			sendResponse(c, response.NewErrorResponse(http.StatusInternalServerError, "Failed to delete task"))
			return
		}
		sendResponse(c, response.NewSuccessResponse(http.StatusOK, "Deleted successfully"))
	}
}
