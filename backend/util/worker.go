package util

import (
	"fmt"
	"sync"
	"task_manager/internal/service"
	"task_manager/model"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Worker struct {
	workers     int
	channel     chan uuid.UUID
	wg          sync.WaitGroup
	taskService service.TaskService
}

func NewWorker(workers int, taskService service.TaskService) *Worker {
	return &Worker{
		workers:     workers,
		channel:     make(chan uuid.UUID),
		taskService: taskService,
	}
}

func (w *Worker) StartWorker() {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go w.processTask()
	}
}

func (w *Worker) AddToQueue(taskId uuid.UUID) {
	w.channel <- taskId
}

func (w *Worker) processTask() {
	defer w.wg.Done()
	for {
		select {
		case taskId := <-w.channel:
			fmt.Println("Processing Task with ID", taskId)
			// Do something here , Like some actual work
			fmt.Println("Cooking something here")
			err := w.taskService.UpdateTask(taskId, model.Task{
				Status: model.StatusCompleted,
			})
			if err != nil {
				fmt.Println("Cannot Process Task")
				log.Err(err).Msg("Cannot update task")
			}
			fmt.Println("Processed Task with ID", taskId)
		default:
			return
		}
	}

}

func (w *Worker) Wait() {
	w.wg.Wait()
}
