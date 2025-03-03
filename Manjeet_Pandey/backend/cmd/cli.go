package main

import (
	"fmt"
	"task_manager/internal/service"
	"task_manager/model"
	"task_manager/util"

	"github.com/rs/zerolog/log"
)

type CliHandler struct {
	taskService service.TaskService
}

func NewCliHandler(service service.TaskService) CliHandler {
	return CliHandler{
		taskService: service,
	}
}

func FormatListOutput(entries []model.Task) {
	fmt.Println("  ---------------------------------------------------------------------------------------------")
	fmt.Printf("| %-20s | %-20s | %-45s |\n", "Name", "Description", "Status")
	fmt.Println("|----------------------|----------------------|-----------------------------------------------|")
	for _, entry := range entries {
		name := entry.Name
		description := entry.Description
		status := entry.Status
		fmt.Printf("| %-20s | %-20s | %-45s |\n", name, description, status)
	}
	fmt.Println("  ---------------------------------------------------------------------------------------------")
}

func (t *CliHandler) AddTask(name, description string) {
	task := model.Task{
		Name:        name,
		Description: description,
		Status:      model.StatusPending,
	}

	if err := t.taskService.CreateTask(task); err != nil {
		log.Err(err).Msg("Error creating task")
		return
	}
	FormatListOutput([]model.Task{task})
	return
}

func (t *CliHandler) ListTask() {
	tasks, err := t.taskService.ListTask()
	if err != nil {
		log.Err(err).Msg("Error Listing Task")
		return
	}
	FormatListOutput(tasks)
	return
}

func (t *CliHandler) ProcessTask() {
	numWorker := 5
	tasks, err := t.taskService.GetPendingTasks()
	if err != nil {
		log.Err(err).Msg("Cannot get Pending Task")
		return
	}
	workers := util.NewWorker(numWorker, t.taskService)

	workers.StartWorker()
	for _, task := range tasks {
		workers.AddToQueue(task.ID)
	}
	workers.Wait()
	return
}
