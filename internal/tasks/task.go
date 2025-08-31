package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func NewTask(title string, description *string) *Task {
	if description == nil {
		description =  &[]string{""}[0]
	}
	return &Task{
		ID: uuid.New().String(),
		Title: title,
		Description: *description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}