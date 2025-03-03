package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "Pending"
	StatusCompleted TaskStatus = "Completed"
)

type Task struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Status      TaskStatus     `gorm:"default:Pending" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (s TaskStatus) Validate() error {
	switch s {
	case StatusPending, StatusCompleted:
		return nil
	default:
		return errors.New("invalid status: must be 'Pending' or 'Completed'")
	}
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return t.Status.Validate()
}

func (t *Task) BeforeUpdate(tx *gorm.DB) error {
	return t.Status.Validate()
}
