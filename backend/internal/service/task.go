package service

import (
	"task_manager/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService {
	return TaskService{db: db}
}

func (s *TaskService) CreateTask(task model.Task) error {
	if err := s.db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (s *TaskService) ListTask() ([]model.Task, error) {
	var tasks []model.Task
	if err := s.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTask(id uuid.UUID) (*model.Task, error) {
	var task model.Task
	if err := s.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(id uuid.UUID, task model.Task) error {
	if err := s.db.Where("id = ?", id).Updates(&task).Error; err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(id uuid.UUID) error {
	if err := s.db.Delete(&model.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
